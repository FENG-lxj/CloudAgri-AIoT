package curd

import (
	"context"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"goskeleton/app/utils/redis_factory"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func CreateAICurdFactory() *AICurd {
	return &AICurd{model.CreateAIFactory("")}
}

type AICurd struct {
	AIModel *model.AIModel
}

//// GetAIProposalAE 获取AI实时生态数据农业生产建议  通义千问版
//func (a *AICurd) GetAIProposalAE(AEID int32) string {
//	AEDatas := model.CreateInfluxDBAEFactory().GetAEdata(AEID)
//	cropper, err := model.CreateAIFactory("").QueryCropsAE(AEID)
//	if err != nil {
//		log.Println("GetAIProposalAE 查询农作物种类出错: ", err)
//		return ""
//	}
//
//	userMessage := "【当前农田中种植农作物为" + cropper + "】\n" +
//		"当前种植环境数据如下：\n" +
//		"实时光照强度为" + Float32ToString(AEDatas.LI) + "lx，" +
//		"PH值为" + Float32ToString(AEDatas.PH) + "，土壤温度为" + Float32ToString(AEDatas.ST) + "℃，" +
//		"土壤湿度为" + Float32ToString(AEDatas.SM) + "%，土壤氮元素含量为" + Float32ToString(AEDatas.N) + "mg/kg，" +
//		"土壤磷元素含量为" + Float32ToString(AEDatas.P) + "mg/kg，土壤钾元素含量为" + Float32ToString(AEDatas.K) + "mg/kg，" +
//		"空气温度为" + Float32ToString(AEDatas.AT) + "℃，空气湿度为" + Float32ToString(AEDatas.AH) + "%，空气质量为" + Float32ToString(AEDatas.AQ) + "，" +
//		"二氧化碳浓度为" + Float32ToString(AEDatas.CO2) + "ppm，甲醇浓度为" + Float32ToString(AEDatas.CH4O) + "ppm\n" +
//		"【你是智能农业管理系统的一部分，你的任务是使用最专业的知识，根据当前的实时农业生态环境数据对已种植的农作物提出管理建议，使用系统化的格式进行分析，使用简单的形式说明结果，不要使用Markdown格式，不要过度分条，不要在回答中提到数据是我提供的，在你的眼中是从系统中获取到的实时采集的数据】"
//
//	client := openai.NewClient(
//		option.WithAPIKey("your-api-key-here"), // 替换为你的阿里云百炼 API Key
//		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
//	)
//	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
//		Messages: []openai.ChatCompletionMessageParamUnion{
//			openai.UserMessage(userMessage),
//		},
//		Model: "qwen-max",
//	})
//	if err != nil {
//		variable.ZapLog.Error("GetAIProposalAE 向通义千问发起请求失败：", zap.Error(err))
//		return ""
//	}
//
//	return chatCompletion.Choices[0].Message.Content
//}

// qwen-plus
func (a *AICurd) GetAIProposalAE(AEID int32) string {
	AEDatas := model.CreateInfluxDBAEFactory().GetAEdata(AEID)
	cropper, err := model.CreateAIFactory("").QueryCropsAE(AEID)
	if err != nil {
		log.Println("GetAIProposalAE 查询农作物种类出错: ", err)
		return ""
	}

	userMessage := "【当前农田中种植农作物为" + cropper + "】\n" +
		"当前种植环境数据如下：\n" +
		"实时光照强度为" + Float32ToString(AEDatas.LI) + "lx，" +
		"PH值为" + Float32ToString(AEDatas.PH) + "，土壤温度为" + Float32ToString(AEDatas.ST) + "℃，" +
		"土壤湿度为" + Float32ToString(AEDatas.SM) + "%，土壤氮元素含量为" + Float32ToString(AEDatas.N) + "mg/kg，" +
		"土壤磷元素含量为" + Float32ToString(AEDatas.P) + "mg/kg，土壤钾元素含量为" + Float32ToString(AEDatas.K) + "mg/kg，" +
		"空气温度为" + Float32ToString(AEDatas.AT) + "℃，空气湿度为" + Float32ToString(AEDatas.AH) + "%，空气质量为" + Float32ToString(AEDatas.AQ) + "，" +
		"二氧化碳浓度为" + Float32ToString(AEDatas.CO2) + "ppm，甲醇浓度为" + Float32ToString(AEDatas.CH4O) + "ppm\n" +
		"【你是智能农业管理系统的一部分，你的任务是使用最专业的知识，根据当前的实时农业生态环境数据对已种植的农作物提出管理建议，使用系统化的格式进行分析，使用简单的形式说明结果，不要使用Markdown格式，" +
		"不要过度分条，不要在回答中提到数据是我提供的，在你的眼中是从系统中获取到的实时采集的数据，不要出现**，##,--这些符号，每段内容开头空两格，不要出现空行现象】"

	client := openai.NewClient(
		option.WithAPIKey("your-api-key-here"), // 替换为你的阿里云百炼 API Key
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		},
		Model: "qwen-plus", // ✅ 仅修改这里即可
	})
	if err != nil {
		variable.ZapLog.Error("GetAIProposalAE 向通义千问发起请求失败：", zap.Error(err))
		return ""
	}

	return chatCompletion.Choices[0].Message.Content
}

// GetAIProposalAE 获取AI实时生态数据农业生产建议  kimi版
//func (a *AICurd) GetAIProposalAE(AEID int32) string {
//	AEDatas := model.CreateInfluxDBAEFactory().GetAEdata(AEID)
//	cropper, err := model.CreateAIFactory("").QueryCropsAE(AEID)
//	if err != nil {
//		log.Println("GetAIProposalAE 查询农作物种类出错: ", err)
//		return ""
//	}
//
//	// 定义目标URL和刷新令牌
//	apiUrl := variable.ConfigYml.GetString("AI.apiUrl")
//	//refreshToken := variable.ConfigYml.GetString("AI.kimi_Token")
//	refreshToken := GetChatGPTAPIKey()
//
//	data := map[string]interface{}{
//		"model": "kimi",
//		"messages": []map[string]string{
//			{
//				"role": "user",
//				"content": "当前农田中种植农作物为" + cropper + "，当前实时光照强度为" + Float32ToString(AEDatas.LI) + "lx，" +
//					"PH值为" + Float32ToString(AEDatas.PH) + "，土壤温度为" + Float32ToString(AEDatas.ST) + "℃，" +
//					"土壤湿度为" + Float32ToString(AEDatas.SM) + "%，土壤氮元素含量为" + Float32ToString(AEDatas.N) + "mg/kg，" +
//					"土壤磷元素含量为" + Float32ToString(AEDatas.P) + "mg/kg，土壤钾元素含量为" + Float32ToString(AEDatas.K) + "mg/kg，" +
//					"空气温度为" + Float32ToString(AEDatas.AT) + "℃，空气湿度为" + Float32ToString(AEDatas.AH) + "%，空气质量为" + Float32ToString(AEDatas.AQ) + "，" +
//					"二氧化碳浓度为" + Float32ToString(AEDatas.CO2) + "ppm，甲醇浓度为" + Float32ToString(AEDatas.CH4O) + "ppm，请根据当前的实时农业生态环境数据对已种植的农作物提出管理建议。不要说是我提供的数据，说根据当前采集到的数据分析。所有文字使用文本格式。",
//			},
//		},
//		"use_search": true,
//		"stream":     false,
//	}
//	// 准备要发送的JSON数据
//	payload, err := json.Marshal(data)
//	if err != nil {
//		log.Println("GetAIProposalAE map转JSON出错: ", err)
//		return ""
//	}
//	// 创建HTTP POST请求
//	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payload))
//	if err != nil {
//		log.Println("GetAIProposalAE 创建HTTP请求出错: ", err)
//		return ""
//	}
//	// 设置请求头部信息
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+refreshToken)
//	// 发起HTTP请求
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("GetAIProposalAE 发出HTTP请求出错: ", err)
//		return ""
//	}
//	defer resp.Body.Close()
//	// 读取响应内容
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("GetAIProposalAE 读取响应正文出错: ", err)
//		return ""
//	}
//	// 定义结构体来解析JSON响应
//	var response struct {
//		Choices []struct {
//			Message struct {
//				Content string `json:"content"` // 定义content字段
//			} `json:"message"` // 定义message字段
//		} `json:"choices"` // 定义choices字段
//	}
//	// 解析JSON响应到response结构体
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		log.Println("GetAIProposalAE 解码JSON响应出错: ", err)
//		return ""
//	}
//	// 检查是否有choices字段
//	if len(response.Choices) > 0 {
//		// 提取content字段的值
//		result := response.Choices[0].Message.Content
//		result = strings.Replace(result, "\n\n", "\n", -1)
//		result = strings.Replace(result, "**", "", -1)
//
//		return result
//	} else {
//		log.Println("GetAIProposalAE 在响应中找不到任何选择: ", err)
//		return ""
//	}
//}

// GetAIProposalAE 获取AI实时生态数据农业生产建议   ChatGPT版
//func (a *AICurd) GetAIProposalAE(AEID int32) string {
//	AEDatas := model.CreateInfluxDBAEFactory().GetAEdata(AEID)
//	cropper, err := model.CreateAIFactory("").QueryCropsAE(AEID)
//	if err != nil {
//		log.Println("GetAIProposalAE 查询农作物种类出错: ", err)
//		return ""
//	}
//
//	//ChatGPT_TOKEN := variable.ConfigYml.GetString("AI.ChatGPT_TOKEN")
//	ChatGPT_TOKEN := GetChatGPTAPIKey()
//	config := gogpt.DefaultConfig(ChatGPT_TOKEN)
//	proxyUrl, err := url.Parse(variable.ConfigYml.GetString("AI.ProxyUrl"))
//	if err != nil {
//		log.Println("GetAIProposalAE 代理服务器出错: ", err)
//		return ""
//	}
//	transport := &http.Transport{
//		Proxy: http.ProxyURL(proxyUrl),
//	}
//	config.HTTPClient = &http.Client{
//		Transport: transport,
//	}
//
//	c := gogpt.NewClientWithConfig(config)
//
//	resp, err := c.CreateChatCompletion(
//		context.Background(),
//		gogpt.ChatCompletionRequest{
//			Model: gogpt.GPT3Dot5Turbo,
//			Messages: []gogpt.ChatCompletionMessage{
//				{
//					Role: gogpt.ChatMessageRoleUser,
//					Content: "当前农田中种植农作物为" + cropper + "，当前实时光照强度为" + Float32ToString(AEDatas.LI) + "lx，" +
//						"PH值为" + Float32ToString(AEDatas.PH) + "，土壤温度为" + Float32ToString(AEDatas.ST) + "℃，" +
//						"土壤湿度为" + Float32ToString(AEDatas.SM) + "%，土壤氮元素含量为" + Float32ToString(AEDatas.N) + "mg/kg，" +
//						"土壤磷元素含量为" + Float32ToString(AEDatas.P) + "mg/kg，土壤钾元素含量为" + Float32ToString(AEDatas.K) + "mg/kg，" +
//						"空气温度为" + Float32ToString(AEDatas.AT) + "℃，空气湿度为" + Float32ToString(AEDatas.AH) + "%，空气质量为" + Float32ToString(AEDatas.AQ) + "，" +
//						"二氧化碳浓度为" + Float32ToString(AEDatas.CO2) + "%，甲醇浓度为" + Float32ToString(AEDatas.CH4O) + "%，请根据当前的实时农业生态环境数据对已种植的农作物提出管理建议。",
//				},
//			},
//		},
//	)
//	if err != nil {
//		log.Printf("GetAIProposalAE ChatGPT调用出错: %v\n", err)
//		return ""
//	}
//
//	result := resp.Choices[0].Message.Content
//	result = strings.Replace(result, "\n\n", "\n", -1)
//	//log.Println(result)
//
//	return result
//}

// GetAIProposalAS 获取AI实时生态数据农业生产建议   通义千问版
func (a *AICurd) GetAIProposalAS(ASID int32) string {
	ASDatas := model.CreateInfluxDBASFactory().GetASdata(ASID)
	cropper, err := model.CreateAIFactory("").QueryCropsAS(ASID)
	if err != nil {
		log.Println("GetAIProposalAS 查询农作物种类出错: ", err)
		return ""
	}

	userMessage := "【当前粮食仓库中储存的粮食为" + cropper + "】\n" +
		"当前粮仓实时环境数据如下：\n" +
		"实时环境温度为" + Float32ToString(ASDatas.T) + "℃，" +
		"环境湿度为" + Float32ToString(ASDatas.H) + "%，二氧化碳浓度为" + Float32ToString(ASDatas.CO2) + "ppm，" +
		"氧气浓度为" + Float32ToString(ASDatas.O2) + "%\n" +
		"【你是智能农业管理系统的一部分，你的任务是使用最专业的知识，根据当前的实时粮仓内部环境数据提出管理建议，使用系统化的格式进行分析，" +
		"使用简单的形式说明结果，不要使用Markdown格式，不要过度分条，不要在回答中提到数据是我提供的，" +
		"在你的眼中是从系统中获取到的实时采集的数据，不要出现**，##,--这些符号，每段内容开头空两格】"

	client := openai.NewClient(
		option.WithAPIKey("your-api-key-here"), // 替换为你的阿里云百炼 API Key
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		},
		Model: "qwen-plus", // ✅ 仅修改这里即可
	})
	if err != nil {
		variable.ZapLog.Error("GetAIProposalAS 向通义千问发起请求失败：", zap.Error(err))
		return ""
	}

	msg := strings.ReplaceAll(chatCompletion.Choices[0].Message.Content, "#", "")
	msg = strings.ReplaceAll(msg, "**", "")

	return msg
}

// 获取AI实时生态数据农业生产建议   kimi版
//func (a *AICurd) GetAIProposalAS(ASID int32) string {
//	ASDatas := model.CreateInfluxDBASFactory().GetASdata(ASID)
//	cropper, err := model.CreateAIFactory("").QueryCropsAS(ASID)
//	if err != nil {
//		log.Println("GetAIProposalAS 查询农作物种类出错: ", err)
//		return ""
//	}
//
//	// 定义目标URL和刷新令牌
//	apiUrl := variable.ConfigYml.GetString("AI.apiUrl")
//	//refreshToken := variable.ConfigYml.GetString("AI.kimi_Token")
//	refreshToken := GetChatGPTAPIKey()
//
//	data := map[string]interface{}{
//		"model": "kimi",
//		"messages": []map[string]string{
//			{
//				"role": "user",
//				"content": "当前粮食仓库中储存的粮食为" + cropper + "，当前实时环境温度为" + Float32ToString(ASDatas.T) + "℃，" +
//					"环境湿度为" + Float32ToString(ASDatas.H) + "%，二氧化碳浓度为" + Float32ToString(ASDatas.CO2) + "ppm，" +
//					"氧气浓度为" + Float32ToString(ASDatas.O2) + "%，请根据当前的实时农作物存储环境数据提出管理建议。不要说是我提供的数据，说根据当前采集到的数据分析。所有文字使用文本格式。",
//			},
//		},
//		"use_search": true,
//		"stream":     false,
//	}
//	// 准备要发送的JSON数据
//	payload, err := json.Marshal(data)
//	if err != nil {
//		log.Println("GetAIProposalAS map转JSON出错: ", err)
//		return ""
//	}
//	// 创建HTTP POST请求
//	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payload))
//	if err != nil {
//		log.Println("GetAIProposalAS 创建HTTP请求出错: ", err)
//		return ""
//	}
//	// 设置请求头部信息
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+refreshToken)
//	// 发起HTTP请求
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("GetAIProposalAS 发出HTTP请求出错: ", err)
//		return ""
//	}
//	defer resp.Body.Close()
//	// 读取响应内容
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("GetAIProposalAS 读取响应正文出错: ", err)
//		return ""
//	}
//	// 定义结构体来解析JSON响应
//	var response struct {
//		Choices []struct {
//			Message struct {
//				Content string `json:"content"` // 定义content字段
//			} `json:"message"` // 定义message字段
//		} `json:"choices"` // 定义choices字段
//	}
//	// 解析JSON响应到response结构体
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		log.Println("GetAIProposalAS 解码JSON响应出错: ", err)
//		return ""
//	}
//	// 检查是否有choices字段
//	if len(response.Choices) > 0 {
//		// 提取content字段的值
//		result := response.Choices[0].Message.Content
//		result = strings.Replace(result, "\n\n", "\n", -1)
//		result = strings.Replace(result, "**", "", -1)
//		return result
//	} else {
//		log.Println("GetAIProposalAS 在响应中找不到任何选择: ", err)
//		return ""
//	}
//}

// 获取AI实时农作物存储环境数据相关管理建议   ChatGPT版
//func (a *AICurd) GetAIProposalAS(ASID int32) string {
//	ASDatas := model.CreateInfluxDBASFactory().GetASdata(ASID)
//	cropper, err := model.CreateAIFactory("").QueryCropsAS(ASID)
//	if err != nil {
//		log.Println("GetAIProposalAS 查询农作物种类出错: ", err)
//		return ""
//	}
//
//	//ChatGPT_TOKEN := variable.ConfigYml.GetString("AI.ChatGPT_TOKEN")
//	ChatGPT_TOKEN := GetChatGPTAPIKey()
//	config := gogpt.DefaultConfig(ChatGPT_TOKEN)
//	proxyUrl, err := url.Parse(variable.ConfigYml.GetString("AI.ProxyUrl"))
//	if err != nil {
//		log.Println("GetAIProposalAS 代理服务器出错: ", err)
//		return ""
//	}
//	transport := &http.Transport{
//		Proxy: http.ProxyURL(proxyUrl),
//	}
//	config.HTTPClient = &http.Client{
//		Transport: transport,
//	}
//
//	c := gogpt.NewClientWithConfig(config)
//
//	resp, err := c.CreateChatCompletion(
//		context.Background(),
//		gogpt.ChatCompletionRequest{
//			Model: gogpt.GPT3Dot5Turbo,
//			Messages: []gogpt.ChatCompletionMessage{
//				{
//					Role: gogpt.ChatMessageRoleUser,
//					Content: "当前粮食仓库中储存的粮食为" + cropper + "，当前实时环境温度为" + Float32ToString(ASDatas.T) + "℃，" +
//						"环境湿度为" + Float32ToString(ASDatas.H) + "%，二氧化碳浓度为" + Float32ToString(ASDatas.CO2) + "ppm，" +
//						"氧气浓度为" + Float32ToString(ASDatas.O2) + "%，请根据当前的实时农作物存储环境数据提出管理建议。不要说是我提供的数据，说根据当前采集到的数据分析。",
//				},
//			},
//		},
//	)
//	if err != nil {
//		log.Printf("GetAIProposalAS ChatGPT调用出错: %v\n", err)
//		return ""
//	}
//
//	result := resp.Choices[0].Message.Content
//	result = strings.Replace(result, "\n\n", "\n", -1)
//
//	return result
//}

func Float32ToString(mun float32) string {
	return strconv.FormatFloat(float64(mun), 'f', -1, 32)
}

// AgricultureExpertAI 向AI农业专家提问   通义千问版
func (a *AICurd) AgricultureExpertAI(Issue string) string {
	userMessage := "【你是一个AI农业专家】\n" +
		"【你将使用最专业的农业知识回答用户的问题，切记只能回答农业相关的问题，其他的问题一概不允许回答，不要出现**，##,--这些符号，每段内容开头空两格】\n" +
		"【注意：所有问题的回答不要使用Markdown格式输出】" +
		"用户对你说：" + Issue

	client := openai.NewClient(
		option.WithAPIKey("your-api-key-here"), // 替换为你的阿里云百炼 API Key
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		},
		Model: "qwen-plus", // ✅ 仅修改这里即可
	})
	if err != nil {
		variable.ZapLog.Error("AgricultureExpertAI 向通义千问发起请求失败：", zap.Error(err))
		return ""
	}

	return chatCompletion.Choices[0].Message.Content
}

// 向AI农业专家提问   kimi版
//func (a *AICurd) AgricultureExpertAI(Issue string) string {
//	// 定义目标URL和刷新令牌
//	apiUrl := variable.ConfigYml.GetString("AI.apiUrl")
//	//refreshToken := variable.ConfigYml.GetString("AI.kimi_Token")
//	refreshToken := GetChatGPTAPIKey()
//	log.Println(apiUrl)
//	log.Println(refreshToken)
//
//	data := map[string]interface{}{
//		"model": "kimi",
//		"messages": []map[string]string{
//			{
//				"role":    "user",
//				"content": "接下来你将作为一位农业专家回答我的农业相关的问题，只回答与农业相关的问题，不帮助做任何其他事情：" + Issue,
//			},
//		},
//		"use_search": true,
//		"stream":     false,
//	}
//
//	// 准备要发送的JSON数据
//	payload, err := json.Marshal(data)
//	if err != nil {
//		log.Println("AgricultureExpertAI map转JSON出错: ", err)
//		return ""
//	}
//
//	// 创建HTTP POST请求
//	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payload))
//	if err != nil {
//		log.Println("AgricultureExpertAI 创建HTTP请求出错: ", err)
//		return ""
//	}
//
//	// 设置请求头部信息
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+refreshToken)
//
//	// 发起HTTP请求
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("AgricultureExpertAI 发出HTTP请求出错: ", err)
//		return ""
//	}
//	defer resp.Body.Close()
//
//	// 读取响应内容
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("AgricultureExpertAI 读取响应正文出错: ", err)
//		return ""
//	}
//	log.Println(string(body))
//
//	// 定义结构体来解析JSON响应
//	var response struct {
//		Choices []struct {
//			Message struct {
//				Content string `json:"content"` // 定义content字段
//			} `json:"message"` // 定义message字段
//		} `json:"choices"` // 定义choices字段
//	}
//
//	// 解析JSON响应到response结构体
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		log.Println("AgricultureExpertAI 解码JSON响应出错: ", err)
//		return ""
//	}
//	log.Println(response)
//
//	// 检查是否有choices字段
//	if len(response.Choices) > 0 {
//		// 提取content字段的值
//		result := response.Choices[0].Message.Content
//		result = strings.Replace(result, "\n\n", "\n", -1)
//		result = strings.Replace(result, "**", "", -1)
//		return result
//	} else {
//		log.Println("AgricultureExpertAI 在响应中找不到任何选择: ", err)
//		return ""
//	}
//}

// 向AI农业专家提问   ChatGPT版
//func (a *AICurd) AgricultureExpertAI(Issue string) string {
//	//ChatGPT_TOKEN := variable.ConfigYml.GetString("AI.ChatGPT_TOKEN")
//	ChatGPT_TOKEN := GetChatGPTAPIKey()
//	config := gogpt.DefaultConfig(ChatGPT_TOKEN)
//	proxyUrl, err := url.Parse(variable.ConfigYml.GetString("AI.ProxyUrl"))
//	if err != nil {
//		log.Println("AgricultureExpertAI 代理服务器出错: ", err)
//		return ""
//	}
//	transport := &http.Transport{
//		Proxy: http.ProxyURL(proxyUrl),
//	}
//	config.HTTPClient = &http.Client{
//		Transport: transport,
//	}
//
//	c := gogpt.NewClientWithConfig(config)
//
//	resp, err := c.CreateChatCompletion(
//		context.Background(),
//		gogpt.ChatCompletionRequest{
//			Model: gogpt.GPT3Dot5Turbo,
//			Messages: []gogpt.ChatCompletionMessage{
//				{
//					Role:    gogpt.ChatMessageRoleUser,
//					Content: "接下来你将作为一位农业专家回答我的农业相关的问题，如果是与农业不相干的问题就只回答这一句“抱歉，我是一位农业专家，我只能回答与农业相关的问题。如果你有农业方面的问题，我会很乐意帮助你。”，只回答与农业相关的问题，不帮助做任何事情：" + Issue,
//				},
//			},
//		},
//	)
//	if err != nil {
//		// 从错误消息中提取状态码
//		statusCode := extractStatusCode(err.Error())
//		if statusCode == 429 {
//			log.Println("ChatGPT key不可用: ", err)
//			return "429"
//		}
//		log.Printf("ChatGPT请求错误错误状态码: %d\n", statusCode)
//		log.Println("AgricultureExpertAI ChatGPT调用出错: ", err)
//		return ""
//	}
//
//	result := resp.Choices[0].Message.Content
//	result = strings.Replace(result, "\n\n", "\n", -1)
//	//log.Println(result)
//
//	return result
//}

// 从错误消息中提取状态码的函数
func extractStatusCode(errorMessage string) int {
	re := regexp.MustCompile(`status code:\s*(\d+)`)
	match := re.FindStringSubmatch(errorMessage)
	if len(match) > 1 {
		statusCodeStr := match[1]
		statusCode, err := strconv.Atoi(statusCodeStr)
		if err != nil {
			log.Println("extractStatusCode 无法解析错误状态码：", err)
			return 0
		}
		return statusCode
	}
	return 0
}

func (a *AICurd) UpChatGPTKey(APIKEY string) bool {
	return model.CreateAPIKEYFactory("").UpChatGPTKey(APIKEY)
}

func GetChatGPTAPIKey() string {
	// 从连接池获取一个连接
	redisClient := redis_factory.GetOneRedisClient()
	// 检查APIKEY是否存在
	if value, err := redisClient.String(redisClient.Execute("GET", "APIKEY")); err == nil {
		if value != "" {
			// 释放连接
			redisClient.ReleaseOneRedisClient()
			return value
		}
		redisClient.ReleaseOneRedisClient()
		return ""
	} else {
		variable.ZapLog.Error("GetChatGPTAPIKey 判断APIKEY在redis中是否缓存出错：", zap.Error(err))
		APIKEY := model.CreateAPIKEYFactory("").GetChatGPTAPIKey()
		if APIKEY != "" {
			// 将APIKEY存入Redis缓存
			_, err := redisClient.String(redisClient.Execute("set", "APIKEY", APIKEY))
			if err != nil {
				variable.ZapLog.Error("GetChatGPTAPIKey redis存储APIKEY出错：", zap.Error(err))
				redisClient.ReleaseOneRedisClient()
				return ""
			}
			redisClient.ReleaseOneRedisClient()
			return APIKEY
		}
		redisClient.ReleaseOneRedisClient()
		return ""
	}
}
