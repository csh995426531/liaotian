package main

import (
	"context"
	"liaotian/test/ws"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendAllService()
	go ws.WebsocketManager.SendAllService()
	go ws.TestSendGroup()
	go ws.TestSendAll()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/:channel", ws.WebsocketManager.WsClient)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Error: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	log.Println("Server Shutdown")
}
