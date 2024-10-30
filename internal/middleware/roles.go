package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func AdminRoleRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserSessionKey)
	fmt.Println("klsnfjlsndfjowneofewnioewnfinewionewiouneuwionfuoiewf")
	log.Print("user session", user)
	c.Next()
}
