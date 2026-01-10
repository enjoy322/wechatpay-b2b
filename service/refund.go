package service

import (
	"context"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/model"
)

// RefundService 处理退款申请与退款查询。
type RefundService struct {
	Client *client.Client
}

// RefundRequest 退款请求参数。
type RefundRequest struct {
	Mchid        string `json:"mchid"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OrderID      string `json:"order_id,omitempty"`
	OutRefundNo  string `json:"out_refund_no"`
	RefundAmount int64  `json:"refund_amount"`
	RefundFrom   int32  `json:"refund_from"`
	RefundReason int32  `json:"refund_reason,omitempty"`
	Description  string `json:"description,omitempty"`
}

// RefundResponse 退款返回参数。
type RefundResponse struct {
	RefundID    string `json:"refund_id"`
	OutRefundNo string `json:"out_refund_no"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// GetRefundRequest 查询退款请求参数。
type GetRefundRequest struct {
	Mchid       string `json:"mchid"`
	OutRefundNo string `json:"out_refund_no,omitempty"`
	RefundID    string `json:"refund_id,omitempty"`
}

// GetRefundResponse 查询退款返回参数。
type GetRefundResponse struct {
	RefundID    string             `json:"refund_id"`
	OutRefundNo string             `json:"out_refund_no"`
	OrderID     string             `json:"order_id"`
	OutTradeNo  string             `json:"out_trade_no"`
	Status      model.RefundStatus `json:"status,omitempty"`
	ErrCode     int                `json:"errcode"`
	ErrMsg      string             `json:"errmsg"`
}

// CreateRefund 发起退款。
func (s *RefundService) CreateRefund(ctx context.Context, req RefundRequest) (*RefundResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// GetRefund 查询退款。
func (s *RefundService) GetRefund(ctx context.Context, req GetRefundRequest) (*GetRefundResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}
