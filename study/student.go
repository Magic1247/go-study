package study

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func NewStudent(name string, age int) Student {
	return Student{name: name,
		age: age}
}

func main() {
	stu := NewStudent("杨幂", 18)
	fmt.Println(stu)
}
