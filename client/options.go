package client

import (
	"errors"
)

// Options Client 初始化参数。
type Options struct {
	// AccessToken 接口调用凭证（可由业务侧定时刷新并更新到 Client 上）。
	AccessToken string
}

// NewClient 创建一个可复用的微信 API Client。
// 注意：业务侧可定时刷新 access_token，并更新到返回的 Client 上。
func NewClient(opts Options) (*Client, error) {
	if opts.AccessToken == "" {
		return nil, errors.New("accessToken is empty")
	}
	c := &Client{}
	c.accessToken = opts.AccessToken
	c.baseURL = "https://api.weixin.qq.com"

	return c, nil
}
