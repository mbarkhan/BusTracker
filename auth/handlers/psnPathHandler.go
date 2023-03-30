// To Do
// to read a record in this model, we may pass userID and PathID. So, survey the method again
// Delete method should be checked too

package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PsnPathHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type psnPathHandler struct {
	service services.PassengerPathServiceInt
}

func NewPsnPathHandler(service services.PassengerPathServiceInt) PsnPathHandlerInt {
	return &psnPathHandler{service}
}

func (r *psnPathHandler) Create(c *gin.Context) {
	var psnpath models.Passenger_Path
	err := c.BindJSON(&psnpath)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	psnpath, err = r.service.Create(&psnpath)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": psnpath})
}

func (r *psnPathHandler) Read(c *gin.Context) {
	psnPathId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user := models.Passenger_Path{}

	user, err = r.service.Read(psnPathId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}

func (r *psnPathHandler) Update(c *gin.Context) {
	psnpath := models.Passenger_Path{}

	err := c.BindJSON(&psnpath)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	psnpath, err = r.service.Update(&psnpath)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, psnpath)
}

func (r *psnPathHandler) Delete(c *gin.Context) {
	psnpathId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(psnpathId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Path for the Passenger deleted successfully."})
}

func (r *psnPathHandler) List(c *gin.Context) {
	psnpathes, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, psnpathes)
}
