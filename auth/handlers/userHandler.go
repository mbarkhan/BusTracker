package handlers

import (
	"fmt"
	utils "githubmbarkhanBusTracker/auth/Utils"
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerInt interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	service services.UserServiceInt
}

func NewUserHandler(service services.UserServiceInt) UserHandlerInt {
	return &userHandler{service}
}

func (r *userHandler) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user, err = r.service.Create(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}

func (r *userHandler) Read(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user := models.User{}

	user, err = r.service.Read(userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}

func (r *userHandler) Update(c *gin.Context) {
	user := models.User{}

	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user, err = r.service.Update(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (r *userHandler) Delete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return

	}

	err = r.service.Delete(userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully."})
}

func (r *userHandler) List(c *gin.Context) {
	users, err := r.service.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
func (r *userHandler) Login(c *gin.Context) {
	mobnumber := c.Param("mobile")
	password := c.Param("password")

	user, err := r.service.GetUserByMobile(mobnumber)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if user == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Mobile Number Not Found"})
		return
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err}) //password Invalid
		return
	}

	token, err := utils.NewToken(user.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err}) //unable to create token
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user, "token": fmt.Sprintf("Bearer %s", token)})
	return
}
