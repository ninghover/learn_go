// // 子协程是基于主协程的，主协程推出，则子协程也会退出。子协程退出不会影响到主协程
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task() {
// 	i := 0
// 	for {
// 		fmt.Printf("子协程 i = %v\n", i)
// 		i++
// 		time.Sleep(1 * time.Second)
// 		// if i==5{
// 		// 	break
// 		// }
// 	}
// }

// func main() {
// 	go task()
// 	i := 0
// 	for {
// 		fmt.Printf("主协程 i = %v\n", i)
// 		i++
// 		time.Sleep(1 * time.Second)
// 		// if i==5{
// 		// 	break
// 		// }
// 	}
// }

package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("这是一个子协程")
		func(a, b int) {
			// return           //这个return只是退出了子协程中的匿名函数
			runtime.Goexit() //如果想要推出子协程
			fmt.Printf("a = %v, b = %v , add = %v\n", a, b, a+b)
		}(1, 2)
	}()
	// for{}
}
