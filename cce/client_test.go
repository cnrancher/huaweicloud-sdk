package cce

import (
	"context"
	"testing"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func Test_CCEGetURL(t *testing.T) {
	baseClient := common.NewClient("abcd", "def", "myhuawei.com", "cn-north-1", "test")
	c := NewClient(baseClient)
	if c.GetAPIEndpointFunc() != "https://cce.cn-north-1.myhuawei.com" {
		t.Fatal("api endpoint is not valid", c.GetAPIEndpointFunc())
	}
	if c.GetURL("clusters", "1234", "nodes") != "https://cce.cn-north-1.myhuawei.com/api/v3/projects/test/clusters/1234/nodes" {
		t.Fatal("get url is not valid", c.GetURL("clusters", "1234", "nodes"))
	}
	if c.GetSignerServiceName() != "cce" {
		t.Fatalf("signer service name is not valid, signer:%s", c.GetSignerServiceName())
	}
}

func Test_CCEClient(t *testing.T) {
	baseClient, err := common.GetBaseClientFromENV()
	if err != nil {
		t.Skip(err)
	}
	cceClient := NewClient(baseClient)
	_, err = cceClient.GetClusters(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
