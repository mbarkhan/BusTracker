package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type vehicleHandler struct {
	service services.VehicleServiceInt
}

func NewVehicleHandler(service services.VehicleServiceInt) VehicleHandlerInt {
	return &vehicleHandler{service}
}

func (r *vehicleHandler) Create(c *gin.Context) {
	var vhcle models.Vehicles
	err := c.BindJSON(&vhcle)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	vhcle, err = r.service.Create(&vhcle)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": vhcle})
}

func (r *vehicleHandler) Read(c *gin.Context) {
	vhclId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	vhcl := models.Vehicles{}

	vhcl, err = r.service.Read(vhclId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": vhcl})
}

func (r *vehicleHandler) Update(c *gin.Context) {
	vhcl := models.Vehicles{}

	err := c.BindJSON(&vhcl)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	vhcl, err = r.service.Update(&vhcl)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, vhcl)
}

func (r *vehicleHandler) Delete(c *gin.Context) {
	vhclId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(vhclId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully."})
}

func (r *vehicleHandler) List(c *gin.Context) {
	vehicles, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, vehicles)
}
