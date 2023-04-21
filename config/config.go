// @Create   : 2023/3/17 17:45
// @Author   : yaho
// @Remark   :

package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConf(path string) (*viper.Viper, error) {
	err := os.Setenv("GO_ENV", "dev")
	if err != nil {
		return nil, err
	}

	// 初始化viper
	v := viper.New()
	v.SetConfigName(os.Getenv("GO_ENV"))
	v.SetConfigType("json")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, err

}
