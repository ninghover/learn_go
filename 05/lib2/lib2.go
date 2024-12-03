package lib2

import "fmt"

func init(){
	fmt.Println("lib2.init()...")
}

func Fun(){
	fmt.Println("lib2.Fun()...")
}

var A = 20

// 变量和方法名大写时其他文件才能访问