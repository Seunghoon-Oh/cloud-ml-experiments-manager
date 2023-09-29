package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-experiments-manager/service"
)

func GetFoos(c *gin.Context) {
	data := service.GetFoos()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateFoo(c *gin.Context) {
	data := service.CreateFoo()
	c.String(http.StatusOK, data)
}
