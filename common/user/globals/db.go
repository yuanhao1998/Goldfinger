// @Create   : 2023/3/20 10:44
// @Author   : yaho
// @Remark   :

package globals

import (
	cache "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var CacheConn cache.UniversalClient
var SessionConn cache.UniversalClient
