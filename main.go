package main

import (
	"fmt"
	"go-study/study"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}

func main() {
	//go printMessage("Hello from Goroutine")
	//
	//printMessage("Hello from Main")
	//study.MapStudy()
	//study.FuncStudyMain()
	//s1 := []int{1, 2, 3}
	//fmt.Println(s1)
	//s1 = append(s1, 4, 5)
	//fmt.Println(s1)
	//
	//defer printMessage("hello")
	//// 多个 defer 逆序执行
	//defer printMessage("world")
	//study.PointerStudyMain()
	//study.SliceStudyMain()
	//study.RangeStudyMain()
	//fmt.Println(study.RecursionStudyMain(10))
	//study.ErrorStudyMain()
	study.GoRoutineStudyMain()
}
