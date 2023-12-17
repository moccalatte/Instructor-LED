package controller

import (
	"errors"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/mock/controller_mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"strings"
	"testing"
)

func TestCreateAttendanceHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := controller_mock.NewAttendanceControllerMock()

	// Set the expected behavior for the CreateHandler method
	expectedResult := dto.AttendanceRequestDto{
		// Set your expected result here
	}
	expectedError := errors.New("expected error")

	mockController.On("CreateHandler", mock.Anything).Return(expectedResult, expectedError)

	// Create a router and register the CreateHandler route
	router := gin.New()
	mockController.Route(router.Group("/mock"))

	// Prepare a request
	reqBody := `{"your": "request", "body": "here"}`
	req, _ := http.NewRequest("POST", "/mock/attendance", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// Assert other expectations as needed

	// Assert that the expected methods were called with the correct arguments
	mockController.AssertExpectations(t)
}

func TestGetAttendanceByIDHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := controller_mock.NewAttendanceControllerMock()

	// Set the expected behavior for the GetHandlerByID method
	expectedResult := dto.AttendanceRequestDto{
		// Set your expected result here
	}
	expectedError := errors.New("expected error")

	mockController.On("GetHandlerByID", mock.Anything).Return(expectedResult, expectedError)

	// Create a router and register the GetHandlerByID route
	router := gin.New()
	mockController.Route(router.Group("/mock"))

	// Prepare a request
	attendanceID := "your-attendance-id"
	req, _ := http.NewRequest("GET", "/mock/attendance/"+attendanceID, nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// Add more assertions based on your expected behavior

	// Assert that the expected methods were called with the correct arguments
	mockController.AssertExpectations(t)
}

func TestGetAllAttendanceHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := controller_mock.NewAttendanceControllerMock()

	// Set the expected behavior for the GetHandlerAll method
	expectedResult := []dto.AttendanceRequestDto{
		// Set your expected result here
	}
	expectedError := errors.New("expected error")

	mockController.On("GetHandlerAll").Return(expectedResult, expectedError)

	// Create a router and register the GetHandlerAll route
	router := gin.New()
	mockController.Route(router.Group("/mock"))

	// Prepare a request
	req, _ := http.NewRequest("GET", "/mock/attendance", nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// Add more assertions based on your expected behavior

	// Assert that the expected methods were called with the correct arguments
	mockController.AssertExpectations(t)
}

func TestUpdateAttendanceHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := controller_mock.NewAttendanceControllerMock()

	// Set the expected behavior for the UpdateHandler method
	expectedResult := dto.AttendanceRequestDto{
		// Set your expected result here
	}
	expectedError := errors.New("expected error")

	mockController.On("UpdateHandler", mock.Anything).Return(expectedResult, expectedError)

	// Create a router and register the UpdateHandler route
	router := gin.New()
	mockController.Route(router.Group("/mock"))

	// Prepare a request
	attendanceID := "your-attendance-id"
	requestBody := dto.AttendanceRequestDto{
		// Set your request body here
	}
	reqBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PUT", "/mock/attendance/"+attendanceID, bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// Add more assertions based on your expected behavior

	// Assert that the expected methods were called with the correct arguments
	mockController.AssertExpectations(t)
}

func TestDeleteAttendanceHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := controller_mock.NewAttendanceControllerMock()

	// Set the expected behavior for the DeleteHandler method
	expectedResult := dto.AttendanceRequestDto{
		// Set your expected result here
	}
	expectedError := errors.New("expected error")

	mockController.On("DeleteHandler", mock.Anything).Return(expectedResult, expectedError)

	// Create a router and register the DeleteHandler route
	router := gin.New()
	mockController.Route(router.Group("/mock"))

	// Prepare a request
	attendanceID := "your-attendance-id"
	req, _ := http.NewRequest("DELETE", "/mock/attendance/"+attendanceID, nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// Add more assertions based on your expected behavior

	// Assert that the expected methods were called with the correct arguments
	mockController.AssertExpectations(t)
}