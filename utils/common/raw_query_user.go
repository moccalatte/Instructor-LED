package common

const (
	CreateUser = `insert into users(fullname,role,email,password, updated_at, is_deleted) 
	values ($1,$2,$3,$4,$5,$6) returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;`

	GetUserById = `select * from users where user_id = $1;`

	UpdateUser = `update users set fullname = $1,role = $2,email = $3,password = $4, updated_at = $5, is_deleted = $6 where user_id = $7 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;`

	DeleteUser    = `update users set is_deleted = $1 where user_id = $2 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;`
	GetByFullname = `select user_id, fullname, role, email, password, is_deleted from users where fullname = $1 OR email = $1;`
)
