package controller

import (
	"final-project-kelompok-1/config"
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	uc             usecase.QuestionUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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

// Modifikasi fungsi extractImageData
func extractImageData(ctx *gin.Context) (string, error) {
    file, err := ctx.FormFile("image")
    if err != nil {
        return "", err
    }

    // Buat direktori jika belum ada
    if err := os.MkdirAll(config.ImageUploadDirectory, 0755); err != nil {
        return "", err
    }

    // Buat path file unik untuk gambar
    imagePath := filepath.Join(config.ImageUploadDirectory, generateUniqueFileName(file.Filename))

    // Buka file gambar
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // Buat file baru untuk menyimpan gambar
    dst, err := os.Create(imagePath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    // Salin konten file gambar
    _, err = io.Copy(dst, src)
    if err != nil {
        return "", err
    }

    // Mengembalikan path file yang baru dibuat
    return imagePath, nil
}

// Fungsi untuk membuat nama file yang unik
func generateUniqueFileName(originalName string) string {
    baseName := strings.TrimSuffix(originalName, filepath.Ext(originalName))
    timestamp := time.Now().UnixNano()
    return fmt.Sprintf("%s_%d%s", baseName, timestamp, filepath.Ext(originalName))
}

func (q *QuestionController) UploadImageHandler(ctx *gin.Context) {
    // Handle upload gambar di sini
    imagePath, err := extractImageData(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to extract image data"})
        return
    }

    // Membuat URL gambar berdasarkan path gambar
    imageURL := generateImageURL(imagePath)

    ctx.JSON(http.StatusOK, gin.H{"message": "Image successfully uploaded", "imageURL": imageURL})
}

// Fungsi untuk membuat URL gambar berdasarkan path gambar
func generateImageURL(imagePath string) string {
    // Mendapatkan nama file dari path gambar
    fileName := filepath.Base(imagePath)

    // Melakukan encoding pada nama file untuk mengatasi spasi
    encodedFileName := url.PathEscape(fileName)

    // Bentuk URL gambar berdasarkan nama file yang telah diencode
    return config.BaseURL + "/uploads/" + encodedFileName
}

func (q *QuestionController) CreateHandler(ctx *gin.Context) {
	

	var payload dto.QuestionRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdQuestion, err := q.uc.AddQuestion(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Question successfully created", "createdQuestion": createdQuestion})
}

func (q *QuestionController) DownloadImageHandler(ctx *gin.Context) {
	questionID := ctx.Param("id")

	// Panggil use case untuk mendapatkan path gambar
	imagePath, err := q.uc.GetImagePath(questionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get image path"})
		return
	}

	// Baca file gambar
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read image file"})
		return
	}

	// Set header untuk respons
	ctx.Header("Content-Type", "image/jpeg")
	ctx.Header("Content-Disposition", "attachment; filename=image.jpg")

	// Kirim file gambar sebagai respons
	ctx.Data(http.StatusOK, "image/jpeg", imageBytes)
}



func (q *QuestionController) Route() {
	q.rg.POST("/question", q.authMiddleware.RequireToken("student"), q.CreateHandler)
	q.rg.GET("/question/:id", q.authMiddleware.RequireToken("student", "trainer"), q.GetHandlerByID)
	q.rg.GET("/question", q.authMiddleware.RequireToken("student", "trainer"), q.GetHandlerAll)
	q.rg.PUT("/question/:id", q.authMiddleware.RequireToken("student"), q.UpdateHandler)
	q.rg.DELETE("/question/:id", q.authMiddleware.RequireToken("student"), q.DeleteHandler)
	q.rg.PUT("/question-answer/:id", q.authMiddleware.RequireToken("trainer"), q.AnswerHandler)
	q.rg.POST("/question/upload", q.authMiddleware.RequireToken("student"), q.UploadImageHandler)
	q.rg.GET("/question/:id/download", q.authMiddleware.RequireToken("student", "trainer"), q.DownloadImageHandler)
}

func NewQuestionController(uc usecase.QuestionUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *QuestionController {
	return &QuestionController{uc: uc, rg: rg, authMiddleware: authMiddleware}

}