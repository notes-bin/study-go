package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	//这里需要替换成自己的,我就用星号代替了
	AccessKeyId = "hello"
	secret      = "world"
)

var Aliyun map[string]interface{}

func PercentEncode(str string) string {
	//替换字符串
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func CreateSignature(secret, StringToSign string) string {
	//创建签名
	key := []byte(secret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(StringToSign))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s
}

func GetUtcTime() string {
	//获取utc时间
	timezero := time.FixedZone("GMT", 0)
	return time.Now().In(timezero).Format("2006-01-02T15:04:05Z")
}

func SendMail() string {
	//获取utc时间
	utc := GetUtcTime()
	//把请求参数一一添加起来
	data := map[string]string{
		"Action":           "SingleSendMail",
		"Format":           "JSON",
		"Version":          "2015-11-23",
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   utc,
		"SignatureVersion": "1.0",
		"AccessKeyId":      AccessKeyId,
		"Timestamp":        utc,
	}

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	params := make(url.Values)
	for _, v := range keys {
		params.Add(v, data[v])
	}

	fmt.Println(params.Encode())

	//参数encode编码,并替换参数中的特殊字符
	percent := PercentEncode(params.Encode())

	//获取待签名数据
	StringToSign := fmt.Sprintf("%s&%s&%s", "GET", url.QueryEscape("/"), url.QueryEscape(percent))

	//生成签名值
	Signature := CreateSignature(secret, StringToSign)
	fmt.Println(Signature)

	//把签名内容带到请求参数中
	params.Add("Signature", Signature)

	//发起get请求
	req := "https://dm.aliyuncs.com/?" + PercentEncode(params.Encode())
	return req
}

func main() {
	fmt.Println(SendMail())
}
