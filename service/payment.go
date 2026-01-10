package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/model"
	"github.com/enjoy322/wechatpay-b2b/signer"
)

// PaymentService 负责构建小程序支付相关参数。
type PaymentService struct {
	Client *client.Client
}

// CommonPaymentSignData 微信支付 signData 结构（普通单）。
type CommonPaymentSignData struct {
	Mchid        string             `json:"mchid"`
	OutTradeNo   string             `json:"out_trade_no"`
	Description  string             `json:"description"`
	Amount       model.Amount       `json:"amount"`
	Attach       string             `json:"attach,omitempty"`
	ProductInfo  *model.ProductInfo `json:"product_info,omitempty"`
	DeliveryType uint32             `json:"delivery_type,omitempty"`
	Env          uint32             `json:"env"`
}

// CombinedPaymentSignData 合单支付的 signData 结构。
type CombinedPaymentSignData struct {
	Env               uint32                `json:"env"`
	CombinedOrderList []model.CombinedOrder `json:"combined_order_list"`
}

// CommonPaymentParams 返回给小程序 wx.requestCommonPayment 的参数集。
type CommonPaymentParams struct {
	SignData  string `json:"signData"`
	Mode      string `json:"mode"`
	PaySig    string `json:"paySig"`
	Signature string `json:"signature"`
}

// BuildCommonPaymentParams 生成小程序 wx.requestCommonPayment 所需参数。
// uri 必须与计算 paySig 的服务端 API 路径一致（例如 /retail/B2b/getorder）。
func (s *PaymentService) BuildCommonPaymentParams(ctx context.Context, uri string, signData any, appKey, sessionKey, mode string) (*CommonPaymentParams, error) {
	if appKey == "" {
		return nil, errors.New("appKey is required")
	}
	body, err := json.Marshal(signData)
	if err != nil {
		return nil, err
	}
	return &CommonPaymentParams{
		SignData:  string(body),
		Mode:      mode,
		PaySig:    signer.PaySig(uri, body, appKey),
		Signature: signer.UserSignature(body, sessionKey),
	}, nil
}
