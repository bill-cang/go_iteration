package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*https://cloud.tencent.com/developer/article/1847917#:~:text=%E6%A6%82%E8%A6%81%20Go%20%E5%B9%B6%E5%8F%91%E7%9B%B8%E5%85%B3%E5%BA%93%20sync%20%E9%87%8C%E9%9D%A2%E6%9C%89%E4%B8%80%E4%B8%AA%E6%9C%89%E8%B6%A3%E7%9A%84%20package%20Pool%EF%BC%8C,sync.Pool%20%E6%98%AF%E4%B8%AA%E6%9C%89%E8%B6%A3%E7%9A%84%E5%BA%93%EF%BC%8C%E7%94%A8%E5%BE%88%E5%B0%91%E7%9A%84%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0%E4%BA%86%E5%BE%88%E5%B7%A7%E7%9A%84%E5%8A%9F%E8%83%BD%E3%80%82%20%E7%AC%AC%E4%B8%80%E7%9C%BC%E7%9C%8B%E5%88%B0%20Pool%20%E8%BF%99%E4%B8%AA%E5%90%8D%E5%AD%97%EF%BC%8C%E5%B0%B1%E8%AE%A9%E4%BA%BA%E6%83%B3%E5%88%B0%E6%B1%A0%E5%AD%90%EF%BC%8C%20%E5%85%83%E7%B4%A0%E6%B1%A0%E5%8C%96%E6%98%AF%E5%B8%B8%E7%94%A8%E7%9A%84%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%E7%9A%84%E6%89%8B%E6%AE%B5%20%EF%BC%88%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%E7%9A%84%E5%87%A0%E6%8A%8A%E6%96%A7%E5%A4%B4%EF%BC%9A%E5%B9%B6%E5%8F%91%EF%BC%8C%E9%A2%84%E5%A4%84%E7%90%86%EF%BC%8C%E7%BC%93%E5%AD%98%EF%BC%89%E3%80%82
 */

// 用来统计实例真正创建的次数
var numCalcsCreated int32

// 创建实例的函数
func createBuffer() interface{} {
	// 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	// 创建实例
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// 多 goroutine 并发测试
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// 申请一个 buffer 实例
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			// 释放一个 buffer 实例
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}
