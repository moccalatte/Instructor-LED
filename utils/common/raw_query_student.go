package common

const (
	CreateStudent = `insert into student(fullname,birth_date,birth_place,address,education,institution,job,email,password,is_deleted) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning student_id,fullname,birth_place,address,education,institution,job,email,password,is_deleted;`

	GetStudentByid = `select * from student where student_id = $1;`

	UpdateStudentbyId = `update student set fullname = $1, shortname = $2, birth_date = $3,birth_place=$4,address=$5,education=$6,institution=$7,job=$8,email=$9,password=$10,is_deleted=$11 where student_id = $12 returning student_id,fullname,shortname,birth_date,birth_place,address,education,institution,job,email,password,is_deleted;`

	DeleteStudentById = `update student set is_deleted = $1 where student_id = $2 returning student_id,fullname,shortname,birth_date,birth_place,address,education,institution,job,email,password,is_deleted;`
)
