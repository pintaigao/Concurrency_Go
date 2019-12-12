package main

import (
	"fmt"
	"time"
)

func coroutine() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 没有go的话，就卡在i=0上面了
		go func(i int) {
			for {
				a[i]++
				// 手动交出控制权
				// runtime.Gosched()
			}
		}(i)
		fmt.Print(a)
	}
	time.Sleep(time.Millisecond)
	fmt.Print(a)
}

func main() {
	/* for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute) */

	coroutine()
}
