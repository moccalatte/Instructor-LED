package common

const (
	CreateQuestion         = `insert into question (session_id, student_id, trainer_id, title, description, course_id, image, answer, status,  updated_at, is_deleted) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;`
	GetQuestionById        = `select * from question where question_id = $1;`
	GetAllDataQ            = `select * from question where is_deleted = $1;`
	GetQuestionByStudentId = `select * from question where student_id = $1`
	GetAllQuestion         = `SELECT question_id,session_id, student_id, trainer_id, title, description, course_id, image, answer, status, is_deleted FROM question WHERE is_deleted = false;`

	UpdateQuestionById = `update question set session_id = $1, student_id = $2, trainer_id = $3, title=$4, description=$5, course_id = $6, image = $7, answer = $8, status = $9, updated_at = $10, is_deleted = $11 where question_id = $12 returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;`

	DeleteQuestionById = `update question set is_deleted = $1 where question_id = $2 returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;`

	AnswerQuestionById = `update question set answer = $1, updated_at = $2 where question_id = $3 returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;`

	SaveImagePath = `"UPDATE question SET image_path = $1 WHERE question_id = $2`

	GetImagePathById = `
    SELECT image FROM questions
    WHERE question_id = $1
`
)
