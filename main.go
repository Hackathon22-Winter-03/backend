package main

import (
	"fmt"

	"github.com/Hackathon22-Winter-03/backend/model"
	"github.com/Hackathon22-Winter-03/backend/router"
)

func main() {
	e := router.SetupRouting()

	_, err := model.InitDB(e)
	if err != nil {
		panic(fmt.Errorf("db error: %w", err))
	}
}
