package main

import (
	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/cookie"
  	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/set", func(c *gin.Context) {
		session := sessions.Default(c)

		var count int
		v := session.Get("count")
		if v == nil {
			count = 69
		} else {
			count = v.(int)
		}

		session.Set("count", count)
		session.Save()

		c.JSON(200, gin.H{"count": count})
	})

	router.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)

		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
		}

		c.JSON(200, gin.H{"count": count})
	})

	router.Run(":8080")

}