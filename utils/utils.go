package utils

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"wtm-tools-go/internal/sc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	MAX_UINT256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
)

// FormatUnits 将大整数按指定单位和精度转换为字符串表示
// value: 大整数值（例如以 wei 为单位）
// decimals: 单位的小数位数（例如 ether 为 18）
// precision: 输出结果保留的小数位数
func FormatUnits(value *big.Int, decimals int, precision int) string {
	// 创建一个除数，10 的 decimals 次方
	divisor := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	// 计算整数部分
	quotient := big.NewInt(0).Div(value, divisor)

	// 计算余数（用于小数部分）
	remainder := big.NewInt(0).Mod(value, divisor)

	// 将余数转换为字符串，并在前补零以达到 decimals 位数
	remainderStr := remainder.String()
	remainderStr = strings.Repeat("0", decimals-len(remainderStr)) + remainderStr

	// 截取指定精度的小数部分
	if precision > 0 {
		if len(remainderStr) > precision {
			remainderStr = remainderStr[:precision]
		} else {
			remainderStr = remainderStr + strings.Repeat("0", precision-len(remainderStr))
		}
		// 去除尾部的零
		remainderStr = strings.TrimRight(remainderStr, "0")
		if remainderStr == "" {
			return quotient.String()
		}
		return fmt.Sprintf("%s.%s", quotient.String(), remainderStr)
	}

	return quotient.String()
}

type TokenBalance struct {
	Balance  *big.Int
	Amount   float64 // human readable amount
	Decimals uint8
}

func CheckBalance(ctx context.Context, tokenAddress string, walletAddress common.Address, client *ethclient.Client) (*TokenBalance, error) {
	erc20, err := sc.NewERC20(common.HexToAddress(tokenAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create erc20: %w", err)
	}
	balance, err := erc20.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	decimals, err := erc20.Decimals(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get decimals: %w", err)
	}
	amountStr := FormatUnits(balance, int(decimals), 6)
	amountFloat, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	return &TokenBalance{
		Balance:  balance,
		Amount:   amountFloat,
		Decimals: decimals,
	}, nil
}

// f: 输入的float64
// decimal: 要扩大到的小数位数倍数，比如18
func FloatToBigIntWithDecimal(f float64, decimal uint8) *big.Int {
	// 构造 10^decimal
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	scaleFloat := new(big.Float).SetInt(scale)

	// 计算 f * 10^decimal
	resultFloat := new(big.Float).Mul(new(big.Float).SetFloat64(f), scaleFloat)

	// 截断小数部分
	resultInt := new(big.Int)
	resultFloat.Int(resultInt)
	return resultInt
}

// DecodeHexToBytes 自动判断字符串类型并返回对应字节数组
func DecodeHexToBytes(input string) ([]byte, error) {
	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		data := input[2:]
		// odd length hex string is invalid
		if len(data)%2 != 0 {
			data = "0" + data
		}
		b, err := hex.DecodeString(data)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	// 直接转字节
	return []byte(input), nil
}
