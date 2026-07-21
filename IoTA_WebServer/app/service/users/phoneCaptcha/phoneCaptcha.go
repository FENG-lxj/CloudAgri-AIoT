package phoneCaptcha

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/redis_factory"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Create 生成验证码
func Create(phone string) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		if i == 0 {
			// 生成非零的随机数作为第一位
			code += strconv.Itoa(1 + rand.Intn(9))
		} else {
			code += strconv.Itoa(rand.Intn(10))
		}
	}
	if StorePhoneCode(phone, code) {
		if sendSMS(phone, code) {
			return code
		}
		return ""
	}
	return ""
}

// StorePhoneCode 存储手机验证码到Redis缓存，并设置过期时间为5分钟
func StorePhoneCode(phone string, code string) bool {
	// 获取当前时间戳
	CreationTime := time.Now().Unix()
	CodeAndTimestamp := code + "_" + strconv.FormatInt(CreationTime, 10)
	// 从连接池获取一个连接
	redisClient := redis_factory.GetOneRedisClient()
	// 将验证码和时间戳存入Redis缓存，设置过期时间为5分钟
	_, err := redisClient.String(redisClient.Execute("set", phone, CodeAndTimestamp, "EX", 300))
	if err != nil {
		variable.ZapLog.Error("StorePhoneCode 存储手机验证码失败", zap.Error(err))
		// 释放连接
		redisClient.ReleaseOneRedisClient()
		return false
	} else {
		variable.ZapLog.Info("StorePhoneCode 存储手机验证码成功", zap.String("phone", phone), zap.String("code_Timestamp", CodeAndTimestamp))
	}
	// 操作完毕记得释放连接
	redisClient.ReleaseOneRedisClient()
	return true
}

// GetPhoneCode 从Redis缓存中取出手机验证码并打印
func GetPhoneCode(phone string) (string, int64) {
	// 从连接池获取一个连接
	redisClient := redis_factory.GetOneRedisClient()
	// 从Redis缓存中取出手机验证码
	CodeAndTimestamp, err := redisClient.String(redisClient.Execute("get", phone))
	if err != nil {
		variable.ZapLog.Error("GetPhoneCode 获取手机验证码失败", zap.Error(err))
		// 释放连接
		redisClient.ReleaseOneRedisClient()
		return "", 0
	}
	variable.ZapLog.Info("GetPhoneCode 获取手机验证码成功", zap.String("phone", phone), zap.String("code_Timestamp", CodeAndTimestamp))
	// 拆分验证码和时间戳
	parts := strings.Split(CodeAndTimestamp, "_")
	code := parts[0]
	timestamp, _ := strconv.ParseInt(parts[1], 10, 64)
	// 记得释放连接
	redisClient.ReleaseOneRedisClient()
	return code, timestamp
}

func sendSMS(phone string, code string) bool {
	//url := "https://uni.apistd.com/?action=sms.message.send&accessKeyId=RkMVXTsVD84XvvvhLW1jQarq4YcmSGXXxwU1wTDg2RryfkrqR"
	//data := map[string]interface{}{
	//	"to":           phone,
	//	"signature":    "螭耀科创",
	//	"templateId":   "pub_verif_ttl3",
	//	"templateData": map[string]string{"code": code, "ttl": "5"},
	//}
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	variable.ZapLog.Error("编码 JSON 时出错：", zap.Error(err))
	//	return false
	//}
	//
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	//if err != nil {
	//	variable.ZapLog.Error("创建 HTTP 请求时出错：", zap.Error(err))
	//	return false
	//}
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	variable.ZapLog.Error("发送请求时出错：", zap.Error(err))
	//	return false
	//}
	//
	//defer resp.Body.Close()
	return true
}
