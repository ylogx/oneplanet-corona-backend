package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)

	r.GET("/home", GetDataForHomepage)

	return r
}
