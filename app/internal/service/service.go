package service

import (
	"crud/internal/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	app     *fiber.App
	storage Storage
}

func NewService(app *fiber.App, storage Storage) service {
	return service{app: app, storage: storage}
}

func (s *service) Start() {
	s.app.Post("/users", func(ctx *fiber.Ctx) error {
		var req models.UserRequest
		if err := ctx.BodyParser(&req); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}
		id, err := s.storage.Create(models.User{
			Name:  req.Name,
			Email: req.Email,
		})
		if err != nil {
			return fmt.Errorf("create in storage: %w", err)
		}
		return ctx.JSON(models.UserResponse{ID: id})
	})
	s.app.Get("/users", func(ctx *fiber.Ctx) error {
		list, err := s.storage.List()
		if err != nil {
			log.Println(err)
			return err
		}
		return ctx.JSON(list)
	})
	s.app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		v, err := s.storage.GetUser(ctx.Params("id"))
		if err != nil {
			return fmt.Errorf("user %q not found", ctx.Params("id"))
		}
		return ctx.JSON(v)
	})
	s.app.Patch("/users/:id", func(ctx *fiber.Ctx) error {
		var req models.UserRequest
		if err := ctx.BodyParser(&req); err != nil {
			return fmt.Errorf("body parser: %w", err)
		}
		err := s.storage.Update(ctx.Params("id"), req.Email, req.Name)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}
		return nil
	})
	s.app.Delete("/users/:id", func(ctx *fiber.Ctx) error {
		err := s.storage.Delete(ctx.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(err)
		return err
	})
	log.Fatal(s.app.Listen(":80"))
}
