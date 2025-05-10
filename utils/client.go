package utils

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"
	"wtm-tools-go/logger"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CheckReceiptOptions 定义可选参数的结构体
type CheckReceiptOptions struct {
	PollInterval time.Duration // 轮询间隔时间
	Timeout      time.Duration // 最长等待时间
}

// DefaultCheckReceiptOptions 返回默认配置
func DefaultCheckReceiptOptions() CheckReceiptOptions {
	return CheckReceiptOptions{
		PollInterval: 2 * time.Second,  // 默认轮询间隔为 2 秒
		Timeout:      30 * time.Second, // 默认超时时间为 30 秒
	}
}

// CheckReceipt 等待交易上链并获取收据
// 参数:
//   - client: 以太坊客户端，用于与节点交互
//   - txHash: 交易的哈希值
//   - opts: 可选参数，传入 CheckReceiptOptions 结构体指针，为 nil 时使用默认值
//
// 返回:
//   - *types.Receipt: 交易收据
//   - error: 可能的错误
func CheckReceipt(client *ethclient.Client, txHash common.Hash, opts *CheckReceiptOptions) (*types.Receipt, error) {
	options := DefaultCheckReceiptOptions()
	if opts != nil {
		if opts.PollInterval > 0 {
			options.PollInterval = opts.PollInterval
		}
		if opts.Timeout > 0 {
			options.Timeout = opts.Timeout
		}
	}
	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel()

	// 轮询交易收据
	for {
		select {
		case <-ctx.Done():
			// 超时或上下文被取消
			return nil, fmt.Errorf("timeout waiting for transaction receipt: %s", txHash.Hex())
		default:
			// 查询交易收据
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if err != nil {
				// 交易还未上链
				if errors.Is(err, ethereum.NotFound) {
					logger.Debug().Str("txHash", txHash.Hex()).Msg("Waiting for transaction to be mined...")
					time.Sleep(options.PollInterval)
					continue
				}
				// 其他错误
				return nil, fmt.Errorf("failed to get transaction receipt: %v", err)
			}

			// 成功获取收据，检查交易状态
			if receipt.Status == types.ReceiptStatusSuccessful {
				logger.Info().Str("txHash", txHash.Hex()).Int64("blockNumber", receipt.BlockNumber.Int64()).Msg("Transaction confirmed successfully")
			} else {
				logger.Error().Str("txHash", txHash.Hex()).Int64("blockNumber", receipt.BlockNumber.Int64()).Msg("Transaction failed")
			}
			return receipt, nil
		}
	}
}

// EstimateGas 估算交易所需 gas 数量
// 参数:
//   - client: 以太坊客户端，用于与节点交互
//   - from: 发送者地址
//   - to: 接收者地址
//   - value: 交易金额
//   - data: 交易数据
//
// 返回:
//   - uint64: 估算的 gas 数量
//   - error: 可能的错误
func EstimateGas(ctx context.Context, client *ethclient.Client, from, to common.Address, value *big.Int, data []byte) (uint64, error) {
	msg := ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: value,
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

// GetChainID
func GetChainID(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	return chainID, nil
}
