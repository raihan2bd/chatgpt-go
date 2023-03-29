package models

import (
	"context"
	"time"
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
