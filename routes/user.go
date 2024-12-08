package routes

import (
	"batkor/restapp/models/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.POST("/api/user/create", func(c *gin.Context) {
		values := make(map[string]string)
		values["login"] = c.PostForm("login")
		values["email"] = c.PostForm("email")
		User := user.New(values)
		User.Save()
		c.JSON(200, gin.H{
			"id": User.Id(),
		})
	})

	routes.GET("/api/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api/user",
		})
	})
}
