package repository

import (
	"database/sql"
	"fmt"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type CourseRepository interface {
	Create(payload model.Course) (model.Course, error)
	GetById(id string) (model.Course, error)
	Update(payload model.Course, id string) (model.Course, error)
	Delete(id string) (model.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func (c *courseRepository) Create(payload model.Course) (model.Course, error) {
	fmt.Print(payload.CourseName, "DATAMASUK")
	tx, err := c.db.Begin()
	if err != nil {
		return model.Course{}, err
	}
	fmt.Print(payload.CourseName, "DATAMASUK")

	var course model.Course

	err = tx.QueryRow(common.CreateCourse,
		payload.CourseName,
		payload.Description,
		time.Now(),
		false,
	).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.IsDeleted,
	)
	fmt.Print(course.CourseName, "DATANYAMANA")
	// fmt.Print(course.CourseDetailID)
	fmt.Print(err)

	if err != nil {
		return model.Course{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *courseRepository) GetById(id string) (model.Course, error) {
	var course model.Course
	err := c.db.QueryRow(common.GetCourseById, id).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.IsDeleted,
	)
	if err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *courseRepository) Update(payload model.Course, id string) (model.Course, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.Course{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var course model.Course
	err = tx.QueryRow(common.UpdateCourseById,
		payload.CourseName,
		payload.Description,
		time.Now(),
		true,
		id).Scan(
			&course.CourseID,
			&course.CourseName,
			&course.Description,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsDeleted,
	)
	if err != nil {
		return model.Course{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *courseRepository) Delete(id string) (model.Course, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.Course{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var course model.Course
	err = tx.QueryRow(common.UpdateCourseById,
		true,
		id).Scan(
			&course.CourseID,
			&course.CourseName,
			&course.Description,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsDeleted,
	)
	if err != nil {
		return model.Course{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}
