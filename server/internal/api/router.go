package api

import "github.com/gofiber/fiber/v2"

func (s *RestServer) createProjectsRouter(projectsAPI fiber.Router) {
	// api/projects/
	projectsAPI.Get("/", s.fetchAll)

	// api/projects/:id
	projectsAPI.Get("/:id", s.fetchByID)

	// api/projects/:id
	projectsAPI.Put("/:id", s.Update)

	// api/projects/:id
	projectsAPI.Delete("/:id", s.Delete)

	// api/projects/
	projectsAPI.Post("/", s.Create)
}
