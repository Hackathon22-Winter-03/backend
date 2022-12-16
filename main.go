package main

import (
	"fmt"

	"github.com/Hackathon22-Winter-03/backend/model"
	"github.com/Hackathon22-Winter-03/backend/router"
)

func main() {
	_, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("db error: %w", err))
	}

	_, err = router.SetupRouting()
	if err != nil {
		panic(fmt.Errorf("routing error: %w", err))
	}
}
