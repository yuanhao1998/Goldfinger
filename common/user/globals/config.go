// @Create   : 2023/3/17 22:30
// @Author   : yaho
// @Remark   :

package globals

var RunConf *Conf // 全局配置变量

// 项目配置
type rpcProject struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type apiProject struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// mysql配置
type mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"DB"`
	UserName string `json:"userName"`
	PWD      string `json:"PWD"`
	MaxConn  int    `json:"maxConn"` // 最大连接数量（保持持续连接数量）
	MaxOpen  int    `json:"maxOpen"` // 最大打开数量
}

// redis配置
type redis struct {
	Addr       []string `json:"Addr"` // 填写多个地址会采用集群模式
	PWD        string   `json:"PWD"`
	SessionDB  int      `json:"sessionDB"`  // session库
	CacheDB    int      `json:"cacheDB"`    // 业务缓存库
	MasterName string   `json:"masterName"` // 填写会采用主从模式
	PoolSize   int      `json:"poolSize"`
}

// 日志配置
type rpcLog struct {
	Level string `json:"level"`
	Path  string `json:"path"`
}
type apiLog struct {
	Level string `json:"level"`
	Path  string `json:"path"`
}

type Conf struct {
	RPCProject rpcProject `json:"rpcProject"`
	APIProject apiProject `json:"apiProject"`
	Mysql      mysql      `json:"mysql"`
	Redis      redis      `json:"redis"`
	RPCLog     rpcLog     `json:"rpcLog"`
	APILog     apiLog     `json:"apiLog"`
	SecretKey  string     `json:"secretKey"`
}
