package controller

import (
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseDetailController struct {
	uc             usecase.CourseDetailUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *CourseDetailController) CreateHandler(ctx *gin.Context) {
	var payload dto.CourseDetailRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdCourseDetail, err := c.uc.AddCourse(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Course Detail successfully created", createdCourseDetail)
}

func (c *CourseDetailController) GetHandlerByID(ctx *gin.Context) {
	courseDetailID := ctx.Param("id")

	courseDetail, err := c.uc.FindCourseDetailByID(courseDetailID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Course Detail by ID", courseDetail)
}

func (c *CourseDetailController) UpdateHandler(ctx *gin.Context) {
	var payload dto.CourseDetailRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	courseDetailID := ctx.Param("id")
	updatedCourseDetail, err := c.uc.UpdateAttendance(payload, courseDetailID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Course Detail Updated", updatedCourseDetail)
}

func (c *CourseDetailController) DeleteHandler(ctx *gin.Context) {
	courseDetailID := ctx.Param("id")

	deletedCourseDetail, err := c.uc.Delete(courseDetailID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Course Detail Deleted", deletedCourseDetail)
}

func (c *CourseDetailController) Route() {
	c.rg.POST("/course-detail", c.authMiddleware.RequireToken("admin"), c.CreateHandler)
	c.rg.GET("/course-detail/:id", c.authMiddleware.RequireToken("admin", "trainer", "admin"), c.GetHandlerByID)
	c.rg.PUT("/course-detail/:id", c.authMiddleware.RequireToken("admin"), c.UpdateHandler)
	c.rg.DELETE("/course-detail/:id", c.authMiddleware.RequireToken("admin"), c.DeleteHandler)
}

func NewCourseDetailController(uc usecase.CourseDetailUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *CourseDetailController {
	return &CourseDetailController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
