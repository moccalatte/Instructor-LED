package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type CourseDetailUseCase interface {
	AddCourse(payload dto.CourseDetailRequestDto) (model.CourseDetail, error)
	FindCourseDetailByID(id string) (model.CourseDetail, error)
	UpdateAttendance(payload dto.CourseDetailRequestDto, id string) (model.CourseDetail, error)
	Delete(id string) (model.CourseDetail, error)
}

type courseDetailUseCase struct {
	repo repository.CourseDetailRepository
}

func (c *courseDetailUseCase) AddCourse(payload dto.CourseDetailRequestDto) (model.CourseDetail, error) {
	newCourse := model.CourseDetail{
		CourseID:      payload.CourseId,
		CourseChapter: payload.CourseChapter,
		CourseContent: payload.CourseContent,
	}

	AddCourse, err := c.repo.Create(newCourse)

	if err != nil {
		return model.CourseDetail{}, fmt.Errorf("failed to add Course : %s", err.Error())
	}

	return AddCourse, nil

}

func (c *courseDetailUseCase) FindCourseDetailByID(id string) (model.CourseDetail, error) {
	courseDetail, err := c.repo.GetById(id)

	if err != nil {
		return model.CourseDetail{}, fmt.Errorf("failed to find Course : %s", err.Error())
	}

	return courseDetail, nil
}

func (c *courseDetailUseCase) UpdateAttendance(payload dto.CourseDetailRequestDto, id string) (model.CourseDetail, error) {
	newCourse := model.CourseDetail{
		CourseID:      payload.CourseId,
		CourseChapter: payload.CourseChapter,
		CourseContent: payload.CourseContent,
	}

	UpdateCourse, err := c.repo.Update(newCourse, id)

	if err != nil {
		return model.CourseDetail{}, fmt.Errorf("failed to Update Course Detail : %s", err.Error())
	}

	return UpdateCourse, nil
}

func (c *courseDetailUseCase) Delete(id string) (model.CourseDetail, error) {
	deletedCoursDetail, err := c.repo.Delete(id)

	if err != nil {
		return model.CourseDetail{}, fmt.Errorf("failed to Delete Course : %s", err.Error())
	}

	return deletedCoursDetail, nil
}

func NewCourseDetailUseCase(repo repository.CourseDetailRepository) CourseDetailUseCase {
	return &courseDetailUseCase{repo: repo}
}
