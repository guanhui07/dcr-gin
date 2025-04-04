package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// OnlyFields 只获取指定 key 的数据
func OnlyFields(fields map[string]any, keys ...string) map[string]any {
	var results = make(map[string]any)

	for _, key := range keys {
		results[key] = fields[key]
	}

	return results
}

func MakeKeysMap(keys ...string) map[string]any {
	var keysMap = map[string]any{}
	for _, key := range keys {
		keysMap[key] = 1
	}

	return keysMap
}

// ExceptFields 只获取指定 key 以外的数据
func ExceptFields(fields map[string]any, keys ...string) map[string]any {
	var (
		results = make(map[string]any)
		keysMap = MakeKeysMap(keys...)
	)

	for key, value := range fields {
		if _, exists := keysMap[key]; !exists {
			results[key] = value
		}
	}

	return results
}

// OnlyExistsFields 只获取指定 key ，不存在或者 nil 则忽略
func OnlyExistsFields(fields map[string]any, keys ...string) map[string]any {
	var results = make(map[string]any)

	for _, key := range keys {
		if value := fields[key]; value != nil {
			results[key] = value
		}
	}

	return results
}

// MergeFields 合并两个 map[string]any
func MergeFields(fields map[string]any, finalFields map[string]any) {
	for key, value := range finalFields {
		fields[key] = value
	}
}

// GetStringField 获取 Fields 中的字符串，会尝试转换类型
func GetStringField(fields map[string]any, key string, defaultValues ...string) string {
	if value, existsString := fields[key]; existsString {
		if str, isString := value.(string); isString {
			return str
		}
	}
	return StringOr(defaultValues...)
}

// GetSubField 获取下级 Fields ，如果没有的话，匹配同前缀的放到下级 Fields 中
func GetSubField(fields map[string]any, key string, defaultValues ...map[string]any) map[string]any {

	if subField, isField := fields[key].(map[string]any); isField {
		return subField
	}

	if len(defaultValues) > 0 {
		return defaultValues[0]
	}

	subField := make(map[string]any)
	prefix := key + "."

	for fieldKey, fieldValue := range fields {
		if strings.HasPrefix(fieldKey, prefix) {
			subField[strings.ReplaceAll(fieldKey, prefix, "")] = fieldValue
		}
	}

	if len(subField) > 0 {
		fields[key] = subField
	}

	return subField
}

// GetInt64Field 获取 Fields 中的 int64，会尝试转换类型
func GetInt64Field(fields map[string]any, key string, defaultValues ...int64) int64 {
	var defaultValue int64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int64); isInt {
			return intValue
		}
		return ToInt64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetIntField 获取 Fields 中的 int，会尝试转换类型
func GetIntField(fields map[string]any, key string, defaultValues ...int) int {
	var defaultValue = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int); isInt {
			return intValue
		}
		return int(ToInt64(value, int64(defaultValue)))
	} else {
		return defaultValue
	}
}

// GetFloatField 获取 Fields 中的 float32，会尝试转换类型
func GetFloatField(fields map[string]any, key string, defaultValues ...float32) float32 {
	var defaultValue float32 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float32); isInt {
			return intValue
		}
		return ToFloat(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetFloat64Field 获取 Fields 中的 float64，会尝试转换类型
func GetFloat64Field(fields map[string]any, key string, defaultValues ...float64) float64 {
	var defaultValue float64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float64); isInt {
			return intValue
		}
		return ToFloat64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetBoolField 获取 Fields 中的 bool，会尝试转换类型
func GetBoolField(fields map[string]any, key string, defaultValues ...bool) bool {
	var defaultValue = false
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if fieldValue, existsValue := fields[key]; existsValue {
		return ToBool(fieldValue, defaultValue)
	}
	return defaultValue
}

// ToFields 尝试把一个变量转换成 Fields 类型
func ToFields(anyValue any) (map[string]any, error) {
	fields := map[string]any{}
	switch paramValue := anyValue.(type) {
	case map[string]any:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[any]any:
		for key, value := range paramValue {
			fields[fmt.Sprintf("%v", key)] = value
		}
	case map[string]int:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int8:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int16:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint8:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint16:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]float64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]float32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]string:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]bool:
		for key, value := range paramValue {
			fields[key] = value
		}
	default:
		paramType := reflect.ValueOf(anyValue)
		switch paramType.Kind() {
		case reflect.Ptr: // 结构体指针
			if paramType.Elem().Kind() == reflect.Struct {
				EachStructField(paramType.Elem(), paramType.Elem().Interface(), func(field reflect.StructField, value reflect.Value) {
					if field.IsExported() {
						fields[SnakeString(field.Name)] = value.Interface()
					} else {
						fields[SnakeString(field.Name)] = nil
					}
				})
			}
		case reflect.Struct: // 结构体
			EachStructField(paramType, anyValue, func(field reflect.StructField, value reflect.Value) {
				if field.IsExported() {
					fields[SnakeString(field.Name)] = value.Interface()
				} else {
					fields[SnakeString(field.Name)] = nil
				}
			})
		case reflect.Map: // 自定义的 map
			for _, key := range paramType.MapKeys() {
				name := key.String()
				fields[name] = paramType.MapIndex(key).Interface()
			}
		default:
			return nil, errors.New("不支持转 map[string]any 的类型： " + paramType.String())
		}
	}
	return fields, nil
}
