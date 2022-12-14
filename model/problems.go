package model

import (
	"context"
	"time"

	"github.com/Hackathon22-Winter-03/backend/utils"
)

type Problem struct {
	ID        string     `json:"id" db:"id"`
	CreatorID string     `json:"creatorId" db:"creator_id"`
	Score     int        `json:"score" db:"score" form:"score"`
	Title     string     `json:"title" db:"title" form:"title"`
	Text      string     `json:"text" db:"text" form:"text"`
	Language  string     `json:"language" db:"language" form:"language"`
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
	Language    string     `json:"language" db:"language"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetProblems(ctx context.Context) ([]ProblemAggregate, error) {
	problems := []ProblemAggregate{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, `text`, `language`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
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
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, `text`, `language`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"JOIN users as u ON u.id = p.creator_id",
	)
	if err != nil {
		return problems, err
	}

	problemsDict := map[string]*ProblemAggregate{}
	for i := range problems {
		problemsDict[problems[i].ID] = &problems[i]
	}

	wa_problems, err := WAProblems(ctx, userID)
	if err != nil {
		return problems, err
	}
	for _, problem := range wa_problems {
		problemsDict[problem].Result = "WA"
	}

	ac_problems, err := ACProblems(ctx, userID)
	if err != nil {
		return problems, err
	}
	for _, problem := range ac_problems {
		problemsDict[problem].Result = "AC"
	}

	return problems, nil
}

func GetProblem(ctx context.Context, problemID string) (ProblemAggregate, error) {
	problem := ProblemAggregate{}
	err := dbx.GetContext(
		ctx,
		&problem,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, `text`, `language`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"JOIN users as u ON u.id = p.creator_id "+
			"WHERE p.`id` = ?",
		problemID,
	)
	if err != nil {
		return problem, err
	}
	return problem, nil
}

func GetProblemByUser(ctx context.Context, problemID string, userID string) (ProblemAggregate, error) {
	problem := ProblemAggregate{}
	err := dbx.GetContext(
		ctx,
		&problem,
		"SELECT p.`id`, `creator_id`, u.name as `creator_name`, p.`score`, `title`, `text`, `language`, p.`created_at`, p.`updated_at`, p.`deleted_at` "+
			"FROM problems as p "+
			"JOIN users as u ON u.id = p.creator_id "+
			"WHERE p.`id` = ?",
		problemID,
	)
	if err != nil {
		return problem, err
	}

	wa_problems, err := WAProblems(ctx, userID)
	if err != nil {
		return problem, err
	}
	for _, wa_problemID := range wa_problems {
		if wa_problemID == problemID {
			problem.Result = "WA"
			return problem, nil
		}
	}

	ac_problems, err := ACProblems(ctx, userID)
	if err != nil {
		return problem, err
	}
	for _, ac_problemID := range ac_problems {
		if ac_problemID == problemID {
			problem.Result = "AC"
			return problem, nil
		}
	}

	return problem, nil
}

func TryCreateProblem(ctx context.Context, creatorID string, score int, title string, text string, language string) (string, error) {
	id, err := utils.GenerateID()
	if err != nil {
		return "", err
	}
	_, err = dbx.ExecContext(
		ctx,
		"INSERT INTO problems (`id`, `creator_id`, `score`, `title`, `text`, `language`) VALUES (?, ?, ?, ?, ?, ?)",
		id,
		creatorID,
		score,
		title,
		text,
		language,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
