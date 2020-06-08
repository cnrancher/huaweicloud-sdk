package network

import (
	"context"
	"errors"
	"net/http"

	"github.com/cnrancher/huaweicloud-sdk/common"
)

func (c *Client) GetPrivateIP(ctx context.Context, privateipID string) (*common.PrivateIpResp, error) {
	if privateipID == "" {
		return nil, errors.New("[GetPrivateIP]private ip id is required")
	}
	rtn := common.PrivateIpResp{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("privateips", privateipID),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}

func (c *Client) GetPrivateIPList(ctx context.Context, subnetID string) (*common.PrivateIpListResp, error) {
	if subnetID == "" {
		return nil, errors.New("[GetPrivateIPList]subnet id is required")
	}
	rtn := common.PrivateIpListResp{}
	_, err := c.DoRequest(
		ctx,
		http.MethodGet,
		c.GetURL("subnets", subnetID, "privateips"),
		nil,
		&rtn,
	)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
