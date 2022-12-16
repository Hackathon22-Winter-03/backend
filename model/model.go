package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Hackathon22-Winter-03/backend/utils"
)

var (
	dbx *sqlx.DB
)

func InitDB() (*sqlx.DB, error) {
	db, err := connectDB(false)
	if err != nil {
		fmt.Printf("failed to connect to db: %v", err)
	}
	dbx = db
	fmt.Println(dbx)

	return dbx, err
}

func connectDB(batch bool) (*sqlx.DB, error) {
	user := utils.GetEnv("DB_USER", "root")
	pass := utils.GetEnv("DB_PASS", "password")
	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "3306")
	dbname := utils.GetEnv("DB_NAME", "bocchi")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s&multiStatements=%t&interpolateParams=true",
		user,
		pass,
		host,
		port,
		dbname,
		"Asia%2FTokyo",
		batch,
	)
	dbx, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// プール内に保持できるアイドル接続数の制限を設定 (default: 2)
	dbx.SetMaxIdleConns(1024)
	// 接続してから再利用できる最大期間
	dbx.SetConnMaxLifetime(0)
	// アイドル接続してから再利用できる最大期間
	dbx.SetConnMaxIdleTime(0)

	return dbx, nil
}
