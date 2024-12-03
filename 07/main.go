package main

import "fmt"

// defer是在当前函数的声明周期结束之后才执行的
// defer是压栈的方式，即先defer的后执行
// return 之后才会调用defer的语句
func fun() int {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	return 0
}

func main() {
	var a int = fun()
	fmt.Println("a=", a)
}
