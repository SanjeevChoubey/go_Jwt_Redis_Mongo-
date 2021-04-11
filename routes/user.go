package routes

import (
	"github.com/SanjeevChoubey/go_Jwt_Redis_Mongo/controller"
	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
