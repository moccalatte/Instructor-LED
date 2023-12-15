package controller

import (
	// "final-project-kelompok-1/model"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"final-project-kelompok-1/utils/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc         usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService common.JwtToken
}

func (a *AuthController) registerHandler(ctx *gin.Context) {
	var payload model.Users
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newResponse, err := a.uc.Register(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Created", newResponse)
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resPayload, err := a.uc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", resPayload)
}

func (a *AuthController) loginStudentHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resPayload, err := a.uc.LoginStudent(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", resPayload)
}

func (a *AuthController) refreshTokenHandler(ctx *gin.Context) {
	tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newToken, err := a.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	common.SendCreateResponse(ctx, "Ok", newToken)
}

func (a *AuthController) Route() {
	ug := a.rg.Group("/auth")
	ug.POST("/register", a.registerHandler)
	ug.POST("/login", a.loginHandler)
	ug.POST("/login-student", a.loginStudentHandler)
	ug.GET("/refresh-token", a.refreshTokenHandler)
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwtService: jwtService}
}
