package account

import "github.com/anhnguyenbk/go-okx/rest/api"

func NewGetPositionsHistory(param *GetPositionsHistoryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/positions-history",
		Method: api.MethodGet,
		Param:  param,
	}, &GetPositionsHistoryResponse{}
}

type GetPositionsHistoryParam struct {
	InstType string `url:"instType,omitempty"`
	InstId   string `url:"instId,omitempty"`
	MgnMode  string `url:"mgnMode,omitempty"`
	Type     string `url:"type,omitempty"`
	PosId    string `url:"posId,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Limit    string `url:"limit,omitempty"`
}

type GetPositionsHistoryResponse struct {
	api.Response
	Data []PositionHistory `json:"data"`
}

type PositionHistory struct {
	CTime          string `json:"cTime"`
	Ccy            string `json:"ccy"`
	CloseAvgPx     string `json:"closeAvgPx"`
	CloseTotalPos  string `json:"closeTotalPos"`
	InstId         string `json:"instId"`
	InstType       string `json:"instType"`
	Lever          string `json:"lever"`
	MgnMode        string `json:"mgnMode"`
	OpenAvgPx      string `json:"openAvgPx"`
	OpenMaxPos     string `json:"openMaxPos"`
	RealizedPnl    string `json:"realizedPnl"`
	Fee            string `json:"fee"`
	FundingFee     string `json:"fundingFee"`
	LiqPenalty     string `json:"liqPenalty"`
	Pnl            string `json:"pnl"`
	PnlRatio       string `json:"pnlRatio"`
	PosId          string `json:"posId"`
	PosSide        string `json:"posSide"`
	Direction      string `json:"direction"`
	TriggerPx      string `json:"triggerPx"`
	Type           string `json:"type"`
	UTime          string `json:"uTime"`
	Uly            string `json:"uly"`
	NonSettleAvgPx string `json:"nonSettleAvgPx"`
	SettledPnl     string `json:"settledPnl"`
}
