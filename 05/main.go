// main包中：import pkg1 -> const -> var -> init() -> main() -> exit
//
//	|
//
// import pkg2 -> const -> var -> init()
package main

import (
	"fmt"
	_ "learn_go/05/lib1"     // 如果导入了一个包的，但是又没有调用这个包的方法，会编译不过。此时可以_匿名，会调用这个包的init函数
	mylib "learn_go/05/lib2" // 给这个包取别名
)

func main() {
	// lib1.Fun()
	mylib.Fun()
	fmt.Println(mylib.A)
}
