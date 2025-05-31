package public

import "github.com/anhnguyenbk/go-okx/rest/api"

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsParam struct {
	InstType   string `url:"instType"`
	InstId     string `url:"instId,omitempty"`
	Uly        string `url:"uly,omitempty"`
	InstFamily string `url:"instFamily,omitempty"`
}

type GetInstrumentsResponse struct {
	api.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	InstType       string `json:"instType"`       // Instrument type, e.g., SPOT, FUTURES
	InstId         string `json:"instId"`         // Instrument ID, e.g., BTC-USDT
	Uly            string `json:"uly"`            // Underlying, e.g., BTC-USD
	InstFamily     string `json:"instFamily"`     // Instrument family, e.g., BTC-USD
	Category       string `json:"category"`       // Currency category (deprecated)
	BaseCcy        string `json:"baseCcy"`        // Base currency, e.g., BTC
	QuoteCcy       string `json:"quoteCcy"`       // Quote currency, e.g., USDT
	SettleCcy      string `json:"settleCcy"`      // Settlement currency
	CtVal          string `json:"ctVal"`          // Contract value
	CtMult         string `json:"ctMult"`         // Contract multiplier
	CtValCcy       string `json:"ctValCcy"`       // Contract value currency
	OptType        string `json:"optType"`        // Option type, C or P
	Stk            string `json:"stk"`            // Strike price
	ListTime       string `json:"listTime"`       // Listing time in milliseconds
	ExpTime        string `json:"expTime"`        // Expiry time
	Lever          string `json:"lever"`          // Maximum leverage
	TickSz         string `json:"tickSz"`         // Tick size
	LotSz          string `json:"lotSz"`          // Lot size
	MinSz          string `json:"minSz"`          // Minimum order size
	CtType         string `json:"ctType"`         // Contract type, linear or inverse
	Alias          string `json:"alias"`          // Contract alias
	State          string `json:"state"`          // Instrument state, e.g., live
	MaxLmtSz       string `json:"maxLmtSz"`       // Maximum limit order size
	MaxMktSz       string `json:"maxMktSz"`       // Maximum market order size
	MaxLmtAmt      string `json:"maxLmtAmt"`      // Maximum limit order amount
	MaxMktAmt      string `json:"maxMktAmt"`      // Maximum market order amount
	MaxTwapSz      string `json:"maxTwapSz"`      // Maximum TWAP order size
	MaxIcebergSz   string `json:"maxIcebergSz"`   // Maximum iceberg order size
	MaxTriggerSz   string `json:"maxTriggerSz"`   // Maximum trigger order size
	MaxStopSz      string `json:"maxStopSz"`      // Maximum stop order size
	AuctionEndTime string `json:"auctionEndTime"` // Auction end time
	RuleType       string `json:"ruleType"`       // Rule type
}
