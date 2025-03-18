package routes

import (
	"fmt"
	"github.com/JokingLove/market-services/services/rest/model"
	"net/http"
)

func (h Routes) GetSupportAsset(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("asset")
	cr := &model.SupportAssetRequest{
		AssetName: assetName,
	}
	supRet, err := h.service.GetSupportAsset(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, supRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response:", "err", err)
	}
}

func (h Routes) GetMarketPrice(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("asset")
	cr := &model.MarketPriceRequest{
		AssetName: assetName,
	}

	supRet, err := h.service.GetMarketPrice(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, supRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response:", "err", err)
	}
}
