// @Create   : 2023/4/19 09:32
// @Author   : yaho
// @Remark   :

package db

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"

	"Goldfinger/config"
	"Goldfinger/globals"
	"Goldfinger/public/model"
)

type StringCache[T model.BaseModel] struct {
	DbConn    *gorm.DB              // 数据库连接
	CacheConn redis.UniversalClient // 缓存连接
}

// CreateStringWithExp 创建数据，缓存使用string存储
// ctx 上下文
// st 要创建的结构体
// cacheKey 缓存键
// expiration 过期时间
func (c *StringCache[T]) CreateStringWithExp(ctx context.Context, st *T, cacheKey string, expiration time.Duration) (int64, error) {

	tx := c.DbConn.Begin()

	res := tx.Create(&st) // 写入数据库
	if res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}

	dbJsonData, err := json.Marshal(st) // json序列化，准备写入cache
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id := reflect.ValueOf(*st).FieldByName("Id")
	if id.IsValid() {

		// 写入cache
		redisKey := fmt.Sprintf("%s#%d", cacheKey, id.Int())
		if err := c.CacheConn.Set(ctx, redisKey, dbJsonData, expiration).Err(); err != nil {
			globals.Logger.Error("CreateString方法中，数据写入数据库成功，但写入缓存时候出现错误，错误信息：%s，写入key：%s，写入数据：%s", err.Error(), redisKey, dbJsonData)
		}
	}

	tx.Commit()

	return id.Int(), nil
}

// CreateString 创建数据，缓存使用string存储，使用默认过期时间
// ctx 上下文
// st 要创建的结构体
// cacheKey 缓存键
func (c *StringCache[T]) CreateString(ctx context.Context, st *T, cacheKey string) (int64, error) {
	return c.CreateStringWithExp(ctx, st, cacheKey, config.CacheDefaultExpiration)
}

// UpdateStringWithExp 根据ID更新缓存和数据库的数据
// ctx 上下文
// st 从此结构体获取数据
// cacheKey 缓存键
// id 根据此id更新数据
func (c *StringCache[T]) UpdateStringWithExp(ctx context.Context, st *T, cacheKey string, expiration time.Duration) (int64, error) {

	tx := c.DbConn.Begin()

	res := tx.Save(&st)
	if res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}

	dbJsonData, err := json.Marshal(st) // json序列化，准备写入cache
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id := reflect.ValueOf(*st).FieldByName("Id")

	if id.IsValid() {
		redisKey := fmt.Sprintf("%s#%d", cacheKey, id.Int())
		c.CacheConn.Set(ctx, redisKey, dbJsonData, expiration) // 写入cache
		tx.Commit()

		return id.Int(), nil
	}

	tx.Commit()

	return 0, nil
}

// UpdateString 更新数据，缓存使用string存储，使用默认过期时间
// ctx 上下文
// st 要创建的结构体
// cacheKey 缓存键
func (c *StringCache[T]) UpdateString(ctx context.Context, st *T, cacheKey string) (int64, error) {
	return c.UpdateStringWithExp(ctx, st, cacheKey, config.CacheDefaultExpiration)
}

// RetrieveStringWithExp 根据ID获取数据详情，优先从缓存获取，不存在的数据会向缓存写入一条空值，可以自定义过期此条空值的时间
// ctx 上下文
// st 将序列化到此结构体
// cacheKey 缓存键
// id 根据此id查询数据
// expiration 当数据不存在时写入缓存的空值的过期时间
func (c *StringCache[T]) RetrieveStringWithExp(ctx context.Context, st *T, cacheKey string, id int64, exp time.Duration) error {
	redisKey := fmt.Sprintf("%s#%d", cacheKey, id)
	cacheBytes, err := c.CacheConn.Get(ctx, redisKey).Bytes()
	if err != nil {
		goto db
	}
	if len(cacheBytes) != 0 {
		if err := json.Unmarshal(cacheBytes, &st); err != nil {
			goto db
		}
	}
	return nil

db:
	if c.DbConn.First(&st, id).RowsAffected == 0 { // 如果不存在此ID
		c.CacheConn.Set(ctx, redisKey, "{}", exp)
		return nil
	} else {

		// 如果此id已被删除，向缓存中写入一条空值
		isDel := reflect.ValueOf(*st).FieldByName("IsDel")
		if isDel.IsValid() || isDel.Bool() {
			c.CacheConn.Set(ctx, redisKey, "{}", exp)
			return nil
		}

		dbJsonData, err := json.Marshal(st)
		if err != nil {
			return err
		}
		c.CacheConn.Set(ctx, redisKey, dbJsonData, exp)
		return nil
	}
}

// RetrieveString 根据ID获取数据详情，优先从缓存获取，不存在的数据会向缓存写入一条空值
// ctx 上下文
// st 将序列化到此结构体
// cacheKey 缓存键
// id 根据此id查询数据
func (c *StringCache[T]) RetrieveString(ctx context.Context, st *T, cacheKey string, id int64) error {
	return c.RetrieveStringWithExp(ctx, st, cacheKey, id, config.CacheDefaultExpiration)
}

// DeleteString 删除一条数据，包括缓存及数据库
// ctx 上下文
// cacheKey 缓存键
// id 根据此ID删除数据
func (c *StringCache[T]) DeleteString(ctx context.Context, cacheKey string, id int64) (int64, error) {
	redisKey := fmt.Sprintf("%s#%d", cacheKey, id)

	tx := c.DbConn.Begin()

	res := tx.First(&T{}, id).Update("is_del", 1) // 更新数据库
	if res.Error != nil {
		tx.Rollback()
		return id, res.Error
	}

	c.CacheConn.Del(ctx, redisKey) // 删除cache

	tx.Commit()

	return id, nil
}
