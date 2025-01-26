package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AshtonMH474/Toytopia/config"
	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/AshtonMH474/Toytopia/routes"
	"github.com/AshtonMH474/Toytopia/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func main() {
	// Load environment configuration
	cfg := config.LoadConfig()

	// Check if in production
	enviroment := os.Getenv("NODE_ENV")
	isProduction := enviroment == "production"

	// Handle CLI commands for seeding or migrations
	handleCLICommands(cfg)

	// Connect to the database
	database.ConnectDB(cfg)

	// Create a new Fiber app
	app := fiber.New()

	// Apply global middleware
	setupGlobalMiddleware(app, isProduction)

	// Set up routes after middleware
	setupRoutes(app)

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupGlobalMiddleware(app *fiber.App, isProduction bool) {
	// CORS middleware for development
	if !isProduction {
		app.Use(cors.New())
	}

	// Logger middleware for debugging
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} (${latency})\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// Security middleware
	app.Use(helmet.New())

	// Set Cross-Origin Resource Policy header
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cross-Origin-Resource-Policy", "cross-origin")
		return c.Next()
	})

}

func handleCLICommands(cfg config.Config) {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed":
			database.ConnectDB(cfg)
			if len(os.Args) > 2 {
				if os.Args[2] == "all" {
					seeders.SeedAll()
					log.Println("Seeding completed.")
				} else if os.Args[2] == "undo" {
					seeders.UndoSeeds()
					log.Println("Undoing Seeds completed.")
				}
			}
			os.Exit(0)
		case "migrate":
			database.ConnectDB(cfg)
			if len(os.Args) > 2 && os.Args[2] == "all" {
				log.Println("Running Migrations")
				database.Database.Db.AutoMigrate(models.User{}, models.Toy{}, models.Wishlist{})
			}
			os.Exit(0)
		}
	}
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	// user routes
	app.Post("/api/users", routes.SignupHandler())
	app.Post("/api/login", routes.LoginHandler)
	app.Delete("/api/logout", routes.Logout)
	app.Get("/api/users/current", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// toys routes
	app.Get("/api/toys", routes.SearchToys)
	app.Post("/api/toys", routes.CreateToy)
	app.Put("/api/toys/:id", routes.UpdateToy)
	app.Delete("/api/toys/:id", routes.DeleteToy)

	// wishlists routes
	app.Get("/api/wishlists", routes.AllWishlists)
	app.Get("/api/wishlists/:id", routes.GetWishlist)
	app.Post("/api/wishlists", routes.CreateWishlist)
	app.Post("/api/wishlists/:Id", routes.AddToy)
	app.Delete("/api/wishlists/:wishlistId/toys/:toyId", routes.RemoveToy)
	app.Put("/api/wishlists/:id", routes.UpdateWishlist)
	app.Delete("/api/wishlists/:id", routes.DeleteWishlist)
}
