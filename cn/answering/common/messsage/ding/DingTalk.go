package ding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
	"time"
)

const DING_TALK_API = "https://oapi.dingtalk.com/robot/send"

// 发送文字信息
func SendTextContent(secret string, accessToken string, content string) error {
	text := map[string]string{
		"content": content,
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	signStr := generateSign(secret, timestamp)

	req := &grequests.RequestOptions{
		Params: map[string]string{
			"access_token": accessToken,
			"timestamp":    timestamp,
			"sign":         signStr,
		},
		JSON: map[string]interface{}{
			"msgtype": "text",
			"text":    text,
		},
	}

	_, err := grequests.Post(DING_TALK_API, req)
	if err != nil {
		return err
	}
	return nil
}

// 生成DingTalk签名
func generateSign(secret string, timestamp string) string {
	signStr := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
