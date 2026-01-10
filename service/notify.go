package service

import (
	"encoding/xml"
)

// PaymentNotify 表示支付/退款的通知结构体，使用微信推送的 XML 载荷。
type PaymentNotify struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	AppID        string `xml:"appid"`
	Mchid        string `xml:"mchid"`
	OutTradeNo   string `xml:"out_trade_no,omitempty"`
	OrderID      string `xml:"order_id,omitempty"`
	OutRefundNo  string `xml:"out_refund_no,omitempty"`
	RefundID     string `xml:"refund_id,omitempty"`
}

func ParseNotify(body []byte) (*PaymentNotify, error) {
	var n PaymentNotify
	if err := xml.Unmarshal(body, &n); err != nil {
		return nil, err
	}
	return &n, nil
}
