package curd

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"time"
)

func CreateUserCurdFactory() *UsersCurd {
	return &UsersCurd{model.CreateUserFactory("")}
}

func CreateUsersCurdFactory() *UsersCurds {
	return &UsersCurds{model.CreateUsersFactory("")}
}

type UsersCurd struct {
	userModel *model.UsersModel
}

type UsersCurds struct {
	UsersModels *model.UsersModels
}

// 判断用户是否存在
func (u *UsersCurd) IsUser(phone string) bool {
	return u.userModel.IsUser(phone)
}

// 登录
func (u *UsersCurd) Login(phone string) *model.UsersModel {
	return u.userModel.Login(phone)
}

// 添加管理用户
func (u *UsersCurd) AddUser(phone, userIP, userAddr, datetime string, isSuper int) bool {
	return u.userModel.AddUser(phone, userIP, userAddr, datetime, isSuper)
}

// 删除管理用户
func (u *UsersCurd) RemoveUser(phone string) bool {
	return u.userModel.RemoveUser(phone)
}

// 获取管理用户列表
func (us *UsersCurds) GetUserList() *model.UsersModels {
	return us.UsersModels.GetUserList()
}

// 超级管理员初始化
func (u *UsersCurd) SuperUserInit() {
	SuperPhone := variable.ConfigYml.GetString("User.Superuser")
	// 判断超级管理是否已初始化
	sqlPhone := u.userModel.IsSuperUser()
	if sqlPhone != "" {
		if SuperPhone != sqlPhone {
			if model.CreateUserFactory("").IsUser(SuperPhone) {
				u.userModel.UpSuperUser(SuperPhone, 1)
				u.userModel.UpSuperUser(sqlPhone, 0)
			} else {
				u.userModel.AddUser(SuperPhone, "127.0.0.1", "", time.Now().Format(variable.DateFormat), 1)
				u.userModel.UpSuperUser(sqlPhone, 0)
			}
		}
	} else {
		u.userModel.AddUser(SuperPhone, "127.0.0.1", "", time.Now().Format(variable.DateFormat), 1)
	}
}
