package routes

import (
	"githubmbarkhanBusTracker/auth/database"
	"githubmbarkhanBusTracker/auth/handlers"
	"githubmbarkhanBusTracker/auth/repositories"
	"githubmbarkhanBusTracker/auth/services"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {

	city_Handler := handlers.NewCityHandler(services.NewCityService(repositories.NewCitiesRepository(database.DB)))
	apiGroup := r.Group("/position")
	apiGroup.POST("/city", city_Handler.Create)
	apiGroup.GET("/city/:id", city_Handler.Read)
	apiGroup.PUT("/city", city_Handler.Update)
	apiGroup.DELETE("/city/:id", city_Handler.Delete)
	apiGroup.GET("/city", city_Handler.List)

	company_Handler := handlers.NewCompanyHandler(services.NewCompanyService(repositories.NewCompanyRepository(database.DB)))
	apiGroup = r.Group("cmp")
	apiGroup.POST("/", company_Handler.Create)
	apiGroup.GET("/:id", company_Handler.Read)
	apiGroup.PUT("/", company_Handler.Update)
	apiGroup.DELETE("/:id", company_Handler.Delete)
	apiGroup.GET("/", company_Handler.List)

	province_Handler := handlers.NewProvinceHandler(services.NewProvinceService(repositories.NewProvincesRepository(database.DB)))
	apiGroup = r.Group("/position")
	apiGroup.POST("/province", province_Handler.Create)
	apiGroup.GET("/province/:id", province_Handler.Read)
	apiGroup.PUT("/province", province_Handler.Update)
	apiGroup.DELETE("/province/:id", province_Handler.Delete)
	apiGroup.GET("/province", province_Handler.List)

	tariff_Handler := handlers.NewTarrrifHandler(services.NewTarrifService(repositories.NewTariffsRepository(database.DB)))
	apiGroup = r.Group("/tariffs")
	apiGroup.POST("/", tariff_Handler.Create)
	apiGroup.GET("/:id", tariff_Handler.Read)
	apiGroup.PUT("/", tariff_Handler.Update)
	apiGroup.DELETE("/:id", tariff_Handler.Delete)
	apiGroup.GET("/", tariff_Handler.List)

	userHandler := handlers.NewUserHandler(services.NewUserService(repositories.NewUserRepository(database.DB)))
	apiGroup = r.Group("/user")
	apiGroup.POST("/", userHandler.Create)
	apiGroup.GET("/:id", userHandler.Read)
	apiGroup.PUT("/", userHandler.Update)
	apiGroup.DELETE("/:id", userHandler.Delete)
	apiGroup.GET("/", userHandler.List)
	apiGroup.POST("/login", userHandler.Login)

	userCredit_Handler := handlers.NewUserCreditHandler(services.NewUserCreditService(repositories.NewUserCreditRepository(database.DB)))
	apiGroup = r.Group("/user")
	apiGroup.POST("/usrCredit", userCredit_Handler.Create)
	apiGroup.GET("/usrCredit/:id", userCredit_Handler.Read)

	vehicle_Handler := handlers.NewVehicleHandler(services.NewvehicleService(repositories.NewVehicleRepository(database.DB)))
	apiGroup = r.Group("/vehicle")
	apiGroup.POST("/", vehicle_Handler.Create)
	apiGroup.GET("/:id", vehicle_Handler.Read)
	apiGroup.PUT("/", vehicle_Handler.Update)
	apiGroup.DELETE("/:id", vehicle_Handler.Delete)
	apiGroup.GET("/", vehicle_Handler.List)

	vehicleLog_Handler := handlers.NewVehicleLogHandler(services.NewVehicleLogService(repositories.NewVehicleLogRepository(database.DB)))
	apiGroup = r.Group("/vehicle")
	apiGroup.POST("/log", vehicleLog_Handler.Create)
	apiGroup.GET("/log/:id", vehicleLog_Handler.Read)
	apiGroup.PUT("/log", vehicleLog_Handler.Update)
	apiGroup.DELETE("/log/:id", vehicleLog_Handler.Delete)
	apiGroup.GET("/log", vehicleLog_Handler.List)

}
