// @Create   : 2023/4/10 15:30
// @Author   : yaho
// @Remark   :

package convert

import (
	"encoding/json"

	"Goldfinger/errors"
)

func JsonConvert(from, to any) error {

	fromJson, err := json.Marshal(from)
	if err != nil {
		return errors.NewConvertError("struct -> struct时发生错误：from结构体序列化为json时：" + err.Error())
	}

	err = json.Unmarshal(fromJson, to)
	if err != nil {
		return errors.NewConvertError("struct -> struct时发生错误：from结构体的json解析到to结构体时：" + err.Error())
	}

	return nil
}
