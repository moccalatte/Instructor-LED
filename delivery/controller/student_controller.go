package controller

import (
	"final-project-kelompok-1/usecase"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	uc usecase.StudentUseCase
	rg *gin.RouterGroup
}

func (s *StudentController) CreateHandler(ctx *gin.Context) {

}

func (s *StudentController) GetHandlerID(ctx *gin.Context) {

}

func (s *StudentController) UpdateHandler(ctx *gin.Context) {

}

func (s *StudentController) DeleteHandler(ctx *gin.Context) {

}

func (s *StudentController) Route() {
	s.rg.POST("/student", s.CreateHandler)
	s.rg.GET("/student/:id", s.GetHandlerID)
	s.rg.PUT("/student/:id", s.UpdateHandler)
	s.rg.DELETE("/student/:id", s.DeleteHandler)
}

func NewStudentController(uc usecase.StudentUseCase, rg *gin.RouterGroup) *StudentController {
	return &StudentController{uc: uc, rg: rg}
}
