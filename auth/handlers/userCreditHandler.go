package handlers

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCreditHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
}

type userCreditHandler struct {
	service services.UserCreditServiceInt
}

func NewUserCreditHandler(service services.UserCreditServiceInt) UserCreditHandlerInt {
	return &userCreditHandler{service}
}

func (r *userCreditHandler) Create(c *gin.Context) {
	var item models.UserCredit
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

func (r *userCreditHandler) Read(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	items := []models.UserCredit{}

	items, err = r.service.Read(uint(itemId))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": items})
}
