package main

import (
	"fmt"
	"log"

	"kg_drive/app"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fiberApp := fiber.New(fiber.Config{
		ReadTimeout: app.RequestTimeout,
		BodyLimit:   10 * 1024 * 1024, // 10MB body limit
	})

	fiberApp.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	calcService := app.NewCalculatorService()
	fiberApp.Get("/numbers/:numberid", calcService.HandleNumberRequest)

	fiberApp.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	port := ":9876"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Printf("Available endpoints:\n")
	fmt.Printf("  - GET /numbers/p (Prime numbers)\n")
	fmt.Printf("  - GET /numbers/f (Fibonacci numbers)\n")
	fmt.Printf("  - GET /numbers/e (Even numbers)\n")
	fmt.Printf("  - GET /numbers/r (Random numbers)\n")
	fmt.Printf("  - GET /health (Health check)\n")
	log.Fatal(fiberApp.Listen(port))
} 