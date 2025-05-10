package oogabooga

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"sync"
	"time"
	"wtm-tools-go/config"
)

// Client represents the OogaBooga API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// singleton instance of the client
var instance *Client
var once sync.Once

// GetClient returns the singleton instance of the OogaBooga client
func GetClient(cfg config.OogaboogaConfig) *Client {
	once.Do(func() {
		instance = NewClient(cfg)
	})
	return instance
}

// NewClient creates a new OogaBooga API client
func NewClient(cfg config.OogaboogaConfig) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		apiKey: cfg.APIKey,
	}
}

// GetTokenPrices fetches token prices in USD
func (c *Client) GetTokenPrices(ctx context.Context) (GetTokenPricesResponse, error) {
	var response GetTokenPricesResponse
	err := c.get(ctx, "v1/prices?currency=USD", &response)
	return response, err
}

// GetTokenList fetches the list of available tokens
func (c *Client) GetTokenList(ctx context.Context) (GetTokenListResponse, error) {
	var response GetTokenListResponse
	err := c.get(ctx, "v1/tokens", &response)
	return response, err
}

// GetLiquiditySources fetches available liquidity sources
func (c *Client) GetLiquiditySources(ctx context.Context) (GetLiquiditySourcesResponse, error) {
	var response GetLiquiditySourcesResponse
	err := c.get(ctx, "v1/liquidity-sources", &response)
	return response, err
}

// GetAllowance fetches the allowance for a token from a specific address
func (c *Client) GetAllowance(ctx context.Context, token, from string) (GetAllowanceResponse, error) {
	// Native token does not require approvals for allowance
	if token == "0x0000000000000000000000000000000000000000" {
		maxUint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
		return GetAllowanceResponse{Allowance: maxUint256.String()}, nil
	}

	params := url.Values{}
	params.Add("token", token)
	params.Add("from", from)

	var response GetAllowanceResponse
	err := c.get(ctx, "v1/approve/allowance?"+params.Encode(), &response)
	return response, err
}

// GetApproveAllowanceTx generates a transaction to approve token allowance
func (c *Client) GetApproveAllowanceTx(ctx context.Context, token string, amount *big.Int) (GetApproveAllowanceTxResponse, error) {
	params := url.Values{}
	params.Add("token", token)
	params.Add("amount", amount.String())

	var response GetApproveAllowanceTxResponse
	err := c.get(ctx, "v1/approve?"+params.Encode(), &response)
	return response, err
}

// GetSwapTx generates a transaction for swapping tokens
func (c *Client) GetSwapTx(ctx context.Context, swapParams SwapParams) (GetSwapTxResponse, error) {
	params := url.Values{}
	params.Add("tokenIn", swapParams.TokenIn)
	params.Add("tokenOut", swapParams.TokenOut)
	params.Add("amount", swapParams.Amount)
	params.Add("slippage", fmt.Sprintf("%f", swapParams.Slippage))

	if swapParams.To != "" {
		params.Add("to", swapParams.To)
	}

	var response GetSwapTxResponse
	err := c.get(ctx, "v1/swap?"+params.Encode(), &response)
	return response, err
}

// get performs a GET request to the API
func (c *Client) get(ctx context.Context, path string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", c.baseURL, path), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
