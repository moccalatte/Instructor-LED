package common

const (
	GetQuestionById    = `SELECT attendance_id, student_id, course_id, trainer_admin_id, status FROM attendance WHERE attendance_id = $1`
	CreateQuestion     = `INSERT INTO attendance (student_id, course_id, trainer_admin_id, status) VALUES ($1, $2, $3, $4, $5)`
	UpdateQuestionById = `UPDATE attendance SET student_id = $1, course_id = $2, trainer_admin_id = $3, status = $4 WHERE attendance_id = $5`
	DeleteQuestionById = `update attendance set is_deleted = $1 where attendance_id = $2 returning attendance_id, student_id, course_id, trainer_admin_id, status,is_deleted ;`
)
