package study

// 当多个参数（连续）类型一致时，只需要在最后一个参数后标明类型
func add(x, y int) int {
	return x + y
}

func FuncStudyMain() int {
	return add(1, 2)
}
