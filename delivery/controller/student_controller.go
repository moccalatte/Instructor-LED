package controller

import (
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	uc             usecase.StudentUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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

func (s *StudentController) GetHandlerAll(ctx *gin.Context) {
	student, err := s.uc.GetAllStudent()
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get all student", student)
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
	s.rg.POST("/student", s.authMiddleware.RequireToken("admin"), s.CreateHandler)
	s.rg.GET("/student/:id", s.authMiddleware.RequireToken("admin", "trainer"), s.GetHandlerID)
	s.rg.GET("/student", s.authMiddleware.RequireToken("admin", "trainer"), s.GetHandlerAll)
	s.rg.PUT("/student/:id", s.authMiddleware.RequireToken("admin"), s.UpdateHandler)
	s.rg.DELETE("/student/:id", s.authMiddleware.RequireToken("admin"), s.DeleteHandler)
}

func NewStudentController(uc usecase.StudentUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *StudentController {
	return &StudentController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
