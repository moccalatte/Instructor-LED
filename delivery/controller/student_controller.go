package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	uc usecase.StudentUseCase
	rg *gin.RouterGroup
}

func (s *StudentController) CreateHandler(ctx *gin.Context) {
	var payload dto.StudentRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	createdStudent, err := s.uc.AddStudent(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Student successfuly Created", createdStudent)
}

func (s *StudentController) GetHandlerID(ctx *gin.Context) {
	studentId := ctx.Param("id")
	student, err := s.uc.FindStudentByID(studentId)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Student successfuly by ID", student)
}

func (s *StudentController) UpdateHandler(ctx *gin.Context) {
	var payload dto.StudentRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	studentId := ctx.Param("id")

	student, err := s.uc.UpdateStudent(payload, studentId)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Student successfuly Updated", student)
}

func (s *StudentController) DeleteHandler(ctx *gin.Context) {
	studentId := ctx.Param("id")

	deletedStudent, err := s.uc.DeleteStudent(studentId)

	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Customer Deleted", deletedStudent)
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