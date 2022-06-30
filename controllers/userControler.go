package controllers

import (
	"crud_gin_gonic/models"
	"crud_gin_gonic/services"
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

func (uc *UserController) CreateUser(ctx *gin.Context){

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err:= uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}

func (uc *UserController) GetUser(ctx *gin.Context){
	username := ctx.Param("name")

	user, err := uc.UserService.GetUser(&username)

	if err!=nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context){
	users, err := uc.UserService.GetAll()
	if err!= nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK,users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context){
	var user models.User
	if err:= ctx.ShouldBindJSON(&user); err !=nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err!=nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"message": "success"})

}
func (uc *UserController) DeleteUser(ctx *gin.Context){
	
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Thanks success"})
} 

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userrout := rg.Group("/user")
	userrout.POST("/create",uc.CreateUser)
	userrout.GET("/get/:name",uc.GetUser)
	userrout.GET("/getall", uc.GetAll)
	userrout.PATCH("/update",uc.UpdateUser)
	userrout.DELETE("/delete/:name",uc.DeleteUser)
}