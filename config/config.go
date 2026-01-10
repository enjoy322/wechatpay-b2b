package config

// Env 表示支付使用的运行环境。
type Env string

const (
	EnvProd    Env = "prod"
	EnvSandbox Env = "sandbox"
)

// Config 保存 SDK 共享配置。
type Config struct {
	AppID         string            // 小程序 appid。
	AppSecret     string            // 小程序 secret，用于获取 access_token。
	AppKey        map[string]string // 生产环境 appKey 映射（mchid -> appKey）。
	AppKeySandbox map[string]string // 沙箱环境 appKey 映射（mchid -> appKey）。
	BaseURL       string            // 默认 https://api.weixin.qq.com，允许自定义。
	Env           Env               // prod 或 sandbox。
}
