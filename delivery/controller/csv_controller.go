package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CsvController struct {
	uc usecase.CsvUseCase
	rg gin.RouterGroup
}

func (c *CsvController) StartPointCsvHandler(ctx *gin.Context) {
	allIdSession, err := c.uc.WriteCsv()
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Success To make Report", allIdSession)
}

func (c *CsvController) Route() {
	c.rg.GET("/report", c.StartPointCsvHandler)
}

func NewCsvController(uc usecase.CsvUseCase, rg *gin.RouterGroup) *CsvController {
	return &CsvController{uc: uc, rg: *rg}
}
