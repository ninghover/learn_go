package main

import "fmt"

func fun(m map[string]int) {
	// map 作为参数传递时引用传递
	// 即函数内修改会影响到外面
	for key, value := range m {
		fmt.Printf("key=%s,value=%d\n", key, value)
	}
}

func main() {
	// m:=make(map[string]int)	//key是string，value是int

	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	// 增
	m["c"] = 3

	// 删
	delete(m, "a")

	// 改
	m["b"] = 20
	// 查
	for key, value := range m {
		fmt.Println("key=", key, ",value=", value)
	}

	mCopy := m //这是浅拷贝，实际上指向的时同一个内存，对任意一个的修改会影响到另一个
	mCopy["d"] = 4
	fun(m)
	// 如果不想浅拷贝，需要自己手动实现一个函数，依次插入

	fun(m)
}
