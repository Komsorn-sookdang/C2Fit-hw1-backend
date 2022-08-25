package router

import (
	"c2fit-hw-backend/controller"
	"c2fit-hw-backend/session"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("websession", session.GetStore()))

	r.POST("/register", controller.CreateUser)

	r.POST("/login", controller.UserLogin)

	r.GET("/testdeploy", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "This API for test deployment."})
	})

	r.GET("/whoami", controller.WhoAmI)

	r.GET("/logout", controller.UserLogOut)

	return r
}
