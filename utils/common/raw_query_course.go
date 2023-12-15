package common

const (
	CreateCourse     = `insert into course (course_name, description, updated_at, is_deleted) values ($1,$2,$3,$4) returning  course_id, course_name, description, created_at, updated_at, is_deleted ;`
	GetCourseById    = `select * from course where course_id = $1;`
	GetAllDataC      = `select * from course where is_deleted = $1;`
	UpdateCourseById = `update course set course_name = $1, description = $2, updated_at = $3, is_deleted = $4 where course_id = $5 returning  course_id, course_name, description, created_at, updated_at, is_deleted;`
	DeleteCourseById = `update course set is_deleted = $1 where course_id = $2 returning  course_id, course_name, description, created_at, updated_at, is_deleted;`
	GetAllCourse     = `SELECT course_id, course_name, description, created_at, is_deleted FROM course WHERE is_deleted = false;`
)
