// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int) //无缓冲管道
// 	fmt.Println(len(c), cap(c))
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		defer fmt.Println("子协程执行完毕...") //defer在最后执行，甚至在return之后
// 		fmt.Println("子协程正在运行...")
// 		c <- 666 // 阻塞操作，会等待主协程从通道接收数据
// 	}()
// 	fmt.Println("主协程开始执行...")
// 	num := <-c //阻塞操作，会等待其他协程写数据
// 	fmt.Printf("主协程读到的数: %v\n", num)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int, 3) //创建一个带缓冲的管道
// 	fmt.Printf("管道容量：%v, 当前长度：%v\n", cap(c), len(c))
// 	// 管道未满，可以直接写，管道满了再写会阻塞
// 	// 管道有数据可以直接读，管道没数据读会阻塞
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		for i := 0; i < 4; i++ {
// 			fmt.Printf("写数据 i = %v\n", i)
// 			c <- i
// 		}
// 	}()
// 	for i := 0; i < 4; i++ {
// 		num := <-c
// 		fmt.Println("读到的数据: ", num)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int, 5)
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 			fmt.Println("发送数据: ", i)
// 		}
// 		close(c)
// 		// 关闭channel后，无法向channel再发数据(panic)
// 		// 关闭channel后，可以继续接收数据
// 	}()

// 	time.Sleep(2 * time.Second)
// 	// for i := 0; i < 6; i++ {
// 	// 	num, ok := <-c
// 	// 	fmt.Println(ok)
// 	// 	fmt.Println("读到了数据: ", num)
// 	// }
// 	// 上述代码可以简写为如下
// 	for num := range c {
// 		fmt.Println("读到了数据：", num)
// 	}
// 	time.Sleep(time.Second)
// }

package main

import (
	"fmt"
	"time"
)

func fun(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // 可写
			x = x + y
		case <-quit: // 可读
			fmt.Println("退出")
			return
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			time.Sleep(time.Second)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fun(c, quit)
}
