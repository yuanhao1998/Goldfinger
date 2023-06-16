// @Create   : 2023/3/20 10:46
// @Author   : yaho
// @Remark   :

package userConfig

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Goldfinger/common/user/globals"
)

func InitDB(conf *userGlobals.Conf) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.UserName, conf.Mysql.PWD, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DB)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	sqlConn, err := db.DB()
	if err != nil {
		panic("获取数据库sql连接失败：" + err.Error())
	}
	sqlConn.SetMaxIdleConns(conf.Mysql.MaxConn)
	sqlConn.SetMaxOpenConns(conf.Mysql.MaxOpen)

	return db
}

func InitCache(conf *userGlobals.Conf, db int) redis.UniversalClient {
	conn := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: conf.Redis.Addr, Password: conf.Redis.PWD, DB: db, MasterName: conf.Redis.MasterName,
		PoolSize: conf.Redis.PoolSize,
	})

	return conn
}
