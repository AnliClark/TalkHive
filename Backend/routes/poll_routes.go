package routes

import "github.com/gin-gonic/gin"
import "TalkHive/controllers"

func pollRoutes(r *gin.Engine) {
	r.POST("/polls", controllers.Poll)
}
