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
	list, err := cceClient.GetClusters(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) == 0 {
		t.Skip()
	}
	clusterID := list.Items[0].MetaData.UID
	originDescription := list.Items[0].Spec.Description
	rtn, err := cceClient.UpdateCluster(context.Background(), clusterID, &common.UpdateCluster{Spec: common.UpdateInfo{Description: "test"}})
	if err != nil {
		t.Fatal(err)
	}
	if rtn.Spec.Description != "test" {
		t.Fatal("update cluster description fail")
	}
	_, err = cceClient.UpdateCluster(context.Background(), clusterID, &common.UpdateCluster{Spec: common.UpdateInfo{Description: originDescription}})
	if err != nil {
		t.Fatal(err)
	}
}
