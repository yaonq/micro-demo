package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	engine := gin.New()

	api := engine.Group("/micro/web")
	{
		api.GET("/gin", func(c *gin.Context) {
			name := c.Query("name")

			logger.Info("name:", name)

			c.String(200, "Micro Web API: hello %s\n", name)
		})
	}

	srv := web.NewService(
		web.Address(":8001"),
		web.Handler(engine),
	)

	srv.Run()

}
