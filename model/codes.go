package model

import (
	"context"

	"github.com/Hackathon22-Winter-03/backend/utils"
)

type Code struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"userId" db:"user_id"`
	ProblemID string `json:"problemId" db:"problem_id"`
	Code      string `json:"code" db:"code"`
	Answer    string `json:"answer" db:"answer"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt *int64 `json:"deletedAt,omitempty" db:"deleted_at"`
}

func GetCodesFromUser(ctx context.Context, userID string) ([]Code, error) {
	codes := []Code{}
	err := dbx.SelectContext(
		ctx,
		&codes,
		"SELECT `id`, `user_id`, `problem_id`, `code`, `answer`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM codes "+
			"WHERE `user_id` = ?",
		userID,
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
		"SELECT `id`, `user_id`, `problem_id`, `code`, `answer`, `created_at`, `updated_at`, `deleted_at` "+
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
	id, err := utils.GenerateID()
	if err != nil {
		return "IE", err
	}
	_, err = dbx.ExecContext(
		ctx,
		"INSERT INTO codes (`id`, `user_id`, `problem_id`, `code`) VALUES (?, ?, ?)",
		id,
		userID,
		problemID,
		code,
	)
	if err != nil {
		return "IE", err
	}
	return "AC", nil
}
