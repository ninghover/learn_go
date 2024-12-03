package main

import "fmt"

// 用const来做枚举
const(
	A=10
	B=20
)
// iota在每个const块里的第一行为0，逐行自增，下面每一行都和上面最近的公式一样
const (
	C=iota+1	// 1
	D			// 2
	E			// 3
	F=iota*2	// 6
	G			// 8
)

const (
	H = iota
)

func main(){
	const a=20	// var定义的变量不使用会报错，const定义的常量不使用只会报警告
	fmt.Println("a = ",a)
	fmt.Println("A = ",A)
	fmt.Println("B = ",B)
	fmt.Println(C,D,E,F,G,H)
}