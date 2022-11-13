package store

import (
	"fmt"
	"go/rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	r.store.db.QueryRow(fmt.Sprintf(
		"INSERT INTO users (email, password_hash) VALUES ('%s', '%s')",
		u.Email,
		u.PasswordHash))

	if err := r.store.db.QueryRow(fmt.Sprintf("SELECT `id` FROM `users` WHERE `email` = '%s'", u.Email)).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		fmt.Sprintf("SELECT `id`, `email`, `password_hash` FROM `users` WHERE `email` = '%s'", email)).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
	); err != nil {
		return nil, err
	}

	return u, nil
}
