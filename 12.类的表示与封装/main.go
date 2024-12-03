//类名，属性名，方法名...首字母大写表示对外开放（其他包可以访问）

package main

import "fmt"

type Man struct {
	name string
	age  int
}

// 不加指针是值传递
func (man *Man) setName(name string) {
	man.name = name
}

// 需要加*，指针传递
func (man *Man) getName() string {
	return man.name
}

func main() {
	man := Man{
		name: "huang",
		age:  15,
	}
	fmt.Println(man.getName())
	man.setName("zhangsan")
	fmt.Println(man.getName())
}
