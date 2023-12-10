package common

const (
	CreateCourse     = `insert into course(course_name, course_detail_id, is_deleted ) values ($1,$2,$3) returning course_id, course_name, course_detail_id, is_deleted ;`
	GetCourseById    = `select * from course where course_id = $1;`
	UpdateCourseById = `update course set course_name = $1, course_detail_id = $2, is_deleted = $3 where course_id = $4 returning course_id, course_name, course_detail_id, is_deleted ;`
	DeleteCourseById = `update course set is_deleted = $1 where course_id = $2 returning course_id, course_name, course_detail_id, is_deleted ;`
)