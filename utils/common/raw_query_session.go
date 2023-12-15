package common

const (
	CreateSesion = `insert into session(title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted) values ($1,$2,$3,$4,$5,$6,$7,$8,$9, $10) returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;`

	GetSessionById = `select * from session where session_id = $1;`
	GetAllDatas    = `select * from session where is_deleted = $1;`

	UpdateSessionById = `update session set title = $1, description = $2, session_date = $3, session_time = $4, session_link = $5, trainer_id = $6,  note = $7,updated_at = $8, is_deleted = $9 where session_id = $10 returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;`

	UpdateNote = `update session set note = $1, updated_at = $2 WHERE session_id = $3 returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;`

	DeleteSessionById = `update session set is_deleted = $1 where session_id = $2 returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;`

	GetAllSession = `SELECT session_id, title, description, session_date, session_time, session_link, trainer_id, note,is_deleted FROM session WHERE is_deleted = false ORDER BY session_date ASC;`
)
