package repository

import (
	"database/sql"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type CourseRepository interface {
	Create(payload model.Course) (model.Course, error)
	GetById(id int) (model.Course, error)
	Update(payload model.Course, id int) (model.Course, error)
	Delete(id int) (model.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func (c *courseRepository) Create(payload model.Course) (model.Course, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.Course{}, err
	}

	var course model.Course
	err = tx.QueryRow(common.CreateCourse,
		payload.CourseName,
		payload.CourseDetailID,
		true,
	).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CourseDetailID,
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

func (c *courseRepository) GetById(id int) (model.Course, error) {
	var course model.Course
	err := c.db.QueryRow(common.GetCourseById, id).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CourseDetailID,
		&course.IsDeleted,
	)
	if err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *courseRepository) Update(payload model.Course, id int) (model.Course, error) {
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
		payload.CourseDetailID,
		true,
		id).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CourseDetailID,
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

func (c *courseRepository) Delete(id int) (model.Course, error) {
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
		false,
		id).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CourseDetailID,
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

func NewRoleRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}
