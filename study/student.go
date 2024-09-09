package study

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) Student {
	return Student{Name: name,
		Age: age}
}

func main() {
	stu := NewStudent("杨幂", 18)
	fmt.Println(stu)
}
