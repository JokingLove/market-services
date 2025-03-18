package service

import (
	"github.com/JokingLove/market-services/database"
	"github.com/JokingLove/market-services/services/rest/model"
)

type RestService interface {
	GetSupportAsset(*model.SupportAssetRequest) (*model.SupportAssetResponse, error)
	GetMarketPrice(*model.MarketPriceRequest) (*model.MarketPriceResponse, error)
}

type HandleSer struct {
	v                    *Validator
	marketPriceView      database.MarketPriceView
	officialCoinRateView database.OfficialCoinRateView
}

func NewHandleService(v *Validator, marketPriceView database.MarketPriceView, officialView database.OfficialCoinRateView) RestService {
	return &HandleSer{
		v:                    v,
		marketPriceView:      marketPriceView,
		officialCoinRateView: officialView,
	}
}

func (h HandleSer) GetSupportAsset(request *model.SupportAssetRequest) (*model.SupportAssetResponse, error) {
	return &model.SupportAssetResponse{
		ReturnCode: 100,
		Message:    "get support asset success",
		IsSupport:  true,
	}, nil
}

func (h HandleSer) GetMarketPrice(request *model.MarketPriceRequest) (*model.MarketPriceResponse, error) {
	assetPriceList, err := h.marketPriceView.QueryMarketPriceByAsset(request.AssetName)
	if err != nil {
		return nil, err
	}
	var marketPriceList []model.MarketPrice
	for _, assetPrice := range assetPriceList {
		mpItem := model.MarketPrice{
			AssetName:   assetPrice.AssetName,
			AssetPrice:  assetPrice.PriceUsdt,
			AssetVolume: assetPrice.Volume,
			AssetRate:   assetPrice.Rate,
		}
		marketPriceList = append(marketPriceList, mpItem)
	}

	ocrList, err := h.officialCoinRateView.QueryOfficialCoinRateByAsset(request.AssetName)
	if err != nil {
		return nil, err
	}
	var officialCoinRateList []model.OfficialCoinRate
	for _, ocrItem := range ocrList {
		officialCoinRateItem := model.OfficialCoinRate{
			Name: ocrItem.AssetName,
			Rate: ocrItem.Price,
		}
		officialCoinRateList = append(officialCoinRateList, officialCoinRateItem)
	}
	return &model.MarketPriceResponse{
		ReturnCode:           100,
		Message:              "get market price success",
		MarketPriceList:      marketPriceList,
		OfficialCoinRateList: officialCoinRateList,
	}, nil
}
