// @Create   : 2023/3/23 10:45
// @Author   : yaho
// @Remark   : 错误定义

package errors

type BaseError struct {
	Msg string
}

func (e BaseError) GenerateErrorMap(errType string, errCode int) map[string]any {
	err := map[string]any{
		"ErrType": errType,
		"ErrCode": errCode,
		"ErrMsg":  e.Msg,
	}
	return err
}

func (e BaseError) Error() string {
	return e.Msg
}

// NameRepeatError 名称重复错误
type NameRepeatError struct {
	BaseError
}

func (e NameRepeatError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("名称重复错误", NameRepeat)
}

func NewNameRepeatError(msg string) *NameRepeatError {
	return &NameRepeatError{BaseError{Msg: msg}}
}

// DBError 数据库错误
type DBError struct {
	BaseError
}

func (e DBError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("数据库错误", DB)
}

func NewDBError(msg string) *DBError {
	return &DBError{BaseError{Msg: msg}}
}

// CacheError 缓存错误
type CacheError struct {
	BaseError
}

func (e CacheError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("缓存错误", Cache)
}

func NewCacheError(msg string) *CacheError {
	return &CacheError{BaseError{Msg: msg}}
}

// ParamsError 参数错误
type ParamsError struct {
	BaseError
}

func (e ParamsError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("参数错误", Params)
}

func NewParamsError(msg string) *ParamsError {
	return &ParamsError{BaseError{Msg: msg}}
}

// TimeOutError 超时错误
type TimeOutError struct {
	BaseError
}

func (e TimeOutError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("超时错误", TimeOut)
}

func NewTimeOutError(msg string) *TimeOutError {
	return &TimeOutError{BaseError{Msg: msg}}
}

// DataConvertError 数据结构转换错误
type DataConvertError struct {
	BaseError
}

func (e DataConvertError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("数据结构转换错误", DataConvert)
}

func NewConvertError(msg string) *DataConvertError {
	return &DataConvertError{BaseError{Msg: msg}}
}

// PWDDecodeError 密码解码错误
type PWDDecodeError struct {
	BaseError
}

func (e PWDDecodeError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("密码解码错误", PWDDecode)
}

func NewPWDDecodeError(msg string) *PWDDecodeError {
	return &PWDDecodeError{BaseError{Msg: msg}}
}

// PWDEncodeError 密码加密错误
type PWDEncodeError struct {
	BaseError
}

func (e PWDEncodeError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("密码加密错误", PWDEncode)
}

func NewPWDEncodeError(msg string) *PWDEncodeError {
	return &PWDEncodeError{BaseError{Msg: msg}}
}

// CaptchaError 验证码错误
type CaptchaError struct {
	BaseError
}

func (e CaptchaError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("验证码错误", Captcha)
}

func NewCaptchaError(msg string) *CaptchaError {
	return &CaptchaError{BaseError{Msg: msg}}
}

// LoginError 登陆错误
type LoginError struct {
	BaseError
}

func (e LoginError) ErrorMap() map[string]any {
	return e.GenerateErrorMap("登陆错误", Login)
}

func NewLoginError(msg string) *LoginError {
	return &LoginError{BaseError{Msg: msg}}
}
