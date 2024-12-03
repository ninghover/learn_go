package main

import "fmt"

func main() {
	//方法一：声明一个变量，默认值是0,""
	var a int
	fmt.Println("a = ", a)
	var b string
	fmt.Println("b = ", b)

	//方法二：声明一个变量，初始化一个值
	var c int = 100
	var d string = "hello"
	fmt.Printf("%v, type of c = %T\n", c, c) //%v是默认格式，%d是int，%s是字符串
	fmt.Printf("%v, type of d = %T\n", d, d)

	//方法三：在初始化的时候省略数据类型
	var e = 200
	var f = "hello"
	fmt.Printf("%d, type of e = %T\n", e, e)
	fmt.Printf("%s, type of f = %T\n", f, f)

	//方法四：:=（常用方法，但不能用来声明全局变量）
	g := 500
	fmt.Println(g)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx, yy)

	var aa, bb = "123", "456"
	fmt.Println(aa, bb) //有一个println函数和print函数
	println(aa, bb)
	var (
		cc int  = 500
		dd bool = false
	)
	fmt.Printf("%v, type of dd = %T\n", cc, cc)
	fmt.Printf("%v, type of dd = %T\n", dd, dd)

}
