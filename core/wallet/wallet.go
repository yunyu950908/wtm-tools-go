package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"wtm-tools-go/config"
	"wtm-tools-go/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    common.Address
}

func NewWallet(walletConfig config.WalletConfig) *Wallet {
	privateKeyStr := strings.TrimPrefix(walletConfig.PrivateKeyHead+walletConfig.PrivateKeyTail, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid public key")
	}
	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
		address:    crypto.PubkeyToAddress(*publicKeyECDSA),
	}
}

func (w *Wallet) Address() common.Address {
	return w.address
}

func (w *Wallet) GetSigner(ctx context.Context, client *ethclient.Client) (bind.SignerFn, error) {
	chainID, err := utils.GetChainID(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.LatestSignerForChainID(chainID), w.privateKey)
	}, nil
}

// SendTransaction 发送交易
func (w *Wallet) SendTransaction(
	ctx context.Context,
	client *ethclient.Client,
	to common.Address,
	value *big.Int,
	data []byte,
) (string, error) {
	// 获取 nonce
	nonce, err := client.PendingNonceAt(ctx, w.address)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// 建议 Gas 价格
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// 估算 gas 数量
	gasLimit, err := utils.EstimateGas(ctx, client, w.address, to, value, data)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas: %v", err)
	}

	// 创建交易
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// 签名交易
	signer, err := w.GetSigner(ctx, client)
	if err != nil {
		return "", fmt.Errorf("failed to get signer: %v", err)
	}
	signedTx, err := signer(w.address, tx)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}
	// 发送交易
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}
