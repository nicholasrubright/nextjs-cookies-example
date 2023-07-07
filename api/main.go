package main

import (
	"fmt"

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

		var token string
		v := session.Get("token")
		if v == nil {
			token = "BIG FUCKING TOKEN AAAAHHH"
		} else {
			token = v.(string)
		}

		session.Set("token", token)
		session.Save()

		c.JSON(200, gin.H{"token": token})
	})

	router.GET("/get", func(c *gin.Context) {

		fmt.Println(c.Request.Header)

		cookie, err := c.Cookie("mysession")

		if err != nil {
			fmt.Println("NO COOKIE FOUND ON REQUEST: ", err)
		} else {
			fmt.Println("COOKEI HAS BEEN FOUND ON REQUET: ", cookie)
		}

		session := sessions.Default(c)

		var token string
		v := session.Get("token")
		if v == nil {
			token = "penis"
		} else {
			token = v.(string)
		}

		c.JSON(200, gin.H{"token": token})
	})

	router.Run(":8080")

}