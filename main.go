package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo-example/internal/app"
	"go-echo-example/internal/global"
	"go-echo-example/internal/routes"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	app.Init()
	e := echo.New()
	app.SetLogger(e)
	app.SetMiddleware(e)
	routes.RegisterRoutes(e)
	go func() {
		fmt.Println(global.Config.App.Addr)
		if err := e.Start(global.Config.App.Addr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), global.Config.App.Timeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
