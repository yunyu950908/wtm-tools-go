package main

import (
	"context"
	"flag"
	"math/big"
	"os"
	"strings"
	"wtm-tools-go/config"
	"wtm-tools-go/core/dex/oogabooga"
	"wtm-tools-go/core/wallet"
	"wtm-tools-go/logger"
	"wtm-tools-go/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Command-line arguments
	configFilePath := flag.String("f", "", "config file path")
	flag.Parse()

	if *configFilePath == "" {
		logger.Error().Msg("config file path is required")
		os.Exit(1)
	}

	// read toml content from config file
	tomlContent, err := os.ReadFile(*configFilePath)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read config file")
		os.Exit(1)
	}

	cfg := config.Provide(string(tomlContent))
	walletConfig := config.ProvideWalletConfig(cfg)
	rpcConfig := config.ProvideRPCConfig(cfg)
	oogaboogaConfig := config.ProvideOogaboogaConfig(cfg)

	client, err := ethclient.Dial(rpcConfig.BerachainHTTP)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to Berachain")
		os.Exit(1)
	}
	wallet := wallet.NewWallet(walletConfig)
	oogaboogaClient := oogabooga.NewClient(oogaboogaConfig)

	ctx := context.Background()

	tokenList, err := oogaboogaClient.GetTokenList(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get token list")
		os.Exit(1)
	}
	tokenPrices, err := oogaboogaClient.GetTokenPrices(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get token prices")
		os.Exit(1)
	}

	type TokenInfo struct {
		Address  string
		Symbol   string
		Decimals int
		Price    float64
	}

	tokenInfoMap := make(map[string]TokenInfo)
	for _, token := range tokenList {
		address := strings.ToLower(token.Address)
		tokenInfoMap[address] = TokenInfo{
			Address:  address,
			Symbol:   token.Symbol,
			Decimals: token.Decimals,
		}
	}
	for _, price := range tokenPrices {
		address := strings.ToLower(price.Address)
		tokenInfoMap[address] = TokenInfo{
			Address:  address,
			Symbol:   tokenInfoMap[address].Symbol,
			Decimals: tokenInfoMap[address].Decimals,
			Price:    price.Price,
		}
	}

	// honey and usdc addresses
	const (
		honeyAddress = "0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce"
		usdcAddress  = "0x549943e04f40284185054145c6e4e9568c1d3241"
	)

	honeyTokenInfo := tokenInfoMap[honeyAddress]
	usdcTokenInfo := tokenInfoMap[usdcAddress]

	// assert honey token address equals tokenInfoMap[honeyAddress].Address
	if honeyTokenInfo.Address != honeyAddress {
		logger.Error().Msgf("HONEY token address does not match: %s != %s", honeyTokenInfo.Address, honeyAddress)
		os.Exit(1)
	}
	// assert usdc token address equals tokenInfoMap[usdcAddress].Address
	if usdcTokenInfo.Address != usdcAddress {
		logger.Error().Msgf("USDC token address does not match: %s != %s", usdcTokenInfo.Address, usdcAddress)
		os.Exit(1)
	}

	// check allowance for honey token
	allowanceOfHoneyToken, err := oogaboogaClient.GetAllowance(ctx, honeyAddress, wallet.Address().Hex())
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get allowance for honey token")
		os.Exit(1)
	}
	logger.Info().Msgf("Allowance for honey token: %s", allowanceOfHoneyToken.Allowance)

	// 0.1 USD
	scaleOfHoneyToken := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(honeyTokenInfo.Decimals-1)), nil)
	logger.Info().Msgf("Scale of honey token: %s", scaleOfHoneyToken.String())
	allowanceOfHoneyTokenInt, ok := new(big.Int).SetString(allowanceOfHoneyToken.Allowance, 10)
	if !ok {
		logger.Error().Msg("Failed to convert allowance for honey token to big.Int")
		os.Exit(1)
	}

	if allowanceOfHoneyTokenInt.Cmp(scaleOfHoneyToken) < 0 {
		approveTx, err := oogaboogaClient.GetApproveAllowanceTx(ctx, honeyAddress, utils.MAX_UINT256)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get approve allowance tx for honey token")
			os.Exit(1)
		}
		logger.Info().Msgf("Approve allowance tx for honey token: %s", approveTx.Tx.Data)
		approveTxData, err := utils.DecodeHexToBytes(approveTx.Tx.Data)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to decode hex to bytes for approve allowance tx for honey token")
			os.Exit(1)
		}
		txHash, err := wallet.SendTransaction(ctx, client, common.HexToAddress(approveTx.Tx.To), nil, approveTxData)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to send approve allowance tx for honey token")
			os.Exit(1)
		}
		receipt, err := utils.CheckReceipt(client, common.HexToHash(txHash), nil)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to check receipt for approve allowance tx for honey token")
			os.Exit(1)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			logger.Error().Msg("Approve allowance tx for honey token failed")
			os.Exit(1)
		}
		logger.Info().Msgf("Approve allowance tx for honey token: %s", txHash)
	}

	// check balance for honey token
	balanceOfHoneyToken, err := utils.CheckBalance(ctx, honeyAddress, wallet.Address(), client)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get balance for honey token")
		os.Exit(1)
	}
	logger.Info().Msgf("Balance for honey token: %s", balanceOfHoneyToken.Balance.String())

	if balanceOfHoneyToken.Balance.Cmp(scaleOfHoneyToken) < 0 {
		logger.Error().Msg("Balance for honey token is less than scale of honey token")
		os.Exit(1)
	}

	// swap honey token to usdc token
	swapTx, err := oogaboogaClient.GetSwapTx(ctx, oogabooga.SwapParams{
		TokenIn:  honeyAddress,
		TokenOut: usdcAddress,
		Amount:   scaleOfHoneyToken.String(),
		To:       wallet.Address().Hex(),
		Slippage: 0.005, // 0.5%
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get swap tx for honey token to usdc token")
		os.Exit(1)
	}
	logger.Info().Msgf("Swap tx for honey token to usdc token: to=%s, data=%s", swapTx.Tx.To, swapTx.Tx.Data)
	swapTxData, err := utils.DecodeHexToBytes(swapTx.Tx.Data)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode hex to bytes for swap tx for honey token to usdc token")
		os.Exit(1)
	}
	txHash, err := wallet.SendTransaction(ctx, client, common.HexToAddress(swapTx.Tx.To), nil, swapTxData)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to send swap tx for honey token to usdc token")
		os.Exit(1)
	}
	receipt, err := utils.CheckReceipt(client, common.HexToHash(txHash), nil)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to check receipt for swap tx for honey token to usdc token")
		os.Exit(1)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		logger.Error().Msg("Swap tx for honey token to usdc token failed")
		os.Exit(1)
	}
	logger.Info().Msgf("Swap tx for honey token to usdc token: %s", txHash)
}
