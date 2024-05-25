package main

import (
	"context"
	"log"

	"gerello/config"
	"gerello/internal/adapters/postgres"
)

// TODO: assembly the Project

func main() {
	ctx := context.Background()

	config, err := config.New()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	pg, err := postgres.New(ctx, config.ConnStrPostgres, config.MaxAttempts)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer pg.Pool.Close()
}
