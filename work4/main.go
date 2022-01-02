package main

import (
	"container/ring"
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	limitCount             = 20 // 限频
	limitWindow            = 6  // 滑动窗口个数,一秒一个，总共6秒窗口统计口径
	curCount    int32      = 0  // 记录限频数量
	head        *ring.Ring      // 环形队列（链表）
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	// 初始化滑动窗口
	head = ring.New(limitWindow)
	for i := 0; i < limitWindow; i++ {
		head.Value = 0
		head = head.Next()
	}
	// 启动执行器
	go func() {
		timer := time.NewTicker(time.Second * 1)
		for range timer.C { // 定时每隔1秒刷新一次滑动窗口数据
			subCount := int32(0 - head.Value.(int))
			newCount := atomic.AddInt32(&curCount, subCount)

			arr := [6]int{}
			for i := 0; i < limitWindow; i++ { // 打印窗口
				arr[i] = head.Value.(int)
				head = head.Next()
			}
			fmt.Println("move subCount,newCount,arr", subCount, newCount, arr)
			head.Value = 0
			head = head.Next()
		}
	}()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

// 常规请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	n := atomic.AddInt32(&curCount, 1)
	if n > int32(limitCount) { // 超出限频
		atomic.AddInt32(&curCount, -1) // add 1 by atomic，业务处理完毕，放回令牌
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("too many request"))
		fmt.Println("too many request")
	} else {
		mu := sync.Mutex{}
		mu.Lock()
		pos := head.Prev()
		pos.Value = pos.Value.(int) + 1
		mu.Unlock()
		time.Sleep(1 * time.Millisecond) // 模拟服务处理时间
		w.Write([]byte("hello world"))
	}

}

// 异常报错的处理
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
