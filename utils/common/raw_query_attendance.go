package common

const (
	CreateAttendance = `insert into attendance(session_id ,student_id, attendance_student, updated_at, is_deleted ) values ($1,$2,$3,$4,$5) returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	GetAllDataActive = `select * from course_detail where is_deleted = $1;`

	GetAttendanceById = `select * from attendance where attendance_id = $1;`

	GetAttandanceBySessionId = `select * from attendance where session_id = $1;`

	UpdateAttendanceById = `update attendance set session_id=$1 ,student_id=$2, attendance_student=$3, updated_at = $4, is_deleted = $5 where attendance_id = $6 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	DeleteAttendanceById = `update attendance set is_deleted = $1 where attendance_id = $2 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	GetAllAttendance = `SELECT * FROM attendance WHERE is_deleted = false;`
)
