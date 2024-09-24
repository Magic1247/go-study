package study

import "fmt"

func pointerTest() {
	var a = 123
	fmt.Println(a)
	// * 代表为指针
	var b *int
	// & 代表取变量的内存地址
	b = &a
	// * 代表取指针指向的变量的值
	*b = 333
	fmt.Println(a, b)

	fmt.Println(a == *b)
	fmt.Println(&a == b)

	var str1 = "佟丽娅"
	p1 := &str1
	*p1 = "古力娜扎"
	fmt.Println(str1)
}

func pointerFunc(p1 *string) {
	*p1 = "迪丽热巴"
}

func PointerStudyMain() {
	var str = "杨幂"
	pointerFunc(&str)
	fmt.Println(str)
	pointerTest()
}
