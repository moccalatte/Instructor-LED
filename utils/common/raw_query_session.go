package common

const (
	CreateSession = `INSERT INTO session(date_session,trainer_admin_id, is_deleted) VALUES ($1, $2, $3);`

	GetSession = `SELECT session_id, date, trainer_admin_id FROM session WHERE session_id = $1;`

	UpdateSessionById = `UPDATE session SET date = $1, trainer_admin_id = $2 WHERE session_id = $3`
	DeleteSessionById = `update session set is_deleted = $1 where session_id = $2 returning session_id, date, trainer_admin_id, is_deleted ;`
)
