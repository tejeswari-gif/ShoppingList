package routes

import (
	"shoppinglist/controller"

	"github.com/gin-gonic/gin"
)

var (
	signupController = new(controller.SignupController)
)

//InitRoutes defines the UserGroup
func InitRoutes(r *gin.Engine) {
	UserGroup(r)
}
