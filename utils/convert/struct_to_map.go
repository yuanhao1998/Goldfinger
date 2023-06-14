// @Create   : 2023/4/27 11:04
// @Author   : yaho
// @Remark   :

package convert

import "reflect"

func StructToMapUseRef(from any) map[string]any {
	objValue := reflect.ValueOf(from)
	objType := objValue.Type()

	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}

	data := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		if fieldValue.CanInterface() {
			value := fieldValue.Interface()
			data[field.Name] = value
		}

	}

	return data
}
