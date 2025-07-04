package api

import (
	"encoding/json"
	"log"
)

type IResponse interface {
	GetCode() string
	GetMessage() string
	GetData() interface{}
	GetDataAsString() string
	IsOk() bool
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r Response) GetCode() string {
	return r.Code
}

func (r Response) GetMessage() string {
	return r.Message
}

func (r Response) GetData() interface{} {
	return r.Data
}

func (r Response) GetDataAsString() string {
	jsonBytes, err := json.Marshal(r.Data)
	if err != nil {
		log.Printf("Marshal okx resp[data] has error %v", err)
		return ""
	}

	jsonString := string(jsonBytes)
	return jsonString
}

func (r Response) IsOk() bool {
	return r.Code == "0"
}
