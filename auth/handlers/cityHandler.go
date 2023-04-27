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

func NewCityHandler(service services.CityServiceInt) CityHandlerInt {
	return &cityHandler{service}
}

func (r *cityHandler) Create(c *gin.Context) {
	var item models.City
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

func (r *cityHandler) Read(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	item := models.City{}

	item, err = r.service.Read(itemId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": item})
}

func (r *cityHandler) Update(c *gin.Context) {
	item := models.City{}

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

func (r *cityHandler) Delete(c *gin.Context) {
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

func (r *cityHandler) List(c *gin.Context) {
	item, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}
