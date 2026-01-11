package types

import "github.com/enjoy322/wechatpay-b2b/model"

// ProfitSharingRequest 请求分账参数。
type ProfitSharingRequest struct {
	Mchid          string                 `json:"mchid"`                     // 微信商户号
	TransactionID  string                 `json:"transaction_id"`            // 微信支付订单号
	OutOrderNo     string                 `json:"out_order_no"`              // 商户分账订单号
	UnfreezeUnpaid bool                   `json:"unfreeze_unpaid,omitempty"` // 是否解冻未分账金额
	Receivers      []model.ProfitReceiver `json:"receivers"`                 // 分账接收方列表
	Description    string                 `json:"description,omitempty"`     // 分账描述
}

// ProfitSharingResponse 请求分账返回参数。
type ProfitSharingResponse struct {
	OrderID       string `json:"order_id"`       // 分账订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账订单号
	TransactionID string `json:"transaction_id"` // 微信支付订单号
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

// QueryProfitSharingRequest 查询分账订单参数。
type QueryProfitSharingRequest struct {
	Mchid         string `json:"mchid"`                  // 微信商户号
	TransactionID string `json:"transaction_id"`         // 微信支付订单号
	OutOrderNo    string `json:"out_order_no,omitempty"` // 商户分账订单号
}

// QueryProfitSharingResponse 查询分账订单返回参数。
type QueryProfitSharingResponse struct {
	OrderID       string                 `json:"order_id"`       // 分账订单号
	OutOrderNo    string                 `json:"out_order_no"`   // 商户分账订单号
	TransactionID string                 `json:"transaction_id"` // 微信支付订单号
	Status        model.ProfitStatus     `json:"status"`         // 分账状态
	Receivers     []model.ProfitReceiver `json:"receivers"`      // 分账接收方列表
	Amount        int64                  `json:"amount"`         // 分账金额
	ErrCode       int                    `json:"errcode"`
	ErrMsg        string                 `json:"errmsg"`
}

// ProfitSharingFinishRequest 分账完结参数。
type ProfitSharingFinishRequest struct {
	Mchid         string `json:"mchid"`                 // 微信商户号
	TransactionID string `json:"transaction_id"`        // 微信支付订单号
	OutOrderNo    string `json:"out_order_no"`          // 商户分账订单号
	Description   string `json:"description,omitempty"` // 完结描述
}

// ProfitSharingFinishResponse 分账完结返回参数。
type ProfitSharingFinishResponse struct {
	OrderID    string `json:"order_id"`     // 分账订单号
	OutOrderNo string `json:"out_order_no"` // 商户分账订单号
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// ProfitSharingReturnRequest 分账回退参数。
type ProfitSharingReturnRequest struct {
	Mchid         string               `json:"mchid"`                 // 微信商户号
	TransactionID string               `json:"transaction_id"`        // 微信支付订单号
	OutOrderNo    string               `json:"out_order_no"`          // 商户分账订单号
	OutReturnNo   string               `json:"out_return_no"`         // 商户回退订单号
	ReturnAmount  int64                `json:"return_amount"`         // 回退金额
	Receiver      model.ProfitReceiver `json:"receiver"`              // 分账接收方
	Description   string               `json:"description,omitempty"` // 描述
}

// ProfitSharingReturnResponse 分账回退返回参数。
type ProfitSharingReturnResponse struct {
	ReturnID    string `json:"return_id"`     // 回退订单号
	OutReturnNo string `json:"out_return_no"` // 商户回退订单号
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// QueryProfitSharingReturnRequest 查询分账回退参数。
type QueryProfitSharingReturnRequest struct {
	Mchid       string `json:"mchid"`               // 微信商户号
	OutReturnNo string `json:"out_return_no"`       // 商户回退订单号
	ReturnID    string `json:"return_id,omitempty"` // 回退订单号
}

// QueryProfitSharingReturnResponse 查询分账回退返回参数。
type QueryProfitSharingReturnResponse struct {
	ReturnID    string                   `json:"return_id"`     // 回退订单号
	OutReturnNo string                   `json:"out_return_no"` // 商户回退订单号
	OrderID     string                   `json:"order_id"`      // 分账订单号
	Amount      int64                    `json:"amount"`        // 回退金额
	Status      model.ProfitReturnStatus `json:"status"`        // 回退状态
	ErrCode     int                      `json:"errcode"`
	ErrMsg      string                   `json:"errmsg"`
}

// AddProfitSharingAccountRequest 添加分账方请求参数。
type AddProfitSharingAccountRequest struct {
	Mchid                     string `json:"mchid"`                        // 微信商户号
	ProfitSharingRelationType string `json:"profit_sharing_relation_type"` // 分账接收方关系类型
	PayeeType                 string `json:"payee_type"`                   // 分账接收方类型
	PayeeID                   string `json:"payee_id"`                     // 分账接收方账号
	PayeeName                 string `json:"payee_name"`                   // 分账接收方名称
}

// AddProfitSharingAccountResponse 添加分账方返回参数。
type AddProfitSharingAccountResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DelProfitSharingAccountRequest 删除分账方请求参数。
type DelProfitSharingAccountRequest struct {
	Mchid     string `json:"mchid"`      // 微信商户号
	PayeeType string `json:"payee_type"` // 分账接收方类型
	PayeeID   string `json:"payee_id"`   // 分账接收方账号
}

// DelProfitSharingAccountResponse 删除分账方返回参数。
type DelProfitSharingAccountResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryProfitSharingAccountRequest 查询分账方请求参数。
type QueryProfitSharingAccountRequest struct {
	Mchid  string `json:"mchid"`  // 微信商户号
	Offset int    `json:"offset"` // 偏移量
	Limit  int    `json:"limit"`  // 分页大小，最大 100
}

// ProfitSharingAccount 分账接收方信息。
type ProfitSharingAccount struct {
	PayeeType                 string `json:"payee_type"`                   // 分账接收方类型
	PayeeID                   string `json:"payee_id"`                     // 分账接收方账号
	PayeeName                 string `json:"payee_name"`                   // 分账接收方名称
	ProfitSharingRelationType string `json:"profit_sharing_relation_type"` // 分账接收方关系类型
}

// QueryProfitSharingAccountResponse 查询分账方返回参数。
type QueryProfitSharingAccountResponse struct {
	AccountList []ProfitSharingAccount `json:"account_list"` // 分账接收方列表
	TotalCount  int                    `json:"total_count"`  // 总数量
	ErrCode     int                    `json:"errcode"`
	ErrMsg      string                 `json:"errmsg"`
}

// QueryProfitSharingRemainAmtRequest 查询分账剩余金额请求参数。
type QueryProfitSharingRemainAmtRequest struct {
	Mchid      string `json:"mchid"`        // 微信商户号
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
}

// QueryProfitSharingRemainAmtResponse 查询分账剩余金额返回参数。
type QueryProfitSharingRemainAmtResponse struct {
	RemainAmount int64  `json:"remain_amount"` // 剩余可分账金额
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}
