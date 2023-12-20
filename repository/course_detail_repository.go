package repository

import (
	"database/sql"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type CourseDetailRepository interface {
	Create(payload model.CourseDetail) (model.CourseDetail, error)
	GetById(id string) (model.CourseDetail, error)
	Update(payload model.CourseDetail, id string) (model.CourseDetail, error)
	Delete(id string) (model.CourseDetail, error)
	GetAll() ([]model.CourseDetail, error)
}

type courseDetailRepository struct {
	db *sql.DB
}

func (c *courseDetailRepository) Create(payload model.CourseDetail) (model.CourseDetail, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.CourseDetail{}, err
	}
	var course_detail model.CourseDetail

	err = tx.QueryRow(common.CreatCourseDetail,
		payload.CourseID,
		payload.CourseChapter,
		payload.CourseContent,
		time.Now(),
		false).Scan(
		&course_detail.CourseDetailID,
		&course_detail.CourseID,
		&course_detail.CourseChapter,
		&course_detail.CourseContent,
		&course_detail.CreatedAt,
		&course_detail.UpdatedAt,
		&course_detail.IsDeleted,
	)

	if err != nil {
		return model.CourseDetail{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.CourseDetail{}, err
	}

	return course_detail, nil
}

func (c *courseDetailRepository) GetById(id string) (model.CourseDetail, error) {
	var course_detail model.CourseDetail
	err := c.db.QueryRow(common.GetCourseDetailById, id).Scan(
		&course_detail.CourseDetailID,
		&course_detail.CourseID,
		&course_detail.CourseChapter,
		&course_detail.CourseContent,
		&course_detail.CreatedAt,
		&course_detail.UpdatedAt,
		&course_detail.IsDeleted,
	)
	if err != nil {
		return model.CourseDetail{}, err
	}
	return course_detail, nil
}

func (c *courseDetailRepository) Update(payload model.CourseDetail, id string) (model.CourseDetail, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.CourseDetail{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var course_detail model.CourseDetail
	err = tx.QueryRow(common.UpdateCourseDetailByid,
		payload.CourseID,
		payload.CourseChapter,
		payload.CourseContent,
		time.Now(),
		false,
		id,
	).Scan(
		&course_detail.CourseDetailID,
		&course_detail.CourseID,
		&course_detail.CourseChapter,
		&course_detail.CourseContent,
		&course_detail.CreatedAt,
		&course_detail.UpdatedAt,
		&course_detail.IsDeleted,
	)
	if err != nil {
		return model.CourseDetail{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.CourseDetail{}, err
	}

	return course_detail, nil

}
func (c *courseDetailRepository) Delete(id string) (model.CourseDetail, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.CourseDetail{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var course_detail model.CourseDetail
	err = tx.QueryRow(common.DeleteCourseDetailById,
		true,
		id,
	).Scan(
		&course_detail.CourseDetailID,
		&course_detail.CourseID,
		&course_detail.CourseChapter,
		&course_detail.CourseContent,
		&course_detail.CreatedAt,
		&course_detail.UpdatedAt,
		&course_detail.IsDeleted,
	)
	if err != nil {
		return model.CourseDetail{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.CourseDetail{}, err
	}

	return course_detail, nil
}

func (c *courseDetailRepository) GetAll() ([]model.CourseDetail, error) {
	rows, err := c.db.Query(common.GetAllDataActiveCd, false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courseDetails []model.CourseDetail
	for rows.Next() {
		var courseDetail model.CourseDetail
		err := rows.Scan(
			&courseDetail.CourseDetailID,
			&courseDetail.CourseID,
			&courseDetail.CourseChapter,
			&courseDetail.CourseContent,
			&courseDetail.CreatedAt,
			&courseDetail.UpdatedAt,
			&courseDetail.IsDeleted,
		)
		if err != nil {
			return nil, err
		}
		courseDetails = append(courseDetails, courseDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courseDetails, nil
}

func NewCourseDetailRepository(db *sql.DB) CourseDetailRepository {
	return &courseDetailRepository{db: db}
}