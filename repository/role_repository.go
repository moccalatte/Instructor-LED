package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
)

type RoleRepository interface {
	Create(payload model.Role) (model.Role, error)
	GetById(id int) (model.Role, error)
	Update(id int) (model.Role, error)
	Delete(id int) (model.Role, error)
}

type roleRepository struct {
	db *sql.DB
}

func (r *roleRepository) Create(payload model.Role) (model.Role, error) {
}

func (r *roleRepository) GetById(id int) (model.Role, error) {
}

func (r *roleRepository) Update(id int) (model.Role, error) {
}

func (r *roleRepository) Delete(id int) (model.Role, error) {
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{db: db}
}
