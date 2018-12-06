package elb

import (
	"context"
	"testing"

	"github.com/cnrancher/cce-sdk/common"
	"github.com/cnrancher/cce-sdk/network"
	"github.com/sirupsen/logrus"
)

func Test_ELBClient(t *testing.T) {
	t.Skip()
	baseClient, err := common.GetBaseClientFromENV()
	if err != nil {
		t.Skip(err)
	}
	elbClient := NewClient(baseClient)
	networkClient := network.NewClient(baseClient)
	list, err := elbClient.GetLoadBalancers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	logrus.Debugf("%#v\n", *list)
	root := context.Background()

	subnets, err := networkClient.GetSubnets(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(subnets.Subnets) == 0 {
		t.Skip("need vpc and subnet to test")
	}

	var vpcID, subnetID string

	for _, subnet := range subnets.Subnets {
		vpcID = subnet.VpcID
		subnetID = subnet.ID
	}

	rtn, err := elbClient.CreateLoadBalancer(root, &common.LoadBalancerRequest{
		AvailableZone: "cn-north-1a",
		ChargeMode:    "traffic",
		EIPType:       "5_bgp",
		TenantID:      baseClient.ProjectID,
		LoadBalancerCommonInfo: common.LoadBalancerCommonInfo{
			Type:        "External",
			VIPSubnetID: subnetID,
			VpcID:       vpcID,
		},
		UpdatableLoadBalancerAttribute: common.UpdatableLoadBalancerAttribute{
			AdminStateUp: 1,
			Bandwidth:    10,
			Name:         "sdk-test",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	logrus.Debugf("%#v\n", *rtn)

	listener, err := elbClient.CreateListener(root, &common.ELBListenerRequest{
		ELBListenerCommon: common.ELBListenerCommon{
			LoadbalancerID:  rtn.ID,
			Protocol:        "TCP",
			BackendProtocol: "TCP",
			SessionSticky:   true,
		},
		UpdatableELBListenerAttribute: common.UpdatableELBListenerAttribute{
			Port:        8080,
			BackendPort: 8080,
			Name:        "sdk-test-8080",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := elbClient.DeleteListener(root, listener.ID); err != nil {
		t.Fatal(err)
	}
	if err := elbClient.DeleteLoadBalancer(root, rtn.ID); err != nil {
		t.Fatal(err)
	}
}

func Test_deleteLB(t *testing.T) {
	t.Skip()
	baseClient, err := common.GetBaseClientFromENV()
	if err != nil {
		t.Skip(err)
	}
	root := context.Background()
	elbClient := NewClient(baseClient)
	lbs, err := elbClient.GetLoadBalancers(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, lb := range lbs.LoadBalancers {
		println(lb.ID)
		if err := elbClient.DeleteLoadBalancer(root, lb.ID); err != nil {
			logrus.Error(err)
		}
	}
}
