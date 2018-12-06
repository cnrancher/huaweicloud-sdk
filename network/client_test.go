package network

import (
	"testing"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func Test_GetURL(t *testing.T) {
	baseClient := common.NewClient("abcd", "def", "myhuawei.com", "cn-north-1", "test")
	c := NewClient(baseClient)
	if c.GetAPIEndpointFunc() != "https://vpc.cn-north-1.myhuawei.com" {
		t.Fatal("api endpoint is not valid", c.GetAPIEndpointFunc())
	}
	if c.GetURL("vpcs", "1234") != "https://vpc.cn-north-1.myhuawei.com/v1/test/vpcs/1234" {
		t.Fatal("get url is not valid", c.GetURL("vpcs", "1234"))
	}
}
