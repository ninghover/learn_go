package  main

import "fmt"

func fun1(){
	fmt.Println("无参，无返回值函数")
}

func fun2(a int) int{
	fmt.Println("有参，有返回值函数, a = ",a)
	return a
}

func fun3(int)(int,string){
	fmt.Println("多返回值函数")
	return 10,"ok"
}

// 返回值匿名
func fun4()(r1 int,r2 int){
	r1=20
	r2=50
	return
}

// 返回值类型相同的一种省略形式
func fun5()(r1,r2 int,r3 string){
	r1,r2=200,500
	r3="hello"
	return
}

func main()  {
	fun1()
	fmt.Println(fun2(10))
	fmt.Println(fun3(20))
	fmt.Println(fun4())
	fmt.Println(fun5())
}