package controller

import (
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc             usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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

func (u *UserController) GetHandlerAll(ctx *gin.Context) {
	user, err := u.uc.GetAllUser()
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get all user", user)
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

// func (u *UserController) Route() {
// 	ur := u.rg.Group("/user", u.authMiddleware.RequireToken("admin", "trainer"))
// 	ur.POST("/", u.CreateHandler)
// 	ur.GET("/user/:id", u.GetHandlerByID)
// 	ur.PUT("/user/:id", u.UpdateHandler)
// 	ur.DELETE("/user/:id", u.DeleteHandler)
// }

func (u *UserController) Route() {
	//u.rg.POST("/user", u.authMiddleware.RequireToken("admin"), u.CreateHandler)
	u.rg.POST("/user", u.CreateHandler)
	u.rg.GET("/user/:id", u.authMiddleware.RequireToken("admin"), u.GetHandlerByID)
	u.rg.GET("/user", u.authMiddleware.RequireToken("admin"), u.GetHandlerAll)
	u.rg.PUT("/user/:id", u.authMiddleware.RequireToken("admin"), u.UpdateHandler)
	u.rg.DELETE("/user/:id", u.authMiddleware.RequireToken("admin"), u.DeleteHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *UserController {
	return &UserController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
