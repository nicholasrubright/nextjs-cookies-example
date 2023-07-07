package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetTokenInSession(session sessions.Session, token string) {
	session.Set("token", token)
	session.Save()
}

func GetTokenFromSession(session sessions.Session) (string, error)  {

	var token string

	tokenVal := session.Get("token")

	if tokenVal == nil {
		token = "BIG FUCKING TOKEN AHAHAHAHAHAHAHA"
	} else {
		token = tokenVal.(string)
	}

	return token, nil
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		sessionToken, err := GetTokenFromSession(session)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		c.Set("token", sessionToken)

		c.Next()
	}
}


func main() {

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := router.Group("/api")

	apiRoutes.Use(SessionMiddleware())
	{
		router.POST("/set", func(c *gin.Context) {
			session := sessions.Default(c)
	
			var token string
			v := session.Get("token")
			if v == nil {
				token = "BIG FUCKING TOKEN AAAAHHH"
			} else {
				token = v.(string)
			}
	
			SetTokenInSession(session, token)
	
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
	
			token, err := GetTokenFromSession(session)
	
			if err != nil {
				panic("balls")
			}
	
			c.JSON(200, gin.H{"token": token})
		})
	}


	router.Run(":8080")

}