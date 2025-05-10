package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// Parse 从toml和env中解析配置出内容
func Parse[Config any](content string) (Config, error) {
	var cfg Config
	if _, err := toml.Decode(content, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse toml: %w", err)
	}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse env: %w", err)
	}
	return cfg, nil
}

// ParseEnv 根据当前env解析配置内容，解析优先级如下
//   - toml中定义的对应env下的配置
//   - env中定义的环境变量配置
//   - toml中定义的default配置
func ParseEnv[Config any](currentEnv string, defaultConfig map[string]Config) (Config, error) {
	conf, ok := defaultConfig[currentEnv]
	if ok {
		return conf, nil
	}
	conf, ok = defaultConfig["default"]
	if !ok {
		return conf, fmt.Errorf("missing env default")
	}
	if err := env.Parse(&conf); err != nil {
		return conf, fmt.Errorf("failed to parse env: %w", err)
	}
	return conf, nil
}
