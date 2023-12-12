// @Create   : 2023/3/20 10:46
// @Author   : yaho
// @Remark   :

package userConfig

import (
	"Goldfinger/common/user/rpc/src/model"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Goldfinger/common/user/globals"
)

func InitDB() *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userGlobals.RunConf.Mysql.UserName, userGlobals.RunConf.Mysql.PWD,
		userGlobals.RunConf.Mysql.Host, userGlobals.RunConf.Mysql.Port, userGlobals.RunConf.Mysql.DB)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	sqlConn, err := db.DB()
	if err != nil {
		panic("获取数据库sql连接失败：" + err.Error())
	}
	sqlConn.SetMaxIdleConns(userGlobals.RunConf.Mysql.MaxConn)
	sqlConn.SetMaxOpenConns(userGlobals.RunConf.Mysql.MaxOpen)

	model.CreateTable(db) // 自动建表

	return db
}

func InitCache() redis.UniversalClient {
	conn := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: userGlobals.RunConf.Redis.Addr, Password: userGlobals.RunConf.Redis.PWD,
		DB: userGlobals.RunConf.Redis.CacheDB, MasterName: userGlobals.RunConf.Redis.MasterName,
		PoolSize: userGlobals.RunConf.Redis.PoolSize,
	})

	return conn
}

func init() {
	userGlobals.DBConn = InitDB()       // 全局数据库连接
	userGlobals.CacheConn = InitCache() // 全局缓存连接
}
