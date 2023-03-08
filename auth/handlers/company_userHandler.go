package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyUserHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type companyUserHandler struct {
	service services.CompanyUserServiceInt
}

func NewCompanyUserHandler(service services.CompanyUserServiceInt) CompanyUserHandlerInt {
	return &companyUserHandler{service}
}

func (r *companyUserHandler) Create(c *gin.Context) {
	var cmpusr models.Company_Users
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

func (r *companyUserHandler) Read(c *gin.Context) {
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

func (r *companyUserHandler) Update(c *gin.Context) {
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

func (r *companyUserHandler) Delete(c *gin.Context) {
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

func (r *companyUserHandler) List(c *gin.Context) {
	cmpusr, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, cmpusr)
}
