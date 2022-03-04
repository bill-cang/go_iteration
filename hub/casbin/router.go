/*
@Time   : 2022/3/3 9:32
@Author : ckx0709
@Remark :
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRouter(r *gin.Engine) {
	//创建请求
	r.GET("/api/:id/visitor", func(c *gin.Context) {
		resultSuccess(c)
	})
	r.POST("/api/:id/visitor", func(c *gin.Context) {
		resultSuccess(c)
	})
	r.PUT("/api/:id/visitor", func(c *gin.Context) {
		resultSuccess(c)
	})
	r.DELETE("/api/:id/visitor", func(c *gin.Context) {
		resultSuccess(c)
	})
}

func resultSuccess(c *gin.Context) {
	var message string = "成功"
	var code int = 200
	var aa string = "data"
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    aa,
		"result":  "true",
	})
}
