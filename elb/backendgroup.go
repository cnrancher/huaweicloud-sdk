package elb

import (
	"context"
	"errors"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func (c *Client) AddBackendGroup(ctx context.Context, backend common.ELBBackendGroupRequest) (common.ELBBackendGroupDetails, error) {
	if backend.Pool.ListenerID == "" && backend.Pool.LoadbalancerID == "" {
		return common.ELBBackendGroupDetails{}, errors.New("you hava to specific one of ListenerID and LoadbalancerID")
	}
	ret := common.ELBBackendGroupDetails{}
	_, err := c.DoRequest(
		ctx,
		http.MethodPost,
		c.GetURL("pools"),
		backend,
		&ret,
	)
	if err != nil {
		return common.ELBBackendGroupDetails{}, err
	}
	return ret, nil
}

func (c *Client) RemoveBackendGroup(ctx context.Context, poolID string) error {
	if poolID == "" {
		return errors.New("[RemoveBackendGroup]pool id is required")
	}
	_, err := c.DoRequest(
		ctx,
		http.MethodDelete,
		c.GetURL("pools", poolID),
		nil,
		nil,
	)
	return err

}

func (c *Client) GetBackendGroup(ctx context.Context, poolID string) (common.ELBBackendGroupDetails, error) {
	if poolID == "" {
		return common.ELBBackendGroupDetails{}, errors.New("[GetBackendGroup]pool id is required")
	}
	rtn := common.ELBBackendGroupDetails{}
	if _, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("pools", poolID),
		nil,
		&rtn,
	); err != nil {
		return common.ELBBackendGroupDetails{}, err
	}
	return rtn, nil
}

func (c *Client) RemoveHealthmonitors(ctx context.Context, healthmonitorID string) error {
	if healthmonitorID == "" {
		return errors.New("[RemoveHealthmonitors]healthmonitor id is required")
	}
	_, err := c.DoRequest(
		ctx,
		http.MethodDelete,
		c.GetURL("healthmonitors", healthmonitorID),
		nil,
		nil,
	)
	return err
}
