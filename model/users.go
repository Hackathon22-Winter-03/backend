package model

import "context"

type User struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Score     int    `json:"score" db:"score"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt *int64 `json:"deletedAt,omitempty" db:"deleted_at"`
}

func GetUser(ctx context.Context, userID string) (User, error) {
	user := User{}
	err := dbx.GetContext(
		ctx,
		&user,
		"SELECT `id`, `name`, `score`, `created_at`, `updated_at`, `deleted_at` "+
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
		"SELECT `id`, `name`, `score`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM users",
	)
	if err != nil {
		return users, err
	}
	return users, nil
}
