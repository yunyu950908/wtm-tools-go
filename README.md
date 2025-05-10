# wtm-tools-go

## 挖提卖工具, 创造财富飞轮 🤪

- 支持熊链 infrared vault 挖提卖 (stake -> claim -> swap -> stake)
- 支持任意 infrared vault 单币池/双币池都可以
- 修改 `config/config.toml` 定义挖提卖任务
- 敏感内容可以通过环境变量提供
- 支持配置无限数量的挖提卖任务
- 支持 cron 表达式自由配置调度周期
- 支持配置 swap 白名单, 自由配置资产兑换, 提高财富飞轮工作效率

## TODO

- 接入借贷平台, 使用中性对冲策略
- 接入预言机, 自动平衡借贷仓位管理
- 接入 Pendle 挖提卖
- 其他