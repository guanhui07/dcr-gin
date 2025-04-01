package utils

import (
	"fmt"
	"reflect"
)

// ContainsString 判断obj是否在target中，target支持的类型array,slice,map
func ContainsString(obj any, target any) int {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return i
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return 1
		}
	}

	return -1
}

func ArrayTest() {
	testMap()

	testArray()
	testSlice()
}

func testArray() {
	a := 1
	b := [3]int{1, 2, 3}

	fmt.Println(ContainsString(a, b))

	c := "a"
	d := [4]string{"b", "c", "d", "a"}
	fmt.Println(ContainsString(c, d))

	e := 1.1
	f := [4]float64{1.2, 1.3, 1.1, 1.4}
	fmt.Println(ContainsString(e, f))

	g := 1
	h := [4]any{2, 4, 6, 1}
	fmt.Println(ContainsString(g, h))

	i := [4]int64{}
	fmt.Println(ContainsString(a, i))
}

func testSlice() {
	a := 1
	b := []int{1, 2, 3}

	fmt.Println(ContainsString(a, b))

	c := "a"
	d := []string{"b", "c", "d", "a"}
	fmt.Println(ContainsString(c, d))

	e := 1.1
	f := []float64{1.2, 1.3, 1.1, 1.4}
	fmt.Println(ContainsString(e, f))

	g := 1
	h := []any{2, 4, 6, 1}
	fmt.Println(ContainsString(g, h))

	var i []int64
	fmt.Println(ContainsString(a, i))
}

func testMap() {
	var a = map[int]string{1: "1", 2: "2"}
	fmt.Println(ContainsString(3, a))

	var b = map[string]int{"1": 1, "2": 2}
	fmt.Println(ContainsString("1", b))

	var c = map[string][]int{"1": {1, 2}, "2": {2, 3}}
	fmt.Println(ContainsString("6", c))
}
