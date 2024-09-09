package study

import (
	"fmt"
	"math/rand"
)

/*
按照约定，包名与导入路径的最后一个元素相同。
例如，"math/rand" 包中的源码均以`package`rand` 语句开始.

*/

func PackageStudyMain() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
