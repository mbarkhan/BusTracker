package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleLogHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type vehicleLogHandler struct {
	service services.VehicleLogServiceInt
}

func NewVehicleLogHandler(service services.VehicleLogServiceInt) VehicleLogHandlerInt {
	return &vehicleLogHandler{service}
}

func (r *vehicleLogHandler) Create(c *gin.Context) {
	var item models.VehicleLog
	err := c.BindJSON(&item)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	item, err = r.service.Create(&item)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": item})
}

func (r *vehicleLogHandler) Read(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	item := models.VehicleLog{}

	item, err = r.service.Read(itemId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": item})
}

func (r *vehicleLogHandler) Update(c *gin.Context) {
	item := models.VehicleLog{}

	err := c.BindJSON(&item)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	item, err = r.service.Update(&item)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func (r *vehicleLogHandler) Delete(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(itemId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "City deleted successfully."})
}

func (r *vehicleLogHandler) List(c *gin.Context) {
	item, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}
