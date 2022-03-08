/*
@Time   : 2022/3/3 9:32
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRouter(r *gin.Engine) {
	//创建请求
	r.GET("/api/:id/visitor", func(c *gin.Context) {
		responseCustom(c)
	})
	r.POST("/api/:id/visitor", func(c *gin.Context) {
		responseCustom(c)
	})
	r.PUT("/api/:id/visitor", func(c *gin.Context) {
		responseCustom(c)
	})
	r.DELETE("/api/:id/visitor", func(c *gin.Context) {
		responseCustom(c)
	})
}

func responseCustom(c *gin.Context) {
	var message string = "成功"
	var code int = 200
	//var aa string = "data"

	if c.Errors.String() != "" {
		message = fmt.Sprintf("Fail :%+v", c.Errors.String())
		code = 440
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
