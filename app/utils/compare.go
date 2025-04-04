package utils

import (
	"fmt"
	"reflect"
)

type CompareHandler func(comparable any, arg any) bool

func Compare(comparable any, operator string, arg any) bool {
	switch operator {
	case "=", "eq":
		return IsEqual(comparable, arg)
	case ">=", "gte":
		return IsGte(comparable, arg)
	case ">", "gt":
		return IsGt(comparable, arg)
	case "<", "lt":
		return IsLt(comparable, arg)
	case "<=", "lte":
		return IsLte(comparable, arg)
	case "in":
		return IsIn(comparable, arg)
	case "not in":
		return IsNotIn(comparable, arg)
	}
	return false
}

// IsEqual 等于 =
func IsEqual(comparable any, arg any) bool {
	comparableType := reflect.TypeOf(comparable)
	argType := reflect.TypeOf(arg)

	if comparableType.Comparable() && argType.Comparable() && comparable == arg {
		return true
	}

	switch comparableType.Kind() {
	case reflect.Bool:
		return comparable.(bool) == ToBool(arg, !comparable.(bool)) // 若不能转换成bool，则默认不相等
	case reflect.String:
		return comparable.(string) == ToString(arg, fmt.Sprintf("%v", arg))
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Float64, reflect.Float32:
		return ToFloat64(comparable, 0) == ToFloat64(arg, 0)
	// 判断结构体是否相等
	case reflect.Struct:
		// 不是同一种结构体
		if !IsSameStruct(comparableType, arg) {
			return false
		}

		argValue := reflect.ValueOf(arg)
		isSame := true
		EachStructField(reflect.ValueOf(comparable), comparable, func(field reflect.StructField, value reflect.Value) {
			if !field.IsExported() {
				isSame = false
			}

			if isSame && !IsEqual(value.Interface(), argValue.FieldByName(field.Name).Interface()) {
				isSame = false
			}
		})
		return isSame
	case reflect.Array, reflect.Slice:
		comparableValue := reflect.ValueOf(comparable)
		argValue := reflect.ValueOf(arg)
		if comparableValue.Len() != argValue.Len() {
			return false
		}

		isSame := true
		EachSlice(comparableValue, func(index int, value reflect.Value) {
			if isSame && !IsEqual(value.Interface(), argValue.Index(index).Interface()) {
				isSame = false
			}
		})
		return isSame
	}
	return false
}

// IsIn 存在 in
func IsIn(comparable any, arg any) (result bool) {
	argValue := reflect.ValueOf(arg)
	if !IsArray(argValue) {
		return false
	}

	EachSlice(argValue, func(index int, value reflect.Value) {
		if !result && IsEqual(comparable, value.Interface()) {
			result = true
		}
	})
	return
}

// IsNotIn 不存在 not in
func IsNotIn(comparable any, arg any) (result bool) {
	argValue := reflect.ValueOf(arg)
	if !IsArray(argValue) {
		return false
	}
	EachSlice(argValue, func(index int, value reflect.Value) {
		if !result && IsEqual(comparable, value.Interface()) {
			result = true
		}
	})
	return !result
}

// IsLt 小于 <
func IsLt(comparable any, arg any) bool {
	return ToFloat64(comparable, 0) < ToFloat64(arg, 0)
}

// IsLte 小于等于 <=
func IsLte(comparable any, arg any) bool {
	return ToFloat64(comparable, 0) <= ToFloat64(arg, 0)
}

// IsGt 大于 >
func IsGt(comparable any, arg any) bool {
	return ToFloat64(comparable, 0) > ToFloat64(arg, 0)
}

// IsGte 大于等于
func IsGte(comparable any, arg any) bool {
	return ToFloat64(comparable, 0) >= ToFloat64(arg, 0)
}

// IsArray 是否是数组或者 slice
func IsArray(comparable any) bool {
	switch value := comparable.(type) {
	case reflect.Type:
		return value.Kind() == reflect.Slice || value.Kind() == reflect.Array
	case reflect.Value:
		return value.Kind() == reflect.Slice || value.Kind() == reflect.Array
	}
	comparableType := reflect.TypeOf(comparable)
	return comparableType.Kind() == reflect.Slice || comparableType.Kind() == reflect.Array
}

func IsNotInT[T string | bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](target T, list []T) bool {
	for _, v := range list {
		if target == v {
			return false
		}
	}
	return true
}

func IsInT[T string | bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](target T, list []T) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}
