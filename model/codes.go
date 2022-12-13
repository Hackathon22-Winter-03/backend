package model

import "context"

type Code struct {
	ID        int64  `json:"id" db:"id"`
	UserID    int64  `json:"userId" db:"user_id"`
	Code      int    `json:"code" db:"code"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt *int64 `json:"deletedAt,omitempty" db:"deleted_at"`
}

func GetCodesFromUser(ctx context.Context, userName string) ([]Code, error) {
	codes := []Code{}
	err := dbx.SelectContext(
		ctx,
		&codes,
		"SELECT `id`, `user_id`, `code`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM codes "+
			"WHERE `user_name` = ?",
		userName,
	)
	if err != nil {
		return codes, err
	}
	return codes, nil
}

func GetCode(ctx context.Context, codeID int64) (Code, error) {
	c := Code{}
	err := dbx.GetContext(
		ctx,
		&c,
		"SELECT `user_name`, `plain_code`, `stdin`, `title`, `options` "+
			"FROM codes "+
			"WHERE `id` = ?",
		codeID,
	)
	if err != nil {
		return c, err
	}
	return c, nil
}

func SubmitCode(ctx context.Context, userID string, problemID int64, code string) (string, error) {
	_, err := dbx.ExecContext(
		ctx,
		"INSERT INTO codes (`user_id`, `problem_id`, `code`) VALUES (?, ?, ?)",
		userID,
		problemID,
		code,
	)
	if err != nil {
		return "IE", err
	}
	return "AC", nil
}
