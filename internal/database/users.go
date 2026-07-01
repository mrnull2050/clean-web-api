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

func (u *UserModels) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := u.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModels) Get(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	return m.getUser(query, id)
}
func (m *UserModels) GetByEmail(email string) (*User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	return m.getUser(query, email)
}
