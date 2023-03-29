package models

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (m *DBModel) AddNewUser(user User, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		insert into users (first_name, last_name, email, password, access_level, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		user.FirstName,
		user.LastName,
		user.Email,
		hash,
		user.AccessLevel,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) Authenticate(email string, pass string) (int, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	var userEmail string
	var password string
	var access_level int

	query := `
	select id, email, password, access_level from users
	where email = $1;`

	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&id, &userEmail, &password, &access_level)
	if err != nil {
		return id, access_level, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, 0, errors.New("email or password doesn't match")
	} else if err != nil {
		return 0, 0, errors.New("invalid credentials")
	}
	if access_level < 1 {
		return 0, access_level, errors.New("your account is temporary disabled")
	}

	return id, access_level, nil

}
