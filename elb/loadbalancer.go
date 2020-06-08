package elb

import (
	"context"
	"errors"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func (c *Client) GetLoadBalancers(ctx context.Context) (*common.LoadBalancerList, error) {
	rtn := common.LoadBalancerList{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("loadbalancers"),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}

func (c *Client) GetLoadBalancer(ctx context.Context, id string) (*common.LoadBalancerInfo, error) {
	if id == "" {
		return nil, errors.New("loadbalancer id is required")
	}
	rtn := common.LoadBalancerInfo{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("loadbalancers", id),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}

func (c *Client) UpdateLoadBalancer(ctx context.Context, id string, request *common.UpdatableLoadBalancerAttribute) (*common.LoadBalancerInfo, error) {
	if id == "" {
		return nil, errors.New("loadbalancer id is required")
	}
	rtn := common.LoadBalancerInfo{}
	_, err := c.DoRequest(
		ctx,
		http.MethodPut,
		c.GetURL("loadbalancers", id),
		request,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}

func (c *Client) DeleteLoadBalancer(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("loadbalancer id is required")
	}
	job := common.LoadBalancerJobInfo{}
	_, err := c.DoRequest(
		ctx,
		http.MethodDelete,
		c.GetURL("loadbalancers", id),
		nil,
		&job,
	)
	if err != nil {
		return err
	}
	_, _, err = c.WaitForELBJob(ctx, common.DefaultDuration, common.DefaultTimeout, job.JobID)
	return err
}

func (c *Client) CreateLoadBalancer(ctx context.Context, request *common.LoadBalancerRequest) (*common.LoadBalancerInfo, error) {
	lbInfo := common.LoadBalancerInfo{}
	_, err := c.DoRequest(
		ctx,
		http.MethodPost,
		c.GetURL("loadbalancers"),
		request,
		&lbInfo,
	)
	if err != nil {
		return nil, err
	}
	return c.GetLoadBalancer(ctx, lbInfo.Loadbalancer.ID)
}
