package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type AttendanceUseCase interface {
	AddAttendance(payload dto.AttendanceRequestDto) (model.Attendance, error)
	FindAttendanceByID(id string) (model.Attendance, error)
	GetAllAttendance() ([]model.Attendance, error)
	FindAttendanceBySessionId(id string) (model.Attendance, error)
	UpdateAttendance(payload dto.AttendanceRequestDto, id string) (model.Attendance, error)
	Delete(id string) (model.Attendance, error)
}

type attendanceUseCase struct {
	repo repository.AttendanceRepository
}

func (c *attendanceUseCase) AddAttendance(payload dto.AttendanceRequestDto) (model.Attendance, error) {
	newAttendance := model.Attendance{
		SessionID:         payload.SessionID,
		StudentID:         payload.StudentID,
		AttendanceStudent: payload.AttendanceStudent,
	}

	AddAttendance, err := c.repo.Create(newAttendance)

	if err != nil {
		return model.Attendance{}, fmt.Errorf("failed to add Course : %s", err.Error())
	}

	return AddAttendance, nil

}

func (c *attendanceUseCase) FindAttendanceByID(id string) (model.Attendance, error) {
	attendance, err := c.repo.GetById(id)

	if err != nil {
		return model.Attendance{}, fmt.Errorf("failed to find Attendance : %s", err.Error())
	}

	return attendance, nil
}

// func (c *attendanceUseCase) GetAllAttendance() ([]model.Attendance, error) {
// 	var attendanceSlice []model.Attendance
// 	attendance, err := c.repo.FindAll()

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get all data: %s", err.Error())
// 	}

// 	return append(attendanceSlice, attendance...), nil
// }

func (c *attendanceUseCase) FindAttendanceBySessionId(id string) (model.Attendance, error) {
	// fmt.Print(id, "USECASEATTSESSION")
	attendance, err := c.repo.GetBySessionId(id)

	if err != nil {
		return model.Attendance{}, fmt.Errorf("failed to find Attendance : %s", err.Error())
	}

	return attendance, nil
}

func (c *attendanceUseCase) GetAllAttendance() ([]model.Attendance, error) {
	attendanceAll, err := c.repo.GetAll()

	if err != nil {
		fmt.Println("Error Get All Data in use case : ", err.Error())
		return attendanceAll, fmt.Errorf("failed to find data : %v", err.Error())
	}

	return attendanceAll, nil
}

func (c *attendanceUseCase) UpdateAttendance(payload dto.AttendanceRequestDto, id string) (model.Attendance, error) {
	newAttendance := model.Attendance{
		SessionID:         payload.SessionID,
		StudentID:         payload.StudentID,
		AttendanceStudent: payload.AttendanceStudent,
	}

	UpdateAttendance, err := c.repo.Update(newAttendance, id)

	if err != nil {
		return model.Attendance{}, fmt.Errorf("failed to Update Attendance : %s", err.Error())
	}

	return UpdateAttendance, nil
}

func (c *attendanceUseCase) Delete(id string) (model.Attendance, error) {
	deletedAtendance, err := c.repo.Delete(id)

	if err != nil {
		return model.Attendance{}, fmt.Errorf("failed to Delete Attendance : %s", err.Error())
	}

	return deletedAtendance, nil
}

func NewAttendanceUseCase(repo repository.AttendanceRepository) AttendanceUseCase {
	return &attendanceUseCase{repo: repo}
}
