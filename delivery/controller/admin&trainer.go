package controller

import (
	"final-project-kelompok-1/usecase"

	"github.com/gin-gonic/gin"
)

type AdminTrainerController struct {
	uc usecase.AdminTrainerUseCase
	rg *gin.RouterGroup
}

func (a *AdminTrainerController) CreateHandler(ctx *gin.Context) {

}

func (a *AdminTrainerController) GetHandlerByID(ctx *gin.Context) {

}

func (a *AdminTrainerController) UpdateHandler(ctx *gin.Context) {

}

func (a *AdminTrainerController) DeleteHandler(ctx *gin.Context) {

}

func (a *AdminTrainerController) Route() {
	a.rg.POST("/----", a.CreateHandler)
	a.rg.GET("/-----/:id", a.GetHandlerByID)
	a.rg.PUT("/-----/:id", a.UpdateHandler)
	a.rg.DELETE("/----/:id", a.DeleteHandler)
}

func NewAdminTrainerController(uc usecase.AdminTrainerUseCase, rg *gin.RouterGroup) *AdminTrainerController {
	return &AdminTrainerController{uc: uc, rg: rg}
}
