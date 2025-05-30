package rest

import (
	"encoding/json"
	"fmt"
)

type OKXError struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
	Tag     string `json:"tag"`
	Ts      string `json:"ts"`
}

type OKXErrorResp struct {
	Code    string     `json:"code"`
	Message string     `json:"msg"`
	Data    []OKXError `json:"data"`
}

var _ error = (*OKXErrorResp)(nil)

func (e OKXErrorResp) Error() string {
	jsonBytes, err := json.Marshal(e.Data)
	if err != nil {
		return "error when marshaling response data: " + err.Error()
	}

	jsonStr := string(jsonBytes)
	return fmt.Sprintf("code: %s, message: %s, data: %s", e.Code, e.Message, jsonStr)
}
