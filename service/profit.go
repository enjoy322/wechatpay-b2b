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

// ProfitService 处理分账相关调用。
type ProfitService interface {
	// ProfitSharing 请求分账。
	ProfitSharing(ctx context.Context, req types.ProfitSharingRequest, appKey string) (*types.ProfitSharingResponse, error)
	// QueryProfitSharing 查询分账订单。
	QueryProfitSharing(ctx context.Context, req types.QueryProfitSharingRequest, appKey string) (*types.QueryProfitSharingResponse, error)
	// ProfitSharingFinish 分账完结。
	ProfitSharingFinish(ctx context.Context, req types.ProfitSharingFinishRequest, appKey string) (*types.ProfitSharingFinishResponse, error)
	// ProfitSharingReturn 分账回退。
	ProfitSharingReturn(ctx context.Context, req types.ProfitSharingReturnRequest, appKey string) (*types.ProfitSharingReturnResponse, error)
	// QueryProfitSharingReturn 查询分账回退。
	QueryProfitSharingReturn(ctx context.Context, req types.QueryProfitSharingReturnRequest, appKey string) (*types.QueryProfitSharingReturnResponse, error)
	// AddProfitSharingAccount 添加分账方。
	AddProfitSharingAccount(ctx context.Context, req types.AddProfitSharingAccountRequest, appKey string) (*types.AddProfitSharingAccountResponse, error)
	// DelProfitSharingAccount 删除分账方。
	DelProfitSharingAccount(ctx context.Context, req types.DelProfitSharingAccountRequest, appKey string) (*types.DelProfitSharingAccountResponse, error)
	// QueryProfitSharingAccount 查询分账方。
	QueryProfitSharingAccount(ctx context.Context, req types.QueryProfitSharingAccountRequest, appKey string) (*types.QueryProfitSharingAccountResponse, error)
	// QueryProfitSharingRemainAmt 查询分账剩余金额。
	QueryProfitSharingRemainAmt(ctx context.Context, req types.QueryProfitSharingRemainAmtRequest, appKey string) (*types.QueryProfitSharingRemainAmtResponse, error)
}

type profitService struct {
	client *client.Client
}

const (
	profitSharingURI               = "/retail/B2b/profitsharing"
	queryProfitSharingURI          = "/retail/B2b/queryprofitsharing"
	profitSharingFinishURI         = "/retail/B2b/profitsharingfinish"
	profitSharingReturnURI         = "/retail/B2b/profitsharingreturn"
	queryProfitSharingReturn       = "/retail/B2b/queryprofitsharingreturn"
	addProfitSharingAccountURI     = "/retail/B2b/addprofitsharingaccount"
	delProfitSharingAccountURI     = "/retail/B2b/delprofitsharingaccount"
	queryProfitSharingAccountURI   = "/retail/B2b/queryprofitsharingaccount"
	queryProfitSharingRemainAmtURI = "/retail/B2b/queryprofitsharingremainamt"
)

// NewProfitService 创建分账服务。
func NewProfitService(c *client.Client) ProfitService {
	return &profitService{client: c}
}

// ProfitSharing 请求分账。
func (s *profitService) ProfitSharing(ctx context.Context, req types.ProfitSharingRequest, appKey string) (*types.ProfitSharingResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.TransactionID == "" {
		return nil, errors.New("transaction_id is required")
	}
	if req.OutOrderNo == "" {
		return nil, errors.New("out_order_no is required")
	}
	if len(req.Receivers) == 0 {
		return nil, errors.New("receivers is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(profitSharingURI, body, appKey)

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

	var out types.ProfitSharingResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// QueryProfitSharing 查询分账订单。
func (s *profitService) QueryProfitSharing(ctx context.Context, req types.QueryProfitSharingRequest, appKey string) (*types.QueryProfitSharingResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.TransactionID == "" {
		return nil, errors.New("transaction_id is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(queryProfitSharingURI, body, appKey)

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

	var out types.QueryProfitSharingResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// ProfitSharingFinish 分账完结。
func (s *profitService) ProfitSharingFinish(ctx context.Context, req types.ProfitSharingFinishRequest, appKey string) (*types.ProfitSharingFinishResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.TransactionID == "" {
		return nil, errors.New("transaction_id is required")
	}
	if req.OutOrderNo == "" {
		return nil, errors.New("out_order_no is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(profitSharingFinishURI, body, appKey)

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

	var out types.ProfitSharingFinishResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// ProfitSharingReturn 分账回退。
func (s *profitService) ProfitSharingReturn(ctx context.Context, req types.ProfitSharingReturnRequest, appKey string) (*types.ProfitSharingReturnResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.TransactionID == "" {
		return nil, errors.New("transaction_id is required")
	}
	if req.OutOrderNo == "" {
		return nil, errors.New("out_order_no is required")
	}
	if req.OutReturnNo == "" {
		return nil, errors.New("out_return_no is required")
	}
	if req.ReturnAmount <= 0 {
		return nil, errors.New("return_amount is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(profitSharingReturnURI, body, appKey)

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

	var out types.ProfitSharingReturnResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// QueryProfitSharingReturn 查询分账回退。
func (s *profitService) QueryProfitSharingReturn(ctx context.Context, req types.QueryProfitSharingReturnRequest, appKey string) (*types.QueryProfitSharingReturnResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.OutReturnNo == "" && req.ReturnID == "" {
		return nil, errors.New("out_return_no or return_id is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(queryProfitSharingReturn, body, appKey)

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

	var out types.QueryProfitSharingReturnResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// AddProfitSharingAccount 添加分账方。
func (s *profitService) AddProfitSharingAccount(ctx context.Context, req types.AddProfitSharingAccountRequest, appKey string) (*types.AddProfitSharingAccountResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.ProfitSharingRelationType == "" {
		return nil, errors.New("profit_sharing_relation_type is required")
	}
	if req.PayeeType == "" {
		return nil, errors.New("payee_type is required")
	}
	if req.PayeeID == "" {
		return nil, errors.New("payee_id is required")
	}
	if req.PayeeName == "" {
		return nil, errors.New("payee_name is required")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(addProfitSharingAccountURI, body, appKey)

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

	var out types.AddProfitSharingAccountResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// DelProfitSharingAccount 删除分账方。
func (s *profitService) DelProfitSharingAccount(ctx context.Context, req types.DelProfitSharingAccountRequest, appKey string) (*types.DelProfitSharingAccountResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.PayeeType == "" {
		return nil, errors.New("payee_type is required")
	}
	if req.PayeeID == "" {
		return nil, errors.New("payee_id is required")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(delProfitSharingAccountURI, body, appKey)

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

	var out types.DelProfitSharingAccountResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// QueryProfitSharingAccount 查询分账方。
func (s *profitService) QueryProfitSharingAccount(ctx context.Context, req types.QueryProfitSharingAccountRequest, appKey string) (*types.QueryProfitSharingAccountResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.Offset < 0 {
		return nil, errors.New("offset must be >= 0")
	}
	if req.Limit <= 0 || req.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(queryProfitSharingAccountURI, body, appKey)

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

	var out types.QueryProfitSharingAccountResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// QueryProfitSharingRemainAmt 查询分账剩余金额。
func (s *profitService) QueryProfitSharingRemainAmt(ctx context.Context, req types.QueryProfitSharingRemainAmtRequest, appKey string) (*types.QueryProfitSharingRemainAmtResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.OutTradeNo == "" {
		return nil, errors.New("out_trade_no is required")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(queryProfitSharingRemainAmtURI, body, appKey)

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

	var out types.QueryProfitSharingRemainAmtResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}
