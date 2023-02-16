package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"uni_Go/src/utils"
)

func main() {

	r := gin.Default()
	data := utils.SelectData("select * from data")
	r.Use(Cors())
	r.GET("/user", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"Data": data,
		})
	})
	r.POST("/data/id", func(c *gin.Context) {
		id := c.Query("id")
		getdata := utils.SelectData("select * from data where id = " + id)
		fmt.Println(getdata)
		fmt.Println(878747)
		c.JSON(http.StatusOK, gin.H{
			"data": getdata,
		})
	})
	r.Run(":9090")
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
