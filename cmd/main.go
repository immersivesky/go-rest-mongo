package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/immersivesky/go-rest-mongo/internal/adapters/handlers"
  "github.com/immersivesky/go-rest-mongo/internal/adapters/repository"
)

func main() {
  app := fiber.New()

  db := repository.NewDB("mongodb://mongouser:mongopass@localhost:27017")
  handlers := handlers.NewHTTPHandlers(db)

  app.Get("/chat/:id", handlers.GetChat)
  app.Post("/chat/:id", handlers.CreateChat)

  app.Listen(":3200")
}
