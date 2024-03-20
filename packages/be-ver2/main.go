package main

import (
	authController "namviek-backend/packages/be-ver2/controllers/auth"
	profileController "namviek-backend/packages/be-ver2/controllers/profile"
	"namviek-backend/packages/be-ver2/initializers"
	"namviek-backend/packages/be-ver2/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	v2 := r.Group("/api/v2")

	v2.POST("/sign-in", authController.GenJwtNRefresh)

	v2.GET("/profile/:uid", middlewares.VerifyAuth, profileController.Get)

	v2.POST("/profile", middlewares.VerifyAuth, profileController.Post)

	r.Run()
}
