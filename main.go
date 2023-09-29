package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-experiments-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-experiments-manager/data"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/exps", v1.GetExps)
	r.POST("/exp", v1.CreateExp)

	return r
}

func main() {
	data.SetupRedisClient()
	r := setupRouter()
	r.Run(":8082")
}
