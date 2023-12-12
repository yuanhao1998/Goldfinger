// @Create   : 2023/3/23 15:07
// @Author   : yaho
// @Remark   :

package userConfig

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	userGlobals "Goldfinger/common/user/globals"
	"Goldfinger/config"
	"Goldfinger/globals"
)

func init() {
	path, err := filepath.Abs("./common/user/config/")
	if err != nil {
		panic("读取配置文件绝对路径失败：" + err.Error())
	}

	v, err := config.InitConf(path)
	println(v.GetInt("Mysql.maxOpen"))
	if err != nil {
		panic("读取配置文件失败：" + err.Error())
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		globals.Logger.Info("监听到配置文件修改，准备重载配置:" + in.Name)
		// 重载配置
		if err := v.Unmarshal(&userGlobals.RunConf); err != nil {
			globals.Logger.Error("重载配置出错：" + err.Error())
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&userGlobals.RunConf); err != nil {
		panic(err)
	}

}
