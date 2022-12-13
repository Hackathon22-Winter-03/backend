package model

import "context"

type Problem struct {
	ID        int64  `json:"id" db:"id"`
	CreatorID string `json:"creatorId" db:"creator_id"`
	Score     int    `json:"score" db:"score"`
	Title     string `json:"title" db:"title"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt *int64 `json:"deletedAt,omitempty" db:"deleted_at"`
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

func GetProblem(ctx context.Context, problemID int64) (Problem, error) {
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

func TryCreateProblemHandler(ctx context.Context, creatorID string, score int, title string) (int64, error) {
	res, err := dbx.ExecContext(
		ctx,
		"INSERT INTO problems (`creator_id`, `score`, `title`) VALUES (?, ?, ?)",
		creatorID,
		score,
		title,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}