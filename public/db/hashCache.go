// @Create   : 2023/4/17 17:54
// @Author   : yaho
// @Remark   :

package db

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"Goldfinger/public/model"
)

type HashCache[T model.BaseModel] struct {
	DbConn    *gorm.DB              // 数据库连接
	CacheConn redis.UniversalClient // 缓存连接
}

// CreateHash 创建数据，缓存使用hash存储
// ctx 上下文
// st 要创建的结构体
// cacheKey 缓存键
func (c *HashCache[T]) CreateHash(ctx context.Context, st *T, cacheKey string) (int64, error) {

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

	id := reflect.ValueOf(*st).FieldByName("Id").Int() // 获取新增数据库的ID
	c.CacheConn.HSet(ctx, cacheKey, dbJsonData)        // 写入cache

	tx.Commit()

	return id, nil
}

// UpdateHash 根据ID更新缓存和数据库的数据
// ctx 上下文
// st 从此结构体获取数据
// cacheKey 缓存键
// id 根据此id更新数据
func (c *HashCache[T]) UpdateHash(ctx context.Context, st *T, cacheKey string) error {

	tx := c.DbConn.Begin()

	res := tx.Save(&st)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	dbJsonData, err := json.Marshal(st) // json序列化，准备写入cache
	if err != nil {
		tx.Rollback()
		return err
	}
	id := reflect.ValueOf(*st).FieldByName("Id").Int() // 获取更新数据库的ID
	c.CacheConn.HSet(ctx, cacheKey, id, dbJsonData)    // 写入cache

	tx.Commit()

	return nil
}

// RetrieveHash 根据ID获取数据详情，优先从缓存获取
// ctx 上下文
// st 将序列化到此结构体
// cacheKey 缓存键
// id 根据此id查询数据
func (c *HashCache[T]) RetrieveHash(ctx context.Context, st *T, cacheKey string, id int64) error {

	cacheBytes, err := c.CacheConn.HGet(ctx, cacheKey, strconv.FormatInt(id, 10)).Bytes()
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
		c.CacheConn.HSet(ctx, cacheKey, id, "{}")
		return nil
	} else {

		// 如果此id已被删除，向缓存中写入一条空值
		isDel := reflect.ValueOf(*st).FieldByName("IsDel")
		if isDel.IsValid() || isDel.Bool() {
			c.CacheConn.HSet(ctx, cacheKey, id, "{}")
			return nil
		}

		dbJsonData, err := json.Marshal(st)
		if err != nil {
			return err
		}
		c.CacheConn.HSet(ctx, cacheKey, id, dbJsonData)
		return nil
	}
}

// DeleteHash 删除一条数据，包括缓存及数据库
// ctx 上下文
// cacheKey 缓存键
// id 根据此ID删除数据
func (c *HashCache[T]) DeleteHash(ctx context.Context, cacheKey string, id int64) error {

	tx := c.DbConn.Begin()

	res := tx.First(ctx, id).Update("is_del", 1) // 更新数据库
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	c.CacheConn.HDel(ctx, cacheKey, strconv.FormatInt(id, 10)) // 删除cache

	tx.Commit()

	return nil
}
