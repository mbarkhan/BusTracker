package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PathHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type pathHandler struct {
	service services.PathesServiceInt
}

func NewPathesHandler(service services.PathesServiceInt) PathHandlerInt {
	return &pathHandler{service}
}

func (r *pathHandler) Create(c *gin.Context) {
	var path models.Pathes
	err := c.BindJSON(&path)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	path, err = r.service.Create(&path)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": path})
}

func (r *pathHandler) Read(c *gin.Context) {
	pathID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	path := models.Pathes{}

	path, err = r.service.Read(pathID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": path})
}

func (r *pathHandler) Update(c *gin.Context) {
	path := models.Pathes{}

	err := c.BindJSON(&path)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	path, err = r.service.Update(&path)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, path)
}

func (r *pathHandler) Delete(c *gin.Context) {
	pathID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(pathID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Path deleted successfully."})
}

func (r *pathHandler) List(c *gin.Context) {
	pathes, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, pathes)
}
