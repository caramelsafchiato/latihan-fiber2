package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"latihan-fiber2/app/repository"
	"latihan-fiber2/config"
	"latihan-fiber2/database"
)

// requireJSON memastikan request body formatnya JSON (middleware)
func requireJSON(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPut || c.Method() == fiber.MethodPatch {
		if c.Get("Content-Type") != "application/json" {
			return fail(c, fiber.StatusUnsupportedMediaType, "hanya menerima application/json")
		}
	}
	return c.Next()
}

func main() {
	// 1. Konfigurasi
	config.LoadEnv() // Memuat variabel dari .env

	// 2. Koneksi basis data
	pool, err := database.NewPool(context.Background()) // Membuat connection pool
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer pool.Close()

	// 3. Perakitan: pool -> repository -> handler
	// Handler tidak lagi memakai variabel global, melainkan disuntikkan dari luar.
	studentRepository := repository.NewStudentRepository(pool)
	studentHandler := NewStudentHandler(studentRepository)

	// 4. Aplikasi (Sama seperti pertemuan 2, dengan middleware)
	app := fiber.New(fiber.Config{})
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.UserContext(), 2*time.Second)
		defer cancel()

		// Kesehatan layanan kini ikut bergantung pada basis data
		if err := pool.Ping(ctx); err != nil {
			return fail(c, fiber.StatusServiceUnavailable, "database tidak dapat dihubungi")
		}
		return ok(c, "server dan database berjalan", nil)
	})

	// Mengelompokkan endpoint untuk students
	s := api.Group("/students", requireJSON)
	s.Get("/", studentHandler.List)
	s.Get("/:id", studentHandler.Get)
	s.Post("/", studentHandler.Create)
	s.Put("/:id", studentHandler.Replace)
	s.Patch("/:id", studentHandler.Patch)
	s.Delete("/:id", studentHandler.Delete)

	port := config.GetEnv("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}