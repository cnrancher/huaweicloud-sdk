package common

import (
	"encoding/json"
	"testing"
)

const errInfo = `{"error_code":"CCE_CM.0201","error_msg":"Request body invalid","errorCode":"E.CFE.4000201","reason":"Request body invalid"}`
const errInfoV1 = `{"error":{"code":"test","message":"test message"}}`

func Test_DecodeErrorInfo(t *testing.T) {
	info := ErrorInfo{}
	if err := json.Unmarshal([]byte(errInfo), &info); err != nil {
		t.Fatal(err)
	}
	if info.Code != "CCE_CM.0201" || info.Description != "Request body invalid" {
		println(info.Code)
		println(info.Description)
		t.Fatal("decode error info not match")
	}
	info = ErrorInfo{}
	if err := json.Unmarshal([]byte(errInfoV1), &info); err != nil {
		t.Fatal(err)
	}
	if info.Code != "test" || info.Description != "test message" {
		t.Fatal("decode error info v1 not match")
	}
}
