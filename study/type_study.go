package study

import (
	"fmt"
	"strconv"
)

// 字符串转整形
func strconvInt() {
	var s string = "42"
	var i, _ = strconv.Atoi(s) // string 转换为 int
	fmt.Println(i)
}

// 整形转字符串
func intConvStr() {
	var i int = 42
	var s string = strconv.Itoa(i) // int 转换为 string
	fmt.Println(s)
}

func floatConvInt() {
	var f float64 = 3.1415
	var s string = strconv.FormatFloat(f, 'f', 2, 64) // float64 转换为 string，保留 2 位小数
	fmt.Printf("%f保留两位小数转换为%s \n", f, s)
}

func strconvFloat() {
	var s string = "3.1415"
	var f, _ = strconv.ParseFloat(s, 64) // string 转换为 float64
	fmt.Printf("字符串%s转换为%f \n", s, f)
}

func boolConvString() {
	var s string = strconv.FormatBool(true) // bool 转换为 string
	fmt.Println(s)
}

func strconvBool() {
	var s string = "true"
	var f, _ = strconv.ParseBool(s)
	fmt.Println(f)
}

func TypeStudyMain() {
	strconvInt()
	intConvStr()
	floatConvInt()
	strconvFloat()
	boolConvString()
	strconvBool()
}
