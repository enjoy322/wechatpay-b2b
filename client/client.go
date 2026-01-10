package client

import (
	"bytes"
	"context"
	"net/http"
)

// TokenProvider 用于获取接口调用所需的 access_token。
type TokenProvider interface {
	AccessToken(ctx context.Context) (string, error)
}

// Client 封装访问微信 API 的 HTTP 客户端。
type Client struct {
	BaseURL       string
	TokenProvider TokenProvider
	HTTPClient    *http.Client
}

// Do 向指定 uri（路径，不含完整域名）发起 HTTP 请求。
func (c *Client) Do(ctx context.Context, method, uri string, body []byte) (*http.Response, error) {
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+uri, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return httpClient.Do(req)
}
