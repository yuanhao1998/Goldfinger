// @Create   : 2023/4/24 11:36
// @Author   : yaho
// @Remark   :

package db

import "reflect"

// 如果data是指针、获取指针指向的值，否则直接返回data
func findRealData(data any) reflect.Value {
	refValue := reflect.ValueOf(data)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	return refValue
}
