package common

const (
	CreatCourseDetail = `insert into course_detail (course_id, course_chapter, course_content, updated_at, is_deleted) values ($1,$2,$3,$4,$5) returning course_detail_id, course_id, course_chapter, course_content, created_at, updated_at, is_deleted;`

	GetCourseDetailById = `select * from course_detail where course_detail_id = $1;`

	GetAllDataActiveCd = `select * from course_detail where is_deleted = $1;`
	UpdateCourseDetailByid = `update course_detail set course_id = $1, course_chapter = $2, course_content = $3, updated_at = $4, is_deleted = $5 where course_detail_id = $6 returning course_detail_id, course_id, course_chapter, course_content, created_at, updated_at, is_deleted;`

	DeleteCourseDetailById = `update course_detail set is_deleted = $1 where course_detail_id = $2 returning course_detail_id, course_id, course_chapter, course_content, created_at, updated_at, is_deleted;`
)
