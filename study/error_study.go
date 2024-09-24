package study

import (
	"errors"
	"fmt"
)

func func1(n int) (int, error) {
	if n < 10 {
		return 0, errors.New("n 不能小于10")
	}
	return n, nil
}

func ErrorStudyMain() {
	a, e := func1(1)
	fmt.Println(a, e)
}
