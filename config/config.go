package config

import (
	_ "embed"
	"wtm-tools-go/utils"
)

//go:embed config.toml
var configToml string

type Config struct {
	NodeEnv   string                     `env:"NODE_ENV"`
	Timezone  string                     `toml:"timezone"`
	Task      map[string]TaskConfig      `toml:"task"`
	Wallet    map[string]WalletConfig    `toml:"wallet"`
	RPC       map[string]RPCConfig       `toml:"rpc"`
	Oogabooga map[string]OogaboogaConfig `toml:"oogabooga"`
}

type OogaboogaConfig struct {
	BaseURL string `toml:"baseURL" env:"OOGABOOGA_BASE_URL"`
	APIKey  string `toml:"apiKey" env:"OOGABOOGA_API_KEY"`
}

type WalletConfig struct {
	PrivateKeyHead string `toml:"privateKeyHead" env:"PRIVATE_KEY_HEAD"`
	PrivateKeyTail string `toml:"privateKeyTail" env:"PRIVATE_KEY_TAIL"`
}

type RPCConfig struct {
	BerachainHTTP string `toml:"berachainHttp" env:"BERACHAIN_RPC_HTTP"`
	BerachainWS   string `toml:"berachainWs" env:"BERACHAIN_RPC_WS"`
}

type TaskConfig struct {
	Claim []ClaimConfig `toml:"claim"`
	Stake []StakeConfig `toml:"stake"`
}

type VaultProvider string

const (
	VaultProviderInfrared VaultProvider = "INFRARED"
)

type SwapProvider string

const (
	SwapProviderOogaboooga SwapProvider = "OOGABOOGA"
)

type ClaimConfig struct {
	Name         string        `toml:"name"`
	Cron         string        `toml:"cron"`
	Provider     VaultProvider `toml:"provider"`
	VaultAddress string        `toml:"vaultAddress"`
}

type StakeConfig struct {
	Name             string        `toml:"name"`
	Cron             string        `toml:"cron"`
	Provider         VaultProvider `toml:"provider"`
	VaultAddress     string        `toml:"vaultAddress"`
	MinAmountInUSD   float64       `toml:"minAmountInUSD"`
	MaxAmountInUSD   float64       `toml:"maxAmountInUSD"`
	AutoSwap         bool          `toml:"autoSwap"`
	AutoSwapAssets   []string      `toml:"autoSwapAssets"`
	AutoSwapSlippage float64       `toml:"autoSwapSlippage"`
	AutoSwapProvider SwapProvider  `toml:"autoSwapProvider"`
}

func Provide() Config {
	cfg, err := Parse[Config](configToml)
	if err != nil {
		panic(err)
	}
	return cfg
}

func ProvideWalletConfig(cfg Config) WalletConfig {
	walletConfig, err := ParseEnv(cfg.NodeEnv, cfg.Wallet)
	if err != nil {
		panic(err)
	}
	return walletConfig
}

func ProvideRPCConfig(cfg Config) RPCConfig {
	rpcConfig, err := ParseEnv(cfg.NodeEnv, cfg.RPC)
	if err != nil {
		panic(err)
	}
	return rpcConfig
}

func ProvideTaskConfig(cfg Config) TaskConfig {
	taskConfig, err := ParseEnv(cfg.NodeEnv, cfg.Task)
	if err != nil {
		panic(err)
	}
	for _, claim := range taskConfig.Claim {
		if err := utils.ValidateCronExpr(claim.Cron); err != nil {
			panic(err)
		}
	}
	for _, stake := range taskConfig.Stake {
		if err := utils.ValidateCronExpr(stake.Cron); err != nil {
			panic(err)
		}
	}
	return taskConfig
}

func ProvideOogaboogaConfig(cfg Config) OogaboogaConfig {
	oogaboogaConfig, err := ParseEnv(cfg.NodeEnv, cfg.Oogabooga)
	if err != nil {
		panic(err)
	}
	return oogaboogaConfig
}
