package study

import "fmt"

func run(name string) {
	fmt.Printf(" %s 在跑!\n", name)
}

func ArrayStudy() {
	a := [3]int{0, 1, 2}
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5, 65}
	fmt.Println(b)

	var c = new([10]int)
	c[5] = 3
	fmt.Println(c)

	zoom := [...]string{"小猫", "小狗"}

	for i, v := range zoom {
		fmt.Println(i, v)
		run(v)
	}
	//len 长度, cap 容量
	fmt.Println(len(zoom), cap(zoom))

	er := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}
	fmt.Println(er)

	s := [...]int{1, 2, 3}
	s1 := s[0:2]
	// 切片再append后再修改不影响原数组
	s1 = append(s1, 4, 5, 6)
	s1[0] = 1
	fmt.Println(s)
	fmt.Println(s1)

}
