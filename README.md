# wechatpay-b2b

微信支付 B2B（门店助手）Golang SDK。

## 项目定位

- 面向小程序 B2B 门店助手的“支付/退款/提现/合单支付”等能力，提供后端 SDK 封装（签名、请求结构体、调用入口等）。
- 小程序侧支付入口为 `wx.requestCommonPayment`，本 SDK 负责在服务端准备 `signData`、`paySig` 等参数并对接相关服务端接口。

## 官方文档

- 文档：<https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html#_3-3-%E6%94%AF%E4%BB%98%E4%B8%8E%E9%80%80%E6%AC%BE>
