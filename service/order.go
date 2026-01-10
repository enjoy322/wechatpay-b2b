package service

import (
	"context"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/model"
)

// OrderService 处理订单查询、关单相关调用。
type OrderService struct {
	Client *client.Client
}

// CloseOrderRequest 关闭订单请求参数。
type CloseOrderRequest struct {
	Mchid      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

// CloseOrderResponse 关闭订单返回参数。
type CloseOrderResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetOrderRequest 查询订单请求参数。
type GetOrderRequest struct {
	Mchid      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

// GetOrderResponse 查询订单返回参数。
type GetOrderResponse struct {
	AppID        string             `json:"appid"`
	Mchid        string             `json:"mchid"`
	OutTradeNo   string             `json:"out_trade_no"`
	OrderID      string             `json:"order_id"`
	PayStatus    model.PayStatus    `json:"pay_status"`
	PayTime      string             `json:"pay_time,omitempty"`
	Attach       string             `json:"attach,omitempty"`
	RefundStatus model.RefundStatus `json:"refund_status,omitempty"`
	ErrCode      int                `json:"errcode,omitempty"`
	ErrMsg       string             `json:"errmsg,omitempty"`
}

func (s *OrderService) CloseOrder(ctx context.Context, req CloseOrderRequest) (*CloseOrderResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// GetOrder 查询订单。
func (s *OrderService) GetOrder(ctx context.Context, req GetOrderRequest) (*GetOrderResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}
