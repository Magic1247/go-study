package study

import "fmt"

// 当多个参数（连续）类型一致时，只需要在最后一个参数后标明类型
func add(x, y int) (result int) {
	// 定制了返回值后就已经声明了同名变量，复制后return可以不用写返回值
	result = x + y
	return
}

// 声明不定项参数
func mo(data1 int, data2 ...int) {
	fmt.Println(data1, data2)
}

// 闭包函数

func mo2() func() {
	return func() {
		fmt.Println("闭包函数调用")
	}
}

func FuncStudyMain() int {
	arr := []int{1, 1, 2}
	mo(1, arr...)
	mo(1, 2, 3)

	(func() {
		fmt.Println("自执行函数")
	})()
	mo2()()
	return add(1, 2)

}
