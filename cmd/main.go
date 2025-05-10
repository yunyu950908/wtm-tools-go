package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"wtm-tools-go/config"
	"wtm-tools-go/core/dex/oogabooga"
	"wtm-tools-go/core/infrared"
	"wtm-tools-go/core/scheduler"
	"wtm-tools-go/core/wallet"
	"wtm-tools-go/logger"
	"wtm-tools-go/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/cobra"
)

type TokenInfo struct {
	Address  string
	Symbol   string
	Decimals int
}

type TokenPrice struct {
	Address string
	Price   float64 // price in USD
}

func main() {
	fmt.Println("Hello, World!")

	var configFilePath string

	// define root command
	var rootCmd = &cobra.Command{
		Use:   "wtm-tools-go",
		Short: "wtm-tools-go",
		Run: func(cmd *cobra.Command, args []string) {
			// read toml content from config file
			tomlContent, err := os.ReadFile(configFilePath)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to read config file")
				os.Exit(1)
			}
			run(tomlContent)
		},
	}

	// add -f flag, for specifying config file path
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "file", "f", "", "config file path")

	// mark config file path flag as required
	if err := rootCmd.MarkPersistentFlagRequired("file"); err != nil {
		logger.Error().Err(err).Msg("Failed to mark config file path flag as required")
		os.Exit(1)
	}

	// cobra will automatically provide -h/--help
	if err := rootCmd.Execute(); err != nil {
		logger.Error().Err(err).Msg("Failed to execute root command")
		os.Exit(1)
	}

}

func run(tomlContent []byte) {
	cfg := config.Provide(string(tomlContent))
	walletConfig := config.ProvideWalletConfig(cfg)
	rpcConfig := config.ProvideRPCConfig(cfg)
	taskConfig := config.ProvideTaskConfig(cfg)
	oogaboogaConfig := config.ProvideOogaboogaConfig(cfg)

	tz, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to load timezone")
		os.Exit(1)
	}

	berachainClient, err := ethclient.Dial(rpcConfig.BerachainHTTP)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to Berachain")
		os.Exit(1)
	}

	w := wallet.NewWallet(walletConfig)

	s := scheduler.NewScheduler(tz)

	tokenInfoCache := cache.New(5*time.Minute, 10*time.Minute)
	tokenPriceCache := cache.New(5*time.Minute, 10*time.Minute)

	oogaboogaClient := new(oogabooga.Client)
	if oogaboogaConfig.BaseURL != "" && oogaboogaConfig.APIKey != "" {
		oogaboogaClient = oogabooga.GetClient(oogaboogaConfig)
		if err := s.RegisterTask("oogabooga-sync", "*/1 * * * *", func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			tokenList, err := oogaboogaClient.GetTokenList(ctx)
			if err != nil {
				return fmt.Errorf("failed to get token list: %w", err)
			}
			for _, token := range tokenList {
				address := strings.ToLower(token.Address)
				tokenInfoCache.Set(address, TokenInfo{
					Address:  address,
					Symbol:   token.Symbol,
					Decimals: token.Decimals,
				}, cache.DefaultExpiration)
				// logger.Debug().
				// 	Str("address", token.Address).
				// 	Str("symbol", token.Symbol).
				// 	Int("decimals", token.Decimals).
				// 	Msg("token info")
			}
			tokenPrices, err := oogaboogaClient.GetTokenPrices(ctx)
			if err != nil {
				return fmt.Errorf("failed to get token prices: %w", err)
			}
			for _, tokenPrice := range tokenPrices {
				address := strings.ToLower(tokenPrice.Address)
				tokenPriceCache.Set(address, TokenPrice{
					Address: address,
					Price:   tokenPrice.Price,
				}, cache.DefaultExpiration)
				// logger.Debug().
				// 	Str("address", tokenPrice.Address).
				// 	Float64("price", tokenPrice.Price).
				// 	Msg("token price")
			}
			return nil
		}); err != nil {
			logger.Error().Err(err).Msg("Failed to register oogabooga token list task")
			os.Exit(1)
		}
	}

	for _, t := range taskConfig.Claim {
		if err := s.RegisterTask(t.Name, t.Cron, func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			switch t.Provider {
			case config.VaultProviderInfrared:
				return infrared.NewVault(t.VaultAddress, berachainClient, w).Claim(ctx)
			default:
				return fmt.Errorf("unsupported provider: %s", t.Provider)
			}
		}); err != nil {
			logger.Error().Err(err).Msg("Failed to register claim task")
			os.Exit(1)
		}
	}

	for _, task := range taskConfig.Stake {
		if err := s.RegisterTask(task.Name, task.Cron, func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			switch task.Provider {
			case config.VaultProviderInfrared:
				vault := infrared.NewVault(task.VaultAddress, berachainClient, w)
				stakingToken, err := vault.GetStakingToken(ctx)
				if err != nil {
					return fmt.Errorf("failed to get staking token: %w", err)
				}
				// check balance
				remainedStakingTokenBalance, err := utils.CheckBalance(ctx, stakingToken.String(), w.Address(), berachainClient)
				if err != nil {
					return fmt.Errorf("failed to check balance: %w", err)
				}

				// check if this token is available in oogabooga
				_, ok := tokenPriceCache.Get(stakingToken.String())
				if !ok {
					logger.Info().Msg("staking token price not found, staking all")
					// stake
					return vault.Stake(ctx, remainedStakingTokenBalance.Balance)
				}

				// check if auto swap is enabled
				if task.AutoSwap {
					prepareStakeVauleInUSD := 0.0
					// check whitelisted assets
					for _, asset := range task.AutoSwapAssets {
						// if prepareStakeVauleInUSD is greater than maxAmountInUSD, break
						if prepareStakeVauleInUSD >= task.MaxAmountInUSD {
							break
						}
						// check balance
						assetBalance, err := utils.CheckBalance(ctx, asset, w.Address(), berachainClient)
						if err != nil {
							return fmt.Errorf("failed to check asset balance: %w", err)
						}
						if assetBalance.Balance.Cmp(big.NewInt(0)) > 0 {
							// check if this token is available in oogabooga
							cached, ok := tokenPriceCache.Get(stakingToken.String())
							if !ok {
								continue
							}
							cachedTokenPrice, ok := cached.(TokenPrice)
							if !ok {
								continue
							}
							valueInUSD := assetBalance.Amount * cachedTokenPrice.Price
							// value in USD is too low, skip
							if valueInUSD < 1 {
								continue
							}
							if valueInUSD >= task.MaxAmountInUSD {
								valueInUSD = task.MaxAmountInUSD
							}
							swapAmountIn := utils.FloatToBigIntWithDecimal(valueInUSD/cachedTokenPrice.Price, assetBalance.Decimals)
							// check allowance
							allowance, err := oogaboogaClient.GetAllowance(ctx, asset, w.Address().String())
							if err != nil {
								return fmt.Errorf("failed to get allowance: %w", err)
							}
							// if allowance is not enough, approve
							allowanceInt, ok := new(big.Int).SetString(allowance.Allowance, 10)
							if !ok {
								return fmt.Errorf("failed to convert allowance to big.Int")
							}
							if allowanceInt.Cmp(swapAmountIn) < 0 {
								// approve
								approveTx, err := oogaboogaClient.GetApproveAllowanceTx(ctx, asset, utils.MAX_UINT256)
								if err != nil {
									return fmt.Errorf("failed to get approve allowance tx: %w", err)
								}
								// send tx
								approveTxData, err := utils.DecodeHexToBytes(approveTx.Tx.Data)
								if err != nil {
									return fmt.Errorf("failed to decode hex to bytes for approve allowance tx: %w", err)
								}
								txHash, err := w.SendTransaction(ctx, berachainClient, common.HexToAddress(approveTx.Tx.To), nil, approveTxData)
								if err != nil {
									return fmt.Errorf("failed to send approve allowance tx: %w", err)
								}
								// wait for tx to be mined
								receipt, err := utils.CheckReceipt(berachainClient, common.HexToHash(txHash), nil)
								if err != nil {
									return fmt.Errorf("failed to check receipt: %w", err)
								}
								if receipt.Status != types.ReceiptStatusSuccessful {
									return fmt.Errorf("approve allowance tx failed")
								}
							}
							// swap
							swapTx, err := oogaboogaClient.GetSwapTx(ctx, oogabooga.SwapParams{
								TokenIn:  asset,
								TokenOut: stakingToken.String(),
								Amount:   swapAmountIn.String(),
								To:       w.Address().String(),
								Slippage: task.AutoSwapSlippage,
							})
							if err != nil {
								return fmt.Errorf("failed to get swap tx: %w", err)
							}
							// send tx
							swapTxData, err := utils.DecodeHexToBytes(swapTx.Tx.Data)
							if err != nil {
								return fmt.Errorf("failed to decode hex to bytes for swap tx: %w", err)
							}
							txHash, err := w.SendTransaction(ctx, berachainClient, common.HexToAddress(swapTx.Tx.To), nil, swapTxData)
							if err != nil {
								return fmt.Errorf("failed to send swap tx: %w", err)
							}
							// wait for tx to be mined
							receipt, err := utils.CheckReceipt(berachainClient, common.HexToHash(txHash), nil)
							if err != nil {
								return fmt.Errorf("failed to check receipt: %w", err)
							}
							if receipt.Status != types.ReceiptStatusSuccessful {
								return fmt.Errorf("swap tx failed")
							}
							prepareStakeVauleInUSD += valueInUSD
						}
					}
				}
				// stake
				// check balance
				balance, err := utils.CheckBalance(ctx, stakingToken.String(), w.Address(), berachainClient)
				if err != nil {
					return fmt.Errorf("failed to check balance: %w", err)
				}
				if balance.Balance.Cmp(big.NewInt(0)) == 0 {
					return fmt.Errorf("staking token balance is 0")
				}
				return vault.Stake(ctx, balance.Balance)
			default:
				return fmt.Errorf("unsupported provider: %s", task.Provider)
			}
		}); err != nil {
			logger.Error().Err(err).Msg("Failed to register stake task")
			os.Exit(1)
		}
	}

	if err := s.Start(); err != nil {
		logger.Error().Err(err).Msg("Failed to start scheduler")
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	if err != nil {
		logger.Error().Err(err).Msg("err")
	}
}
