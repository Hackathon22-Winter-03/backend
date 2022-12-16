package model

import (
	"context"
	"time"
)

type User struct {
	ID        string     `json:"id" db:"id" form:"id"`
	Name      string     `json:"name" db:"name" form:"name"`
	Comment   string     `json:"comment" db:"comment" form:"comment"`
	Score     int        `json:"score" db:"score"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetUser(ctx context.Context, userID string) (User, error) {
	user := User{}
	err := dbx.GetContext(
		ctx,
		&user,
		"SELECT `id`, `name`, `comment`, `score`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM users "+
			"WHERE `id` = ?",
		userID,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUsers(ctx context.Context) ([]User, error) {
	users := []User{}
	err := dbx.SelectContext(
		ctx,
		&users,
		"SELECT `id`, `name`, `comment`, `score`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM users",
	)
	if err != nil {
		return users, err
	}
	return users, nil
}

func CreateUser(ctx context.Context, userID string, name string, comment string) error {
	_, err := dbx.ExecContext(
		ctx,
		"INSERT INTO users (`id`, `name`, `comment`, `score`) VALUES (?, ?, ?, 0)",
		userID,
		name,
		comment,
	)
	if err != nil {
		return err
	}
	return nil
}
