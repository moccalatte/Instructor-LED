package common

const(
	CreateAttendance = `insert into(session_id ,student_id, attendance_student, updated_at, is_deleted ) values ($1,$2,$3,$4,$5) returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	GetAttendanceById = `select * from attendance where attendance_id = $2 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	UpdateAttendanceById = `update attendance set session_id=$1 ,student_id=$2, attendance_student=$3, updated_at = $4, is_deleted = $5 where attendance_id = $6 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`

	DeleteAttendanceById = `update attendance set is_deleted = $1 where attendance_id = $2 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;`
)