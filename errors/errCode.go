// @Create   : 2023/4/18 10:04
// @Author   : yaho
// @Remark   :

package errors

const (
	DB          = iota // 数据库错误
	Params             // 参数错误
	TimeOut            // 超时错误
	NameRepeat         // 名称重复
	DataConvert        // 数据结构转换错误
	PWDDecode          // 密码解码错误
	PWDEncode          // 密码加密错误
)
