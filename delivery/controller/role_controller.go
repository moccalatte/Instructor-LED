package controller

import (
	"final-project-kelompok-1/usecase"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	uc usecase.RoleUseCase
	rg *gin.RouterGroup
}

func (r *RoleController) CreateHandler(ctx *gin.Context) {

}

func (r *RoleController) GetHandlerByID(ctx *gin.Context) {

}

func (r *RoleController) UpdateHandler(ctx *gin.Context) {

}

func (r *RoleController) DeleteHandler(ctx *gin.Context) {

}

func (r *RoleController) Route() {
	r.rg.POST("/----", r.CreateHandler)
	r.rg.GET("/-----/:id", r.GetHandlerByID)
	r.rg.PUT("/-----/:id", r.UpdateHandler)
	r.rg.DELETE("/----/:id", r.DeleteHandler)
}

func NewRoleController(uc usecase.RoleUseCase, rg *gin.RouterGroup) *RoleController {
	return &RoleController{uc: uc, rg: rg}
}
