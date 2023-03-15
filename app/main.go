package main

import (
	"log"
	"strconv"

	"github.com/ajinori-256/distribu/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	log.Print("Hello World Distri部！！！")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) { c.JSON(200, "Hello World Distri部！！！") })
	r.GET("/test/post", func(c *gin.Context) {
		database.AddMessage(25, "はろわ", "0.0.0.0")
		c.JSON(200, "Hello World Distri部！！！")
	})
	r.GET("/test/post/:number", func(c *gin.Context) {
		number, _ := strconv.ParseInt(c.Param("number"), 10, 64)
		message, err := database.GetMessage(number)
		if err != nil {
			c.JSON(404, "error")
		} else {
			c.JSON(200, message.ToJson())
		}
	})
	r.GET("/guilds/:guilds-id", func(c *gin.Context) { c.JSON(200, "Hello World Distri部！！！") })
	r.Run(":3000")
}
