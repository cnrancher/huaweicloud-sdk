package cce

import (
	"context"
	"testing"

	"github.com/cnrancher/huaweicloud-sdk/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCCEClient(t *testing.T) {
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

	k8sClient, err := GetClusterClient(&list.Items[0], cceClient)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = k8sClient.CoreV1Client.Namespaces().List(metav1.ListOptions{}); err != nil {
		t.Fatal(err)
	}
}
