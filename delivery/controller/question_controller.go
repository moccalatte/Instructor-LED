package controller

import (
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	uc             usecase.QuestionUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (q *QuestionController) CreateHandler(ctx *gin.Context) {
	var payload dto.QuestionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdQuestion, err := q.uc.AddQuestion(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Question successfully created", createdQuestion)
}

func (q *QuestionController) GetHandlerByID(ctx *gin.Context) {
	questionID := ctx.Param("id")

	question, err := q.uc.FindQuestionById(questionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Question by ID", question)
}

func (q *QuestionController) GetHandlerAll(ctx *gin.Context) {

	question, err := q.uc.GetAllQuestion()
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Question All", question)
}

func (q *QuestionController) UpdateHandler(ctx *gin.Context) {
	var payload dto.QuestionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	questionID := ctx.Param("id")
	updatedQuestion, err := q.uc.Update(payload, questionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Question Updated", updatedQuestion)
}

func (q *QuestionController) DeleteHandler(ctx *gin.Context) {
	questionID := ctx.Param("id")

	deletedQuestion, err := q.uc.Delete(questionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Question Deleted", deletedQuestion)
}

func (q *QuestionController) AnswerHandler(ctx *gin.Context) {
	var payload dto.QuestionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	questionID := ctx.Param("id")
	answeredQuestion, err := q.uc.Answer(payload, questionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Question Answered", answeredQuestion)
}

func (q *QuestionController) Route() {
	q.rg.POST("/question", q.authMiddleware.RequireToken("student"), q.CreateHandler)
	q.rg.GET("/question/:id", q.authMiddleware.RequireToken("student", "trainer"), q.GetHandlerByID)
	q.rg.GET("/question", q.GetHandlerAll)
	q.rg.PUT("/question/:id", q.authMiddleware.RequireToken("student"), q.UpdateHandler)
	q.rg.DELETE("/question/:id", q.authMiddleware.RequireToken("student"), q.DeleteHandler)
	q.rg.PUT("/question-answer/:id", q.authMiddleware.RequireToken("trainer"), q.AnswerHandler)
}

func NewQuestionController(uc usecase.QuestionUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *QuestionController {
	return &QuestionController{uc: uc, rg: rg, authMiddleware: authMiddleware}

}
