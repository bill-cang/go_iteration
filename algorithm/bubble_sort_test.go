/*
@Time   : 2021/12/24 16:57
@Author : ckx0709
@Remark :
*/
package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	nubs := []int{5, 3, 2, 1, 9, 4, 7, 5, 8, 6, 3}
	sort := BubbleSort(nubs)
	fmt.Printf("%+v", sort)
}

func TestBubbleSort2(t *testing.T) {
	done := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(500 * time.Microsecond)
			fmt.Println(i)
		}

		fmt.Println("hello world")
		<-done
	}()
	fmt.Printf("****\n")
	done <- 1
}
