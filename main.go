package main

import (
	"net/http"
	"os"

	"github.com/0x-buidl/simple-chat-app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pool := services.NewPool()
	go pool.Start()

	r.Any("/api/ws", func(c *gin.Context) { services.HandleWs(pool, c) })

	port, staticPath := initServer()
	r.LoadHTMLFiles(staticPath + "index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(port) // listen and serve on port :8080
}

func initServer() (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	port := "localhost:8080"
	staticPath := wd + "/frontend/"
	if gin.Mode() == gin.ReleaseMode {
		staticPath = wd + "/frontend/dist/"
		port = ":8080"
	}
	return port, staticPath
}
