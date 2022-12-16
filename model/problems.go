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
	Text      string     `json:"text" db:"text"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type ProblemAggregate struct {
	ID          string     `json:"id" db:"id"`
	CreatorID   string     `json:"creatorId" db:"creator_id"`
	CreatorName string     `json:"creatorName" db:"creator_name"`
	Result      string     `json:"result" db:"result"`
	Score       int        `json:"score" db:"score"`
	Title       string     `json:"title" db:"title"`
	Text        string     `json:"text" db:"text"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetProblems(ctx context.Context) ([]ProblemAggregate, error) {
	problems := []ProblemAggregate{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"JOIN users as u ON u.id = p.creator_id",
	)
	if err != nil {
		return problems, err
	}
	return problems, nil
}

func GetProblemsByUser(ctx context.Context, userID string) ([]ProblemAggregate, error) {
	problems := []ProblemAggregate{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"JOIN users as u ON u.id = p.creator_id",
	)
	if err != nil {
		return problems, err
	}

	problemsDict := map[string]*ProblemAggregate{}
	for _, problem := range problems {
		problemsDict[problem.ID] = &problem
	}

	ac_problems, err := ACProblems(ctx, userID)
	if err != nil {
		return problems, err
	}
	for _, problem := range ac_problems {
		problemsDict[problem].Result = "AC"
	}

	wa_problems, err := WAProblems(ctx, userID)
	if err != nil {
		return problems, err
	}
	for _, problem := range wa_problems {
		problemsDict[problem].Result = "WA"
	}

	return problems, nil
}

func GetProblem(ctx context.Context, problemID string) (ProblemAggregate, error) {
	problem := ProblemAggregate{}
	err := dbx.GetContext(
		ctx,
		&problem,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"WHERE `id` = ? "+
			"JOIN users as u ON u.id = p.creator_id",
		problemID,
	)
	if err != nil {
		return problem, err
	}
	return problem, nil
}

func TryCreateProblemHandler(ctx context.Context, creatorID string, score int, title string, text string) (string, error) {
	id, err := utils.GenerateID()
	if err != nil {
		return "", err
	}
	_, err = dbx.ExecContext(
		ctx,
		"INSERT INTO problems (`id`, `creator_id`, `score`, `title`, `text`) VALUES (?, ?, ?, ?, ?)",
		id,
		creatorID,
		score,
		title,
		text,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
