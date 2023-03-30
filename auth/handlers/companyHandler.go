package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type companyHandler struct {
	service services.CompanyServiceInt
}

func NewCompanyHandler(service services.CompanyServiceInt) CompanyHandlerInt {
	return &companyHandler{service}
}

func (r *companyHandler) Create(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	company, err = r.service.Create(&company)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": company})
}

func (r *companyHandler) Read(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	company := models.Company{}

	company, err = r.service.Read(companyId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": company})
}

func (r *companyHandler) Update(c *gin.Context) {
	company := models.Company{}

	err := c.BindJSON(&company)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	company, err = r.service.Update(&company)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, company)
}

func (r *companyHandler) Delete(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(companyId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Company deleted successfully."})
}

func (r *companyHandler) List(c *gin.Context) {
	companys, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, companys)
}
