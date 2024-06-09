package main

import (
	"context"
	"log"

	"gerello/config"
	"gerello/internal/adapters/primary/api"
	"gerello/internal/adapters/secondary/postgres"
	"gerello/internal/adapters/secondary/postgres/projects"
	"gerello/internal/core/services"
)

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

	projectsRepo := projects.New(pg.Pool)
	projectsServices := services.New(projectsRepo)

	err = api.New(*projectsServices, "8080").Run()
	if err != nil {
		log.Printf("%v", err)
		return
	}
}
