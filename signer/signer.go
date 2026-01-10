package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// PaySig 计算 pay_sig/paySig，算法为 HMAC-SHA256(appKey, uri+"&"+body)。
func PaySig(uri string, body []byte, appKey string) string {
	msg := uri + "&" + string(body)
	return hmacHex(appKey, msg)
}

// UserSignature 计算用户态签名，算法为 HMAC-SHA256(sessionKey, body)。
func UserSignature(body []byte, sessionKey string) string {
	return hmacHex(sessionKey, string(body))
}

func hmacHex(key, msg string) string {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
