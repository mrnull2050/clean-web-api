package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModels struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (u *UserModels) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
INSERT INTO users (email, name, password)
VALUES ($1, $2, $3)
RETURNING id
`
	return u.DB.QueryRowContext(ctx, query, user.Email, user.Name, user.Password).Scan(&user.Id)

}
