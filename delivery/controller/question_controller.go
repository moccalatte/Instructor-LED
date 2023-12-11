package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"final-project-kelompok-1/utils/common"
	"final-project-kelompok-1/delivery/middleware"
	"net/http"
    "encoding/base64"
	"bytes"
	"io"
    "fmt"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	uc usecase.QuestionUseCase
	rg *gin.RouterGroup
	mw  middleware.AuthMiddleware
	jwt common.JwtToken
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

func (c *QuestionController) AddQuestion(ctx *gin.Context) {
	var payload dto.QuestionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle upload gambar di sini
	imageData, err := extractImageData(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to extract image data"})
		return
	}

	// Sertakan data gambar ke payload
	payload.ImagePath = imageData

	// Panggil use case untuk menambah pertanyaan
	question, err := c.uc.AddQuestion(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to add question: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, question)
}

func extractImageData(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("image")
	if err != nil {
		return "", err
	}

	// Buka file gambar
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Buat buffer untuk menampung konten file
	var buf bytes.Buffer
	_, err = io.Copy(&buf, src)
	if err != nil {
		return "", err
	}

	// Encode file gambar ke base64
	imageData := base64.StdEncoding.EncodeToString(buf.Bytes())

	return imageData, nil
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
	q.rg.POST("/question", q.CreateHandler)
	q.rg.GET("/question/:id", q.GetHandlerByID)
	q.rg.PUT("/question/:id", q.UpdateHandler)
	q.rg.DELETE("/question/:id", q.DeleteHandler)
	q.rg.POST("/question/:id/answer", q.AnswerHandler)
	q.rg.POST("/questions", q.mw.RequireToken(), q.AddQuestion)
}

func NewQuestionController(uc usecase.QuestionUseCase, rg *gin.RouterGroup, mw middleware.AuthMiddleware, jwt common.JwtToken) *QuestionController {
	return &QuestionController{
		uc:  uc,
		rg:  rg,
		mw:  mw,
		jwt: jwt,
	}
}