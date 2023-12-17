package common

const (
	CreateStudent = `insert into student (fullname, birth_date, birth_place, address, education, institution, job, email, password, updated_at, is_deleted, role) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning student_id, fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at,is_deleted, role;`

	GetStudentByid    = `SELECT * FROM student WHERE student_id = $1;`
	GetAllDataStd     = `select * from student where is_deleted = $1;`
	UpdateStudentbyId = `update student set fullname = $1, birth_date = $2, birth_place = $3, address = $4, education = $5, institution = $6, job = $7, email = $8, password = $9, updated_at = $10, is_deleted = $11 where student_id = $12 returning student_id, fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted, role;`

	DeleteStudentById = `update student set is_deleted = $1 where student_id = $2 returning student_id, fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted, role;`

	GetByStudentEmail = `select student_id, fullname, birth_date, birth_place, address, education, institution, job, role,email, password, created_at, updated_at, is_deleted from student where fullname = $1 OR email = $1;`
)
