package model

import (
	"context"
	"time"

	"github.com/Hackathon22-Winter-03/backend/utils"
)

type Problem struct {
	ID        string     `json:"id" db:"id"`
	CreatorID string     `json:"creatorId" db:"creator_id"`
	Score     int        `json:"score" db:"score"`
	Title     string     `json:"title" db:"title"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

func GetProblems(ctx context.Context) ([]Problem, error) {
	problems := []Problem{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT `id`, `creator_id`, `score`, `title`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM problems",
	)
	if err != nil {
		return problems, err
	}
	return problems, nil
}

func GetProblem(ctx context.Context, problemID string) (Problem, error) {
	p := Problem{}
	err := dbx.GetContext(
		ctx,
		&p,
		"SELECT `id`, `creator_id`, `score`, `title`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM problems "+
			"WHERE `id` = ?",
		problemID,
	)
	if err != nil {
		return p, err
	}
	return p, nil
}

func TryCreateProblemHandler(ctx context.Context, creatorID string, score int, title string) (string, error) {
	id, err := utils.GenerateID()
	if err != nil {
		return "", err
	}
	_, err = dbx.ExecContext(
		ctx,
		"INSERT INTO problems (`id`, `creator_id`, `score`, `title`) VALUES (?, ?, ?, ?)",
		id,
		creatorID,
		score,
		title,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
