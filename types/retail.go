package types

// BatchCreateRetailRequest 预录入门店信息请求参数（Body 部分）。
type BatchCreateRetailRequest struct {
	RetailInfoList []RetailInfo `json:"retail_info_list"`
}

// RetailInfo 门店信息。
type RetailInfo struct {
	MobilePhone        string   `json:"mobile_phone"`
	RetailName         string   `json:"retail_name"`
	RetailType         string   `json:"retail_type,omitempty"`
	SubRetailType      string   `json:"sub_retail_type,omitempty"`
	AddressProvince    string   `json:"address_province"`
	AddressCity        string   `json:"address_city"`
	AddressRegion      string   `json:"address_region"`
	AddressStreet      string   `json:"address_street"`
	RegistrationNumber string   `json:"registration_number,omitempty"`
	BizName            string   `json:"biz_name,omitempty"`
	CorporationName    string   `json:"corporation_name,omitempty"`
	Latitude           float64  `json:"latitude,omitempty"`
	Longitude          float64  `json:"longitude,omitempty"`
	BusinessType       []string `json:"business_type,omitempty"`
	OtherBusinessType  string   `json:"other_business_type,omitempty"`
}

// BatchCreateRetailResponse 预录入门店信息返回参数。
type BatchCreateRetailResponse struct {
	ErrCode           int                           `json:"errcode"`
	ErrMsg            string                        `json:"errmsg"`
	NumSuccess        int                           `json:"num_success,omitempty"`
	NumFailure        int                           `json:"num_failure,omitempty"`
	FailureRecordList []BatchCreateRetailFailRecord `json:"failure_record_list,omitempty"`
}

// BatchCreateRetailFailRecord 单条导入失败记录。
type BatchCreateRetailFailRecord struct {
	MobilePhone        string `json:"mobile_phone"`
	RegistrationNumber string `json:"registration_number,omitempty"`
	FailureCode        int    `json:"failure_code"`
}

// GetRetailInfoRequest 查询门店信息请求参数。
type GetRetailInfoRequest struct {
	OpenID      string `json:"openid,omitempty"`       // 门店 openid
	MobilePhone string `json:"mobile_phone,omitempty"` // 门店手机号
}

// GetRetailInfoResponse 查询门店信息返回参数。
type GetRetailInfoResponse struct {
	OpenID             string   `json:"openid"`                        // 门店 openid
	MobilePhone        string   `json:"mobile_phone"`                  // 门店手机号
	RetailName         string   `json:"retail_name"`                   // 门店名称
	RetailType         string   `json:"retail_type,omitempty"`         // 门店类型
	SubRetailType      string   `json:"sub_retail_type,omitempty"`     // 子门店类型
	AddressProvince    string   `json:"address_province"`              // 省份
	AddressCity        string   `json:"address_city"`                  // 城市
	AddressRegion      string   `json:"address_region"`                // 区县
	AddressStreet      string   `json:"address_street"`                // 街道
	RegistrationNumber string   `json:"registration_number,omitempty"` // 注册号
	BizName            string   `json:"biz_name,omitempty"`            // 商户名称
	CorporationName    string   `json:"corporation_name,omitempty"`    // 法人姓名
	Latitude           float64  `json:"latitude,omitempty"`            // 纬度
	Longitude          float64  `json:"longitude,omitempty"`           // 经度
	BusinessType       []string `json:"business_type,omitempty"`       // 经营类型
	OtherBusinessType  string   `json:"other_business_type,omitempty"` // 其他经营类型
	ErrCode            int      `json:"errcode"`
	ErrMsg             string   `json:"errmsg"`
}

// GetRetailOpenIDListRequest 全量授权门店查询请求参数。
type GetRetailOpenIDListRequest struct {
	Limit       int    `json:"limit"`                  // 分页大小，最大 100
	PageContext string `json:"page_context,omitempty"` // 分页上下文，首次查询不传
}

// GetRetailOpenIDListResponse 全量授权门店查询返回参数。
type GetRetailOpenIDListResponse struct {
	OpenIDList  []string `json:"openid_list"`            // openid 列表
	PageContext string   `json:"page_context,omitempty"` // 分页上下文，用于下次查询
	HasMore     bool     `json:"has_more"`               // 是否还有更多数据
	ErrCode     int      `json:"errcode"`
	ErrMsg      string   `json:"errmsg"`
}
