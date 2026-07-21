package model

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/service/users/token_cache_redis"
	"time"
)

// 创建 userFactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）

func CreateUserFactory(sqlType string) *UsersModel {
	return &UsersModel{GormM: GormM{DB: UseDbConn(sqlType)}}
}

func CreateUsersFactory(sqlType string) *UsersModels {
	return &UsersModels{GormM: GormM{DB: UseDbConn(sqlType)}}
}

type UsersModel struct {
	GormM
	Phone           string `gorm:"column:phone" json:"phone"`
	HeadImg         int    `gorm:"column:headImg" json:"headImg"`
	LastLoginIp     string `gorm:"column:last_login_ip" json:"last_login_ip"`
	LastLoginIpAddr string `gorm:"column:last_login_ipAddr" json:"last_login_ipAddr"`
	CreatedTime     string `gorm:"column:createdTime" json:"createdTime"`
	IsSuper         int    `gorm:"column:isSuper" json:"isSuper"`
}

type UsersModels struct {
	GormM
	Users []UsersModel `json:"Users"`
}

// 表名
func (u *UsersModel) TableName() string {
	return "user"
}

// 添加新管理用户
func (u *UsersModel) AddUser(phone, userIp, userAddr, createdTime string, isSuper int) bool {
	sql := "INSERT  INTO user(phone, headImg, last_login_ip, last_login_ipAddr, createdTime, isSuper) SELECT ?,?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM user WHERE  phone=?)"
	result := u.Exec(sql, phone, 1, userIp, userAddr, createdTime, isSuper, phone)
	if result.RowsAffected > 0 {
		return true
	} else {
		variable.ZapLog.Error("AddUser 存入用户信息出错：", zap.Error(result.Error))
		return false
	}
}

func (u *UsersModel) RemoveUser(phone string) bool {
	sql := "DELETE FROM user WHERE phone=?"
	result := u.Exec(sql, phone)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// 获取管理用户列表
func (us *UsersModels) GetUserList() *UsersModels {
	sql := "select phone, headImg, last_login_ip, last_login_ipAddr, createdTime, isSuper from user where isSuper=0"
	rows, err := us.Raw(sql).Rows() // 执行查询语句，返回查询结果的行数据集
	if err != nil {
		variable.ZapLog.Error("GetUserList 查询管理用户列表出错:", zap.Error(err))
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var user UsersModel
		err := rows.Scan(&user.Phone, &user.HeadImg, &user.LastLoginIp, &user.LastLoginIpAddr, &user.CreatedTime, &user.IsSuper)
		if err != nil {
			variable.ZapLog.Error("GetUserList 读取查询结果出错:", zap.Error(err))
			return nil
		}
		us.Users = append(us.Users, user)
	}
	return us
}

// 判断是否存在该用户
func (u *UsersModel) IsUser(phone string) bool {
	sql := "SELECT phone FROM user WHERE phone=?"
	u.Raw(sql, phone).First(u)
	if u.Phone != "" {
		return true
	} else {
		return false
	}
}

// 判断数据库中是否存在超级管理员
func (u *UsersModel) IsSuperUser() string {
	sql := "SELECT phone FROM user WHERE isSuper=1"
	u.Raw(sql).First(u)
	if u.Phone != "" {
		return u.Phone
	} else {
		return ""
	}
}

// 更新数据库中的超级管理员状态
func (u *UsersModel) UpSuperUser(phone string, isSuper int) {
	sql := "UPDATE user SET isSuper=? WHERE phone=?"
	res := u.Exec(sql, isSuper, phone)
	if res.Error != nil {
		variable.ZapLog.Error("UpSuperUser 数据库超级管理员状态更新出错:", zap.Error(res.Error))
	}
}

// 用户登录
func (u *UsersModel) Login(phone string) *UsersModel {
	sql := "select headImg, createdTime from user where phone=? limit 1"
	result := u.Raw(sql, phone).First(u)
	if result.Error == nil {
		return u
	} else {
		variable.ZapLog.Error("Login 根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}

// 记录用户登陆（login）生成的token，每次登陆记录一次token
func (u *UsersModel) OauthLoginToken(phone, token, created_at string, expiresAt int64, clientIp string) bool {
	sql := "INSERT  INTO userToken(phone,token,created_at,client_ip,expires_at) SELECT ?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM userToken WHERE  token=?)"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if u.Exec(sql, phone, token, created_at, clientIp, time.Unix(expiresAt, 0).Format(variable.DateFormat), token).Error == nil {
		// 异步缓存用户有效的token到redis
		if variable.ConfigYml.GetInt("Token.IsCacheToRedis") == 1 {
			go u.ValidTokenCacheToRedis(phone)
		}
		return true
	}
	return false
}

// 用户刷新token,条件检查: 相关token在过期的时间之内，就符合刷新条件
func (u *UsersModel) OauthRefreshConditionCheck(phone string, oldToken string) bool {
	// 首先判断旧token在本系统自带的数据库已经存在，才允许继续执行刷新逻辑
	var oldTokenIsExists int
	sql := "SELECT count(*)  as  counts FROM userToken  WHERE phone =? and token=? and NOW()<DATE_ADD(expires_at,INTERVAL  ? SECOND)"
	if u.Raw(sql, phone, oldToken, variable.ConfigYml.GetInt64("Token.JwtTokenRefreshAllowSec")).First(&oldTokenIsExists).Error == nil && oldTokenIsExists == 1 {
		return true
	}
	return false
}

// 用户刷新token
func (u *UsersModel) OauthRefreshToken(expiresAt int64, phone, oldToken, newToken, clientIp, ipAddr string) bool {
	sql := "UPDATE userToken SET token=?, expires_at=?, client_ip=?, updated_at=NOW() WHERE token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, oldToken).Error == nil {
		// 异步缓存用户有效的token到redis
		if variable.ConfigYml.GetInt("Token.IsCacheToRedis") == 1 {
			go u.ValidTokenCacheToRedis(phone)
		}
		go u.UpdateUserloginInfo(clientIp, ipAddr, phone)
		return true
	}
	return false
}

// 更新用户最近一次登录ip及所在地
func (u *UsersModel) UpdateUserloginInfo(last_login_ip, last_login_ipAddr, phone string) {
	sql := "UPDATE user SET last_login_ip=?, last_login_ipAddr=? WHERE phone=?"
	res := u.Exec(sql, last_login_ip, last_login_ipAddr, phone)
	if res.Error != nil {
		variable.ZapLog.Error("UpdateUserloginInfo 根据账号更新IP登录记录出错:", zap.Error(res.Error))
	}
}

// 当 user 删除数据，相关的token同步删除
func (u *UsersModel) OauthDestroyToken(phone string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	sql := "DELETE FROM  userToken WHERE  phone=?  "
	//判断>=0, 有些没有登录过的用户没有相关token，此语句执行影响行数为0，但是仍然是执行成功
	if u.Exec(sql, phone).Error == nil {
		return true
	}
	return false
}

// 判断用户token是否在数据库存在+状态OK
func (u *UsersModel) OauthCheckTokenIsOk(phone string, token string) bool {
	sql := "SELECT   token  FROM  `userToken`  WHERE   phone=?  AND  expires_at>NOW() ORDER  BY  expires_at  DESC , updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, phone, maxOnlineUsers).Rows()
	defer func() {
		//  凡是查询类记得释放记录集
		_ = rows.Close()
	}()
	if err == nil && rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					return true
				}
			}
		}
	}
	return false
}

// 删除用户以及关联的token记录
func (u *UsersModel) Destroy(phone string) bool {

	// 删除用户时，清除用户缓存在redis的全部token
	if variable.ConfigYml.GetInt("Token.IsCacheToRedis") == 1 {
		go u.DelTokenCacheFromRedis(phone)
	}
	if u.Delete(u, phone).Error == nil {
		if u.OauthDestroyToken(phone) {
			return true
		}
	}
	return false
}

// 后续两个函数专门处理用户 token 缓存到 redis 逻辑
func (u *UsersModel) ValidTokenCacheToRedis(phone string) {
	tokenCacheRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(phone)
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return
	}
	defer tokenCacheRedisFact.ReleaseRedisConn()

	sql := "SELECT token,expires_at FROM `userToken` WHERE phone=? AND  expires_at>NOW() ORDER BY expires_at DESC , updated_at DESC LIMIT ?"
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, phone, maxOnlineUsers).Rows()
	defer func() {
		//  凡是获取原生结果集的查询，记得释放记录集
		_ = rows.Close()
	}()

	var tempToken, expires string
	if err == nil && rows != nil {
		for i := 1; rows.Next(); i++ {
			err = rows.Scan(&tempToken, &expires)
			if err == nil {
				if ts, err := time.ParseInLocation(variable.DateFormat, expires, time.Local); err == nil {
					tokenCacheRedisFact.SetTokenCache(ts.Unix(), tempToken)
					// 因为每个用户的token是按照过期时间倒叙排列的，第一个是有效期最长的，将该用户的总键设置一个最大过期时间，到期则自动清理，避免不必要的数据残留
					if i == 1 {
						tokenCacheRedisFact.SetUserTokenExpire(ts.Unix())
					}
				} else {
					variable.ZapLog.Error("expires_at 转换位时间戳出错", zap.Error(err))
				}
			}
		}
	}
	// 缓存结束之后删除超过系统设置最大在线数量的token
	tokenCacheRedisFact.DelOverMaxOnlineCache()
}

// DelTokenCacheFromRedis 用户信息修改后，删除redis所有的token
func (u *UsersModel) DelTokenCacheFromRedis(phone string) {
	tokenCacheRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(phone)
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return
	}
	tokenCacheRedisFact.ClearUserToken()
	tokenCacheRedisFact.ReleaseRedisConn()
}
