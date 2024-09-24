package study

import "fmt"

type Person interface {
	say()
}

type Star struct {
	name string
	age  int
}

func (s Star) say() {
	fmt.Printf("我是208，叫做%s，%d岁\n", s.name, s.age)
}

type Singer struct {
	name string
	song string
}

func (s Singer) say() {
	fmt.Printf("我是一名歌手,叫做%s,《%s》就是我的作品\n", s.name, s.song)
}

func InterfaceStudyMain() {

	s1 := Star{name: "迪丽热巴", age: 18}
	Person(s1).say()

	s2 := Singer{name: "杨幂", song: "爱的供养"}
	Person(s2).say()
}
