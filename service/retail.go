package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// RetailService 处理门店助手（wxa/business）相关接口调用。
type RetailService interface {
	// BatchCreateRetail 预录入门店信息。
	BatchCreateRetail(ctx context.Context, req types.BatchCreateRetailRequest) (*types.BatchCreateRetailResponse, error)
	// GetRetailInfo 查询门店信息。
	GetRetailInfo(ctx context.Context, req types.GetRetailInfoRequest) (*types.GetRetailInfoResponse, error)
	// GetRetailOpenIDList 全量授权门店查询。
	GetRetailOpenIDList(ctx context.Context, req types.GetRetailOpenIDListRequest) (*types.GetRetailOpenIDListResponse, error)
}

type retailService struct {
	client *client.Client
}

const (
	batchCreateRetailURI   = "/wxa/business/batchcreateretail"
	getRetailInfoURI       = "/wxa/business/getretailinfo"
	getRetailOpenIDListURI = "/wxa/business/getretailopenidlist"
)

// NewRetailService 创建门店助手服务。
func NewRetailService(c *client.Client) RetailService {
	return &retailService{client: c}
}

// BatchCreateRetail 预录入门店信息。
func (s *retailService) BatchCreateRetail(ctx context.Context, req types.BatchCreateRetailRequest) (*types.BatchCreateRetailResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if len(req.RetailInfoList) == 0 {
		return nil, errors.New("retail_info_list is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(batchCreateRetailURI)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.BatchCreateRetailResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetRetailInfo 查询门店信息。
func (s *retailService) GetRetailInfo(ctx context.Context, req types.GetRetailInfoRequest) (*types.GetRetailInfoResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.OpenID == "" && req.MobilePhone == "" {
		return nil, errors.New("openid or mobile_phone is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(getRetailInfoURI)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.GetRetailInfoResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetRetailOpenIDList 全量授权门店查询。
func (s *retailService) GetRetailOpenIDList(ctx context.Context, req types.GetRetailOpenIDListRequest) (*types.GetRetailOpenIDListResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Limit <= 0 || req.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(getRetailOpenIDListURI)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.GetRetailOpenIDListResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}
