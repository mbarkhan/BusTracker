package routes

import (
	"githubmbarkhanBusTracker/auth/database"
	"githubmbarkhanBusTracker/auth/handlers"
	"githubmbarkhanBusTracker/auth/repositories"
	"githubmbarkhanBusTracker/auth/services"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	apiGroup := r.Group("/user")

	/*company_user_Handler := handlers.NewCompanyUserHandler(services.NewCompanyUserService(repositories.NewCompanyUserRepository(database.DB)))
	apiGroup.POST("/cmpusr", company_user_Handler.Create)
	apiGroup.GET("/cmpusr/:id", company_user_Handler.Read)
	apiGroup.PUT("/cmpusr", company_user_Handler.Update)
	apiGroup.DELETE("/cmpusr/:id", company_user_Handler.Delete)
	apiGroup.GET("/cmpusr", company_user_Handler.List)
	*/

	company_Handler := handlers.NewCompanyHandler(services.NewCompanyService(repositories.NewCompanyRepository(database.DB)))
	apiGroup = r.Group("/user")
	apiGroup.POST("/cmp", company_Handler.Create)
	apiGroup.GET("/cmp/:id", company_Handler.Read)
	apiGroup.PUT("/cmp", company_Handler.Update)
	apiGroup.DELETE("/cmp/:id", company_Handler.Delete)
	apiGroup.GET("/cmp", company_Handler.List)

	/*pathes_Handler := handlers.NewPathesHandler(services.NewPathService(repositories.NewPathesRepository(database.DB)))
	apiGroup = r.Group("/pathes")
	apiGroup.POST("/path", pathes_Handler.Create)
	apiGroup.GET("/path/:id", pathes_Handler.Read)
	apiGroup.PUT("/path", pathes_Handler.Update)
	apiGroup.DELETE("/path/:id", pathes_Handler.Delete)
	apiGroup.GET("/path", pathes_Handler.List)
	*/
	/*	psn_path_Handler := handlers.NewPathesHandler(services.NewPathService(repositories.NewPathesRepository(database.DB)))
		apiGroup = r.Group("/pathes")
		apiGroup.POST("/psnpath", psn_path_Handler.Create)
		apiGroup.GET("/psnpath/:id", psn_path_Handler.Read)
		apiGroup.PUT("/psnpath", psn_path_Handler.Update)
		apiGroup.DELETE("/psnpath/:id", psn_path_Handler.Delete)
		apiGroup.GET("/psnpath", psn_path_Handler.List)
	*/
	userHandler := handlers.NewUserHandler(services.NewUserService(repositories.NewUserRepository(database.DB)))
	apiGroup = r.Group("/user")
	apiGroup.POST("/", userHandler.Create)
	apiGroup.GET("/:id", userHandler.Read)
	apiGroup.PUT("/", userHandler.Update)
	apiGroup.DELETE("/:id", userHandler.Delete)
	apiGroup.GET("/", userHandler.List)
	apiGroup.POST("/login", userHandler.Login)

	vehicleHandler := handlers.NewVehicleHandler(services.NewvehicleService(repositories.NewVehicleRepository(database.DB)))
	apiGroup = r.Group("/user")
	apiGroup.POST("/vehicle", vehicleHandler.Create)
	apiGroup.GET("/vehicle/:id", vehicleHandler.Read)
	apiGroup.PUT("/vehicle", vehicleHandler.Update)
	apiGroup.DELETE("/vehicle/:id", vehicleHandler.Delete)
	apiGroup.GET("/vehicle", vehicleHandler.List)

}
