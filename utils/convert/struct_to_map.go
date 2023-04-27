// @Create   : 2023/4/27 11:04
// @Author   : yaho
// @Remark   :

package convert

import "reflect"

func StructToMapUseRef(from any) map[string]any {
	m := make(map[string]interface{})
	v := reflect.ValueOf(from)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		m[field.Name] = value
	}

	return m
}
