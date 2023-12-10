package common

const (
	CreateUser = `insert into users(fullname,role,email,password,is_deleted) 
	values ($1,$2,$3,$4,$5) returning user_id, fullname,role,email,password,is_deleted
	;`

	GetUserById = `select user_id, fullname, role, email, password, is_deleted from users where user_id = $1;`
	// GetUserById = `SELECT user_id, fullname, role, email, password, is_deleted FROM users WHERE is_deleted = $1`

	UpdateUser = `update users set fullname=$1,role=$2,email=$3,password = $4,is_deleted=$5 where user_id = $6 returning user_id, name,role,email,password,is_deleted;`

	DeleteUser = `update users set is_deleted=$1 where user_id = $2 returning user_id, name,role,email,password,is_deleted;`

	GetByFullname = `select user_id, fullname, role, email, password, is_deleted from users where fullname = $1 OR email = $1;`
)
