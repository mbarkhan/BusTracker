package routes

import (
	"githubmbarkhanBusTracker/auth/database"
	"githubmbarkhanBusTracker/auth/handlers"
	"githubmbarkhanBusTracker/auth/repositories"
	"githubmbarkhanBusTracker/auth/services"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	userHandler := handlers.NewUserHandler(services.NewUserService(repositories.NewUserRepository(database.DB)))
	apiGroup := r.Group("/auth")
	apiGroup.POST("/user", userHandler.Create)
	apiGroup.GET("/user/:id", userHandler.Read)
	apiGroup.PUT("/user", userHandler.Update)
	apiGroup.DELETE("/user/:id", userHandler.Delete)
	apiGroup.GET("/user", userHandler.List)

}
