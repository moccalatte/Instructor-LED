package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type CourseUseCase interface {
	AddCourse(payload dto.CourseRequestDto) (model.Course, error)
	FindCourseByID(id string) (model.Course, error)
	UpdateCourse(payload dto.CourseRequestDto, id string) (model.Course, error)
	DeleteCourse(id string) (model.Course, error)
	GetAllCourse() ([]model.Course, error)
}

type courseUseCase struct {
	repo repository.CourseRepository
}

func (c *courseUseCase) AddCourse(payload dto.CourseRequestDto) (model.Course, error) {
	newCourse := model.Course{
		CourseName: payload.CourseName,
	}

	addedCourse, err := c.repo.Create(newCourse)

	if err != nil {
		return model.Course{}, fmt.Errorf("failed to add Course : %s", err.Error())
	}

	return addedCourse, nil

}

func (c *courseUseCase) FindCourseByID(id string) (model.Course, error) {
	courseWithId, err := c.repo.GetById(id)

	if err != nil {
		return model.Course{}, fmt.Errorf("failed to find Course : %s", err.Error())
	}

	return courseWithId, nil
}

func (c *courseUseCase) GetAllCourse() ([]model.Course, error) {
	courseAll, err := c.repo.GetAll()

	if err != nil {
		fmt.Println("Error Get All Data in use case : ", err.Error())
		return courseAll, fmt.Errorf("failed to find data : %v", err.Error())
	}

	return courseAll, nil
}

func (c *courseUseCase) UpdateCourse(payload dto.CourseRequestDto, id string) (model.Course, error) {
	newCourse := model.Course{
		CourseName: payload.CourseName,
		// CourseDetailID: payload.CourseDetailID,
	}

	updatedCourse, err := c.repo.Update(newCourse, id)

	if err != nil {
		return model.Course{}, fmt.Errorf("failed to Update Course : %s", err.Error())
	}

	return updatedCourse, nil
}

func (c *courseUseCase) DeleteCourse(id string) (model.Course, error) {
	deletedCourse, err := c.repo.Delete(id)

	if err != nil {
		return model.Course{}, fmt.Errorf("failed to Delete Course : %s", err.Error())
	}

	return deletedCourse, nil
}

func NewCourseUseCase(repo repository.CourseRepository) CourseUseCase {
	return &courseUseCase{repo: repo}
}