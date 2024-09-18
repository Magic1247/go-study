package study

import "fmt"

func MapStudy() {
	var m map[string]int
	m = map[string]int{}

	m1 := map[string]int{"a": 1, "b": 2}

	// 使用空接口接收任何类型的参数
	m2 := map[int]interface{}{}
	m2[1] = "测试"
	m2[2] = 111

	delete(m1, "a")

	fmt.Println(m, m1, m2)

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
