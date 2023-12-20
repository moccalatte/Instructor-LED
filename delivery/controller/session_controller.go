package controller

import (
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	uc             usecase.SessionUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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

	dto.SendSingleResponse(ctx, http.StatusCreated, "Session successfully created", createdSession)
}

func (s *SessionController) GetHandlerByID(ctx *gin.Context) {
	sessionID := ctx.Param("id")

	session, err := s.uc.FindSessionById(sessionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Session by ID", session)
}

func (s *SessionController) GetHandlerAll(ctx *gin.Context) {
	session, err := s.uc.GetAllSession()
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get all session", session)
}

func (s *SessionController) UpdateHandler(ctx *gin.Context) {
	var payload dto.SessionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	sessionID := ctx.Param("id")
	updatedSession, err := s.uc.Update(payload, sessionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Session Updated", updatedSession)
}

func (s *SessionController) UpdateNoteHandler(ctx *gin.Context) {
	var payload dto.SessionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	sessionID := ctx.Param("id")
	updatedSession, err := s.uc.UpdateNote(payload, sessionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Session Updated", updatedSession)
}

func (s *SessionController) DeleteHandler(ctx *gin.Context) {
	sessionID := ctx.Param("id")

	deletedSession, err := s.uc.Delete(sessionID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Session Deleted", deletedSession)
}

func (s *SessionController) Route() {
	s.rg.POST("/session", s.authMiddleware.RequireToken("admin"), s.CreateHandler)
	s.rg.GET("/session/:id", s.authMiddleware.RequireToken("admin", "trainer", "student"), s.GetHandlerByID)
	s.rg.GET("/session", s.authMiddleware.RequireToken("admin", "trainer", "student"), s.GetHandlerAll)
	s.rg.PUT("/session/:id", s.authMiddleware.RequireToken("admin"), s.UpdateHandler)
	s.rg.PUT("/session/note/:id", s.authMiddleware.RequireToken("trainer"), s.UpdateNoteHandler)
	s.rg.DELETE("/session/:id", s.authMiddleware.RequireToken("admin"), s.DeleteHandler)
}

func NewSessionController(uc usecase.SessionUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *SessionController {
	return &SessionController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
