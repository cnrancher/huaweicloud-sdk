package cce

import (
	"context"
	"testing"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func Test_MasterIP(t *testing.T) {
	baseClient, err := common.GetBaseClientFromENV()
	if err != nil {
		t.Skip(err)
	}
	cceClient := NewClient(baseClient)
	root := context.Background()
	if _, err := cceClient.AddMasterIP(root, "32f63655-f924-11e8-978a-0255ac101f1e", &common.CCEClusterIPBindInfo{
		Spec: common.BindInfoSpec{
			Action: "bind",
			ActionSpec: &common.BindActionSpec{
				ID: "99a813d5-29ba-42ff-88e6-59a01515c3c2",
			},
			ElasticIP: "49.4.122.234",
		},
	}); err != nil {
		t.Fatal(err)
	}
}
