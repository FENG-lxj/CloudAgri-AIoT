package MySQL

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var sql *gorm.DB

func Init() {
	//连接到mysql数据库
	dsn := viper.GetString("MySQL.User") + ":" + viper.GetString("MySQL.Passwd") + "@tcp(" + viper.GetString("MySQL.Host") + ":" + viper.GetString("MySQL.Port") + ")/" + viper.GetString("MySQL.DataBase") + "?charset=" + viper.GetString("MySQL.Charset") + "&parseTime=True&loc=Local"
	var err error
	sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("连接到mysql出错:")
		log.Fatal(err)
	}
}

func GetSQL() *gorm.DB {
	return sql
}
