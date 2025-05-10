package infrared

import (
	"context"
	"fmt"
	"math/big"
	"wtm-tools-go/core/wallet"
	"wtm-tools-go/internal/sc"
	"wtm-tools-go/logger"
	"wtm-tools-go/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Vault struct {
	client       *ethclient.Client
	vaultAddress common.Address
	vault        *sc.Vault
	wallet       *wallet.Wallet
}

var vaults = make(map[string]*Vault)

func NewVault(vaultAddress string, client *ethclient.Client, wallet *wallet.Wallet) *Vault {
	if _, ok := vaults[vaultAddress]; ok {
		return vaults[vaultAddress]
	}
	vault, err := sc.NewVault(common.HexToAddress(vaultAddress), client)
	if err != nil {
		panic(err)
	}
	vaults[vaultAddress] = &Vault{
		client:       client,
		vaultAddress: common.HexToAddress(vaultAddress),
		vault:        vault,
		wallet:       wallet,
	}
	return vaults[vaultAddress]
}

func (v *Vault) GetAllRewardTokens(ctx context.Context) ([]common.Address, error) {
	rewardTokens, err := v.vault.GetAllRewardTokens(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get all reward tokens: %w", err)
	}
	return rewardTokens, nil
}

func (v *Vault) GetAllRewardsForUser(ctx context.Context) ([]sc.IInfraredVaultUserReward, error) {
	rewards, err := v.vault.GetAllRewardsForUser(&bind.CallOpts{
		Context: ctx,
	}, v.wallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to get all rewards for user: %w", err)
	}
	return rewards, nil
}

func (v *Vault) Claim(ctx context.Context) error {
	signer, err := v.wallet.GetSigner(ctx, v.client)
	if err != nil {
		return fmt.Errorf("failed to get signer: %w", err)
	}
	tx, err := v.vault.GetReward(&bind.TransactOpts{
		Context: ctx,
		From:    v.wallet.Address(),
		Signer:  signer,
	})
	if err != nil {
		return fmt.Errorf("failed to claim reward: %w", err)
	}
	receipt, err := utils.CheckReceipt(v.client, tx.Hash(), nil)
	if err != nil {
		return fmt.Errorf("failed to check receipt: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("claim failed")
	}
	logger.Info().Str("tx", tx.Hash().String()).Msg("claim reward success")
	return nil
}

func (v *Vault) Withdraw(ctx context.Context, amount *big.Int) error {
	signer, err := v.wallet.GetSigner(ctx, v.client)
	if err != nil {
		return fmt.Errorf("failed to get signer: %w", err)
	}
	tx, err := v.vault.Withdraw(&bind.TransactOpts{
		Context: ctx,
		From:    v.wallet.Address(),
		Signer:  signer,
	}, amount)
	if err != nil {
		return fmt.Errorf("failed to withdraw: %w", err)
	}
	receipt, err := utils.CheckReceipt(v.client, tx.Hash(), nil)
	if err != nil {
		return fmt.Errorf("failed to check receipt: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("withdraw failed")
	}
	return nil
}

func (v *Vault) Stake(ctx context.Context, amount *big.Int) error {
	signer, err := v.wallet.GetSigner(ctx, v.client)
	if err != nil {
		return fmt.Errorf("failed to get signer: %w", err)
	}
	tx, err := v.vault.Stake(&bind.TransactOpts{
		Context: ctx,
		From:    v.wallet.Address(),
		Signer:  signer,
	}, amount)
	if err != nil {
		return fmt.Errorf("failed to stake: %w", err)
	}
	receipt, err := utils.CheckReceipt(v.client, tx.Hash(), nil)
	if err != nil {
		return fmt.Errorf("failed to check receipt: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("stake failed")
	}
	logger.Info().Str("tx", tx.Hash().String()).Msg("stake success")
	return nil
}

func (v *Vault) GetStakingToken(ctx context.Context) (common.Address, error) {
	stakingToken, err := v.vault.StakingToken(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get staking token: %w", err)
	}
	return stakingToken, nil
}

func (v *Vault) GetStakedAmount(ctx context.Context) (*big.Int, error) {
	balance, err := v.vault.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, v.wallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to get balance of staked token: %w", err)
	}
	return balance, nil
}
