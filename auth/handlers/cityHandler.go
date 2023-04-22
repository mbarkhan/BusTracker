package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CityHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type cityHandler struct {
	service services.CityServiceInt
}

func NewCityHandler(service services.CompanyUserServiceInt) CityHandlerInt {
	return &cityHandler{service}
}

func (r *cityHandler) Create(c *gin.Context) {
	var item models.City
	err := c.BindJSON(&cmpusr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	cmpusr, err = r.service.Create(&cmpusr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": cmpusr})
}

func (r *cityHandler) Read(c *gin.Context) {
	cmpusrId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	cmpusr := models.Company_Users{}

	cmpusr, err = r.service.Read(cmpusrId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": cmpusr})
}

func (r *cityHandler) Update(c *gin.Context) {
	cmpusr := models.Company_Users{}

	err := c.BindJSON(&cmpusr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	cmpusr, err = r.service.Update(&cmpusr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cmpusr)
}

func (r *cityHandler) Delete(c *gin.Context) {
	cmpusrId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(cmpusrId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Company_User deleted successfully."})
}

func (r *cityHandler) List(c *gin.Context) {
	cmpusr, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cmpusr)
}
