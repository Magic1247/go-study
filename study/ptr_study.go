package study

import "fmt"

func f(pa *int) {
	*pa = 3
}

func PtrStudy() {
	var a int = 0
	// 传递的是指针
	f(&a)
	fmt.Println(a)
}
