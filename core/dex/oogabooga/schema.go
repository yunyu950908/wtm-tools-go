package oogabooga

// TokenPrice represents price information for a token
type TokenPrice struct {
	Address string  `json:"address"`
	Price   float64 `json:"price"`
}

// GetTokenPricesResponse is an array of TokenPrice
type GetTokenPricesResponse []TokenPrice

// Token represents token information
type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
	TokenURI string `json:"tokenURI"`
}

// GetTokenListResponse is an array of Token
type GetTokenListResponse []Token

// LiquiditySource represents a liquidity source as a string
type LiquiditySource string

// GetLiquiditySourcesResponse is an array of LiquiditySource
type GetLiquiditySourcesResponse []LiquiditySource

// GetAllowanceResponse represents the response for an allowance query
type GetAllowanceResponse struct {
	Allowance string `json:"allowance"`
}

// Transaction represents a transaction with destination and data
type Transaction struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

// GetApproveAllowanceTxResponse represents the response for an approve allowance transaction
type GetApproveAllowanceTxResponse struct {
	Tx Transaction `json:"tx"`
}

// SwapParams represents parameters for a swap operation
type SwapParams struct {
	TokenIn  string  `json:"tokenIn"`
	TokenOut string  `json:"tokenOut"`
	Amount   string  `json:"amount"`
	To       string  `json:"to,omitempty"`
	Slippage float64 `json:"slippage"`
}

// GetSwapTxResponse represents the response for a swap transaction
type GetSwapTxResponse struct {
	Tx Transaction `json:"tx"`
}
