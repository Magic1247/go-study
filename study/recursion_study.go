package study

func recall(num int) (result int) {
	if num == 1 {
		return 1
	}
	return num * recall(num-1)
}

func RecursionStudyMain(num int) int {
	return recall(num)
}
