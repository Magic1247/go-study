package study

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

// 结构体作为参数
func printBook(book Book) {
	fmt.Println(book)
}

func StructStudyMain() {
	b1 := Book{
		"《Go语言从入门到拉垮》",
		"张三",
		"Go语言教程",
		123,
	}

	fmt.Println(b1)
	// 使用键值对赋值时，空余的键值会自动赋值为对应类型的零值
	fmt.Println(Book{title: "Python4.*", author: "吴川之父"})
	// 访问结构体成员
	fmt.Println(b1.title)

	printBook(b1)
	var structPtr1 *Book

	structPtr1 = &b1
	// 语法糖，结构体类型的指针，可以通过.直接访问成员，自动解引用，也可以显式操作
	fmt.Println(structPtr1.subject)
	fmt.Println((*structPtr1).subject)
}
