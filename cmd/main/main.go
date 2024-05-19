package main

import (
	"context"
	"log"

	"gerello/internal/config"
	"gerello/internal/db"
)

// TODO: assembly the Project

func main() {
	ctx := context.Background()

	config, err := config.New()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	pg, err := db.New(ctx, config.ConnStrPostgres, config.MaxAttempts)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer pg.Pool.Close()
}
