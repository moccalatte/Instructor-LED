package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *UserController) CreateHandler(ctx *gin.Context) {
	var payload dto.UserRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdUser, err := u.uc.AddUser(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "user successfuly Created", createdUser)

}

func (u *UserController) GetHandlerByID(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := u.uc.FindUserByID(userId)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get user by ID", user)
}

func (u *UserController) UpdateHandler(ctx *gin.Context) {
	var payload dto.UserRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userId := ctx.Param("id")

	user, err := u.uc.UpdateUser(payload, userId)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "user successfuly Updated", user)
}

func (u *UserController) DeleteHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	deletedUser, err := u.uc.DeleteUser(userId)

	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "User Deleted", deletedUser)
}

func (u *UserController) Route() {
	u.rg.POST("/user", u.CreateHandler)
	u.rg.GET("/user/:id", u.GetHandlerByID)
	u.rg.PUT("/user/:id", u.UpdateHandler)
	u.rg.DELETE("/user/:id", u.DeleteHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{uc: uc, rg: rg}
}
