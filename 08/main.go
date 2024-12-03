package main

import (
	"fmt"
)

func fun(shuzu []int){
	fmt.Println("len of shuzu : ",len(shuzu))
}

func main() {
	// 切片 (变长数组), 当函数参数时是引用传递
	a:=[]int{1,2,3}
	for index,value:=range a{
		fmt.Println("index = ",index," ,value = ",value)
	}
	fun(a)

	var b [4]int
	for index :=0;index<len(b);index++{
		fmt.Println(b[index])
	}

	c:=[10]int{10,20,30,40}
	for i:=0;i<len(c);i++{
		fmt.Println(c[i])
	}

	fmt.Printf("type of a %T\n",a)
	fmt.Printf("type of b %T\n",b)
	fmt.Printf("type of c %T\n",c)
}