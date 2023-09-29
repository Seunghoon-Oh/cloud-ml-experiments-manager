package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-experiments-manager/service"
)

func GetExps(c *gin.Context) {
	data := service.GetExps()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateExp(c *gin.Context) {
	data := service.CreateExp()
	c.String(http.StatusOK, data)
}
