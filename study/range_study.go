package study

import "fmt"

func RangeStudyMain() {
	// 使用range 遍历切片
	var s1 = []int{1, 2, 3, 4, 5}
	//可以使用_，忽略索引
	//for _, v := range s1 {
	for i, v := range s1 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// 遍历map
	var m1 = map[string]string{"name": "杨幂", "age": "18", "gender": "女"}

	for k, v := range m1 {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}
