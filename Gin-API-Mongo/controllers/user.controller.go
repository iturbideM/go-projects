package controllers

import (
	"main/models"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	return ctx.Json(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	return ctx.Json(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	return ctx.Json(200, "")
}

func (uc *UserController) RegisterUserRouters(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser())
	userRoute.GET("/get/:name", uc.GetUser())
	userRoute.GET("/getall", uc.GetAll())
	userRoute.PATCH("/update", uc.UpdateUser())
	userRoute.DELETE("/delete", uc.DeleteUser())
}
