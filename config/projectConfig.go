// @Create   : 2023/4/19 09:51
// @Author   : yaho
// @Remark   : 项目常量配置

package config

import (
	"time"
)

const (
	CacheDefaultExpiration = time.Minute * 5 // 缓存默认过期时间

	LogMaxSize      = 100              // 单个日志最大 单位MB
	LogFileMaxNum   = 60               // 日志文件最大数量
	LogFileSplitDay = 1                // 日志文件切割时间 单位天
	APITimeOut      = 10 * time.Second // 接口超时时间
)
