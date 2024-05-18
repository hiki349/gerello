package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"gerello/internal/domain/services"
)

type Server interface {
	Start()
}

type RestServer struct {
	services services.ProjectService
}

func New(services services.ProjectService) *RestServer {
	return &RestServer{
		services: services,
	}
}

func (s *RestServer) Start() {
	app := fiber.New()
	api := app.Group("/api")

	projectsAPI := api.Group("/projects")

	s.createProjectsRouter(projectsAPI)

	log.Fatal(app.Listen(":7070"))
}
