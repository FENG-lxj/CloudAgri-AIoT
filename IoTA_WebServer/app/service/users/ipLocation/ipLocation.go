package ipLocation

import (
	"encoding/json"
	"go.uber.org/zap"
	"golang.org/x/net/html/charset"
	"goskeleton/app/global/variable"
	"io"
	"net/http"
	"strings"
)

type IPInfo struct {
	IP          string `json:"ip"`
	Pro         string `json:"pro"`
	ProCode     string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

// GetIPLocation 解析IP属地信息
func GetIPLocation(ip string) string {
	url := "https://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"
	resp, err := http.Get(url)
	if err != nil {
		variable.ZapLog.Error("GetIPLocation http.Get(url) 获取IP属地失败", zap.Error(err))
		return ""
	}
	defer resp.Body.Close()

	// 转码为UTF-8
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		variable.ZapLog.Error("GetIPLocation charset.NewReader 获取IP属地失败", zap.Error(err))
		return ""
	}

	out, err := io.ReadAll(reader)
	if err != nil {
		variable.ZapLog.Error("GetIPLocation io.ReadAll(reader) 获取IP属地失败", zap.Error(err))
		return ""
	}

	outStr := string(out)
	// 移除可能的前缀
	outStr = strings.TrimPrefix(outStr, "if(window.IPCallBack) {IPCallBack(")
	// 移除可能的后缀
	outStr = strings.TrimSuffix(outStr, ")}")

	var result IPInfo
	if err := json.Unmarshal([]byte(outStr), &result); err != nil {
		variable.ZapLog.Error("GetIPLocation json.Unmarshal 获取IP属地失败", zap.Error(err))
		return ""
	}
	return result.Pro + result.City
}
