package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/validator/common/upload_files"
	"goskeleton/app/http/validator/web/AI"
	"goskeleton/app/http/validator/web/BigData"
	"goskeleton/app/http/validator/web/users"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string
	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	key = consts.ValidatorPrefix + "PhoneCaptcha"
	containers.Set(key, users.Captcha{})
	key = consts.ValidatorPrefix + "UsersLogin"
	containers.Set(key, users.Login{})
	key = consts.ValidatorPrefix + "RefreshToken"
	containers.Set(key, users.RefreshToken{})

	// 超级管理员操作
	key = consts.ValidatorPrefix + "SuperAdminAddUser"
	containers.Set(key, users.AddUser{})
	key = consts.ValidatorPrefix + "GetUserList"
	containers.Set(key, users.GetUserList{})
	key = consts.ValidatorPrefix + "UpChatGPTKey"
	containers.Set(key, AI.UpChatGPTKey{})
	key = consts.ValidatorPrefix + "RemoveUser"
	containers.Set(key, users.RemoveUser{})
	key = consts.ValidatorPrefix + "UpDev"
	containers.Set(key, BigData.UpDev{})
	key = consts.ValidatorPrefix + "RemoveDev"
	containers.Set(key, BigData.RemoveDev{})

	// AI基本操作
	key = consts.ValidatorPrefix + "ProposalAI_AE"
	containers.Set(key, AI.AEProposalAI{})
	key = consts.ValidatorPrefix + "ProposalAI_AS"
	containers.Set(key, AI.ASProposalAI{})
	key = consts.ValidatorPrefix + "AgricultureExpertAI"
	containers.Set(key, AI.AgricultureExpertAI{})
	key = consts.ValidatorPrefix + "DetectionInfo"
	containers.Set(key, AI.DetectionInfo{})
	key = consts.ValidatorPrefix + "DetectionImage"
	containers.Set(key, AI.DetectionImage{})
	key = consts.ValidatorPrefix + "GetMonitorLogs"
	containers.Set(key, AI.GetMonitorLogs{})
	key = consts.ValidatorPrefix + "PestProposal"
	containers.Set(key, AI.PestProposal{})

	// 农业大数据基本操作
	key = consts.ValidatorPrefix + "GetAEdata"
	containers.Set(key, BigData.GetAEdata{})
	key = consts.ValidatorPrefix + "GetASdata"
	containers.Set(key, BigData.GetASdata{})
	key = consts.ValidatorPrefix + "GetDevInfo"
	containers.Set(key, BigData.GetDevInfo{})
	key = consts.ValidatorPrefix + "GetMemCPU"
	containers.Set(key, BigData.GetMemCPU{})

	// 文件上传
	key = consts.ValidatorPrefix + "UploadFiles"
	containers.Set(key, upload_files.UpFiles{})
	key = consts.ValidatorPrefix + "AiImageInfo"
	containers.Set(key, upload_files.AiImageInfo{})
}
