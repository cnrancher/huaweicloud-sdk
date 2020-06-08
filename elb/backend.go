package elb

import (
	"context"
	"errors"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func (c *Client) AddBackend(ctx context.Context, poolID string, backend common.ELBBackendRequest) (common.ELBBackendResponce, error) {
	if poolID == "" {
		return common.ELBBackendResponce{}, errors.New("[AddBackend]pool id is required")
	}
	rtn := common.ELBBackendResponce{}
	if _, err := c.DoRequest(
		ctx,
		http.MethodPost,
		c.GetURL("pools", poolID, "members"),
		backend,
		&rtn,
	); err != nil {
		return common.ELBBackendResponce{}, err
	}
	return rtn, nil
}

func (c *Client) RemoveBackend(ctx context.Context, poolID string, memberID string) error {
	if poolID == "" || memberID == "" {
		return errors.New("pool id and member id is both required")
	}
	_, err := c.DoRequest(
		ctx,
		http.MethodDelete,
		c.GetURL("pools", poolID, "members", memberID),
		nil,
		nil,
	)
	return err
}

func (c *Client) GetBackend(ctx context.Context, poolID string, memberID string) (common.ELBBackendResponce, error) {
	if poolID == "" || memberID == "" {
		return common.ELBBackendResponce{}, errors.New("pool id and member id is both required")
	}
	rtn := common.ELBBackendResponce{}
	if _, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("pools", poolID, "members", memberID),
		nil,
		&rtn,
	); err != nil {
		return common.ELBBackendResponce{}, err
	}
	return rtn, nil
}
