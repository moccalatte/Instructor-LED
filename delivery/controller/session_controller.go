package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	uc usecase.SessionUseCase
	rg *gin.RouterGroup
}

func (s *SessionController) CreateHandler(ctx *gin.Context) {
	var payload dto.SessionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdSession, err := s.uc.AddSession(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	dto.SendSingleResponse(ctx, http.StatusCreated, "Session successfuly Created", createdSession)
}

func (s *SessionController) Route() {
	s.rg.POST("/session", s.CreateHandler)
}

func NewSessionController(uc usecase.SessionUseCase, rg *gin.RouterGroup) *SessionController {
	return &SessionController{uc: uc, rg: rg}
}
