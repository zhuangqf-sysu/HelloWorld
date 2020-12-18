package main

import (
	"HelloWorld/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	server, close, err := InitServer("configs/config.yaml")
	if err != nil {
		panic(err)
	}
	defer close()

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return listenSign(ctx, server)
	})
	group.Go(func() error {
		return server.Start()
	})

	if err := group.Wait(); err != nil {
		log.Printf("app stop for err(%+v)\n", err)
		return
	}
	log.Println("app stop")
}

func listenSign(ctx context.Context, server *server.Server) error {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("panic in signal proc, err: %v, stack: %s\n", r, buf)
		}
		// 优雅关停
		server.Stop()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case s := <-c:
			log.Printf("app get a signal %s\n", s.String())
			return nil
		case <-ctx.Done():
			// errgroup 中有线程报错退出，则全部关闭
			log.Printf("app context done \n")
			return nil
		}
	}
}
