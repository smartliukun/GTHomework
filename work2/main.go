package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye bye ,shutdown the server")) // 没有输出
		cancel()
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	group.Go(func() error {
		// 接收退出信号,或者师远程请求关闭服务
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("ctx canceled"))
		case sign, ok := <-quit:
			fmt.Println(fmt.Sprintf("sign=%v,ok=%v", sign, ok))
		}

		fmt.Println("closing server ...")
		return server.Shutdown(context.Background())
	})

	group.Go(func() error {
		fmt.Println("start the http server ")
		err := server.ListenAndServe()
		return err
	})

	group.Wait()

	fmt.Println(fmt.Sprintf("errGroup is stoped,reason is %v", errCtx.Err()))
	fmt.Println("all goroutine stop ")

}

// 常规请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// 关闭http请求
