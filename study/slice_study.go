package study

import (
	"fmt"
)

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func SliceStudyMain() {
	// 切片声明
	var s1 []int
	// append本质上是新建了一个更大的切片，然后将原切片的数据复制到新切片中，再追加元素
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(s1)
	// 使用make创建切片，容量为可选项
	var s2 = make([]string, 10)
	s2 = append(s2, "a", "b", "c", "d")
	fmt.Println(s2)

	fmt.Printf("切片s2的容量 %d \n", cap(s2))

	// 遍历切片
	//for i := 0; i < len(s2); i++ {
	//	fmt.Println(s2[i])
	//}

	// 切片初始化
	s3 := []int{1, 2, 3}
	// 追加元素，容量会自动扩容，扩容为原来的2倍
	s3 = append(s3, 31, 21)
	fmt.Println(cap(s3))

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(arr))
	// 数组切片 [start:end]
	s4 := arr[1:5]
	printSlice(s4)

	var s5 []int
	fmt.Printf("%T\n", s5)
	if s5 == nil {
		fmt.Println("切片为空")
	}
}
