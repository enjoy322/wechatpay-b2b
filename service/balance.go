package service

import (
	"context"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/model"
)

// BalanceService 处理余额查询与提现流程。
type BalanceService struct {
	Client *client.Client
}

// BalanceRequest 余额查询请求参数。
type BalanceRequest struct {
	Mchid string `json:"mchid"`
}

// BalanceResponse 余额查询返回参数。
type BalanceResponse struct {
	BalanceList []model.BalanceInfo `json:"balance_list"`
	ErrCode     int                 `json:"errcode"`
	ErrMsg      string              `json:"errmsg"`
}

// WithdrawRequest 提现请求参数。
type WithdrawRequest struct {
	Mchid          string `json:"mchid"`
	WithdrawAmount int64  `json:"withdraw_amount"`
	OutWithdrawNo  string `json:"out_withdraw_no"`
}

// WithdrawResponse 提现返回参数。
type WithdrawResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryWithdrawRequest 查询提现状态请求参数。
type QueryWithdrawRequest struct {
	Mchid         string `json:"mchid"`
	OutWithdrawNo string `json:"out_withdraw_no"`
}

// QueryWithdrawResponse 查询提现状态返回参数。
type QueryWithdrawResponse struct {
	OutWithdrawNo  string               `json:"out_withdraw_no"`
	WithdrawAmount int64                `json:"withdraw_amount"`
	Status         model.WithdrawStatus `json:"status"`
	FailReason     string               `json:"fail_reason,omitempty"`
	ErrCode        int                  `json:"errcode"`
	ErrMsg         string               `json:"errmsg"`
}

// GetBalance 查询账户余额。
func (s *BalanceService) GetBalance(ctx context.Context, req BalanceRequest) (*BalanceResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// Withdraw 发起提现。
func (s *BalanceService) Withdraw(ctx context.Context, req WithdrawRequest) (*WithdrawResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// QueryWithdraw 查询提现状态。
func (s *BalanceService) QueryWithdraw(ctx context.Context, req QueryWithdrawRequest) (*QueryWithdrawResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}
