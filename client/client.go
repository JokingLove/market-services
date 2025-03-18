package client

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/JokingLove/market-services/services/rest/model"
)

var errMarketPriceHttpError = errors.New("call market price http server fail")

type MarketPriceClient interface {
	GetSupportAsset(assetName string) (bool, error)
	GetMarketPrice(assetName string) (*model.MarketPriceResponse, error)
}

type Client struct {
	client *resty.Client
}

func NewMarketPriceClient(url string) *Client {
	client := resty.New()
	client.SetBaseURL(url)
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		statusCode := response.StatusCode()
		if statusCode >= 400 {
			method := response.Request.Method
			baseUrl := response.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, baseUrl, errMarketPriceHttpError)
		}
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetSupportAsset(assetName string) (bool, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"asset_name": assetName,
	}).SetResult(model.SupportAssetResponse{}).Get("/api/v1/get_support_asset")
	if err != nil {
		return false, errors.New("get support asset fail")
	}
	rept, ok := res.Result().(*model.SupportAssetResponse)
	if !ok {
		return false, errors.New("parse result fail")
	}

	return rept.IsSupport, nil
}

func (c *Client) GetMarketPrice(assetName string) (*model.MarketPriceResponse, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"asset_name": assetName,
	}).SetResult(model.MarketPriceResponse{}).Get("/api/v1/get_market_price")
	if err != nil {
		return nil, errors.New("get support asset fail")
	}
	rept, ok := res.Result().(*model.MarketPriceResponse)
	if !ok {
		return nil, errors.New("parse result fail")
	}

	return rept, nil
}
