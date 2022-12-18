package model

import (
	/*
	   #cgo LDFLAGS: -L../lang -llang
	   #include <stdlib.h>
	   #include "../lang/simulate.h"
	*/
	"C"
	"context"
	"database/sql"
	"time"
	"unsafe"

	"github.com/Hackathon22-Winter-03/backend/utils"
)
import "fmt"

type Code struct {
	ID        string     `json:"id" db:"id"`
	UserID    string     `json:"userId" db:"user_id"`
	ProblemID string     `json:"problemId" db:"problem_id"`
	Code      string     `json:"code" db:"code"`
	Result    string     `json:"result" db:"result"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type CodeAggregate struct {
	ID          string     `json:"id" db:"id"`
	UserID      string     `json:"userId" db:"user_id"`
	UserName    string     `json:"userName" db:"user_name"`
	ProblemID   string     `json:"problemId" db:"problem_id"`
	ProblemName string     `json:"problemName" db:"problem_name"`
	Code        string     `json:"code" db:"code"`
	Result      string     `json:"result" db:"result"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" db:"deleted_at"`
}

type Testcase struct {
	ID        string `json:"id" db:"id"`
	ProblemID string `json:"problemId" db:"problem_id"`
	Input     string `json:"input" db:"input"`
	Output    string `json:"output" db:"output"`
}

func GetCodes(ctx context.Context, userID string, problemID string) ([]Code, error) {
	codes := []Code{}
	err := dbx.SelectContext(
		ctx,
		&codes,
		"SELECT `id`, `user_id`, `problem_id`, `code`, `result`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM codes "+
			"WHERE `user_id` = ? AND `problem_id` = ?",
		userID,
		problemID,
	)
	if err != nil {
		return codes, err
	}
	return codes, nil
}

func GetCode(ctx context.Context, problemID string, codeID string) (Code, error) {
	c := Code{}
	err := dbx.GetContext(
		ctx,
		&c,
		"SELECT `id`, `user_id`, `problem_id`, `code`, `result`, `created_at`, `updated_at`, `deleted_at` "+
			"FROM codes "+
			"WHERE `id` = ? AND `problem_id` = ?",
		codeID,
		problemID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, utils.ErrNotFound
		}
		return c, err
	}
	return c, nil
}

func SubmitCode(ctx context.Context, userID string, problemID string, code string) (string, error) {
	result := executeCode(ctx, problemID, code)

	id, err := utils.GenerateID()
	if err != nil {
		return "IE", err
	}
	_, err = dbx.ExecContext(
		ctx,
		"INSERT INTO codes (`id`, `user_id`, `problem_id`, `code`, `result`) VALUES (?, ?, ?, ?, ?)",
		id,
		userID,
		problemID,
		code,
		result,
	)
	if err != nil {
		return "IE", err
	}
	return result, nil
}

func executeCode(ctx context.Context, problemID string, code string) string {
	testcases := []Testcase{}
	err := dbx.SelectContext(
		ctx,
		&testcases,
		"SELECT `id`, `problem_id`, `input`, `output` FROM testcases WHERE `problem_id` = ?",
		problemID,
	)
	if err != nil {
		return "IE"
	}

	// Rust FFI
	cstr_code := C.CString(code)
	defer C.free(unsafe.Pointer(cstr_code))
	// cstr_lang := C.CString(language)
	// defer C.free(unsafe.Pointer(cstr_lang))

	for _, testcase := range testcases {
		cstr_input := C.CString(testcase.Input)
		defer C.free(unsafe.Pointer(cstr_input))
		if C.GoString(C.simulate(cstr_code, cstr_input /* cstr_lang */)) == testcase.Output {
			continue
		} else {
			return "WA"
		}
	}
	return "AC"
}

func StepExecute(ctx context.Context, code string, state string, language string) (string, error) {
	// Rust FFI
	cstr_code := C.CString(code)
	defer C.free(unsafe.Pointer(cstr_code))
	cstr_state := C.CString(state)
	defer C.free(unsafe.Pointer(cstr_state))
	cstr_lang := C.CString(language)
	defer C.free(unsafe.Pointer(cstr_lang))
	fmt.Println(state, cstr_state)

	return C.GoString(C.step_execute(cstr_code, cstr_state, cstr_lang)), nil
}

func ACProblems(ctx context.Context, userID string) ([]string, error) {
	problems := []string{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT `problem_id` FROM codes WHERE `user_id` = ? AND `result` = 'AC'",
		userID,
	)
	if err != nil {
		return problems, err
	}
	return problems, nil
}

func WAProblems(ctx context.Context, userID string) ([]string, error) {
	problems := []string{}
	err := dbx.SelectContext(
		ctx,
		&problems,
		"SELECT `problem_id` FROM codes WHERE `user_id` = ? AND `result` = 'WA'",
		userID,
	)
	if err != nil {
		return problems, err
	}
	return problems, nil
}
