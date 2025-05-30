package trade

import (
	"encoding/json"
	"testing"

	"github.com/anhnguyenbk/go-okx/rest"
)

func TestErrorResponse(t *testing.T) {
	var resp rest.OKXErrorResp
	dataJson := `{"code":"1","data":[{"clOrdId":"","ordId":"","sCode":"51121","sMsg":"Order quantity must be a multiple of the lot size.","tag":"","ts":"1748617662056"}],"inTime":"1748617662056837","msg":"All operations failed","outTime":"1748617662056989"}`

	data := []byte(dataJson)
	json.Unmarshal(data, &resp)
	t.Log("error resp", resp)

	result := resp.Error()
	expected := `code: 1, message: All operations failed, data: [{"clOrdId":"","ordId":"","sCode":"51121","sMsg":"Order quantity must be a multiple of the lot size.","tag":"","ts":"1748617662056"}]`
	if result != expected {
		t.Errorf("TestErrorResponse result=%s; want=%s", result, expected)
	}
}
