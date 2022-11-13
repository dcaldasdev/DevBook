package repository

import (
	"api/internal/entity"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(entity.User) (uint64, error) {
	return 0, nil
}
