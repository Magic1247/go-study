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
	study.PtrStudy()

}
