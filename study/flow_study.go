package study

import "fmt"

func FlowStudy() {
	num1 := 199
	num1++
	fmt.Println(num1)
	if num1 > 200 {
		fmt.Println("大于两百")
	} else if num1 == 200 {
		fmt.Println("等于两百")
	} else {
		fmt.Println("内部异常")
	}

	switch num1 {
	case 200:

		fmt.Println("等于200")
		// 满足后继续往下执行
		fallthrough

	case 201:

		fmt.Println("等于201")
	case 202:
		fmt.Println("等于202")

	default:

		fmt.Println("都不成立")

	}

	a := 0
	for {
		a++
		fmt.Println(a)
		if a > 10 {
			break
		}
	}

	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	a = 0
	for a < 10 {
		a++
		if a == 5 {
			continue
		}
		fmt.Println(a)
	}
A:

	for a := 0; a < 10; a++ {
		if a == 5 {

			break A
			goto B

		}
	}

B:
	fmt.Println("跳转到B代码块")

}
