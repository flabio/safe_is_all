package main

import (
	"log"

	"github.com/safe/auth"
	"github.com/safe/city"
	"github.com/safe/course"
	_ "github.com/safe/docs"
	"github.com/safe/emergencycontact"
	"github.com/safe/handlers"
	"github.com/safe/middleware"
	"github.com/safe/parentesco"
	"github.com/safe/rol"
	"github.com/safe/school"
	"github.com/safe/states"
	"github.com/safe/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration
	// Ruta para la documentación Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // swagger.HandlerDefault es el handler predeterminado
	// default
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/", // Specify allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true, // Allow credentials
	}))

	// Ruta para el login
	app.Post("/login", handlers.Login)

	// Grupo de rutas protegidas
	app.Use("/api", middleware.Protected())
	//router
	auth.NewAuthRouter(app)
	rol.NewRolRouter(app)
	parentesco.NewParentescoRouter(app)
	emergencycontact.NewEmergencyContactRouter(app)
	states.NewStatesRouter(app)
	user.NewUserRouter(app)
	school.NewSchoolRouter(app)
	course.NewCourseRouter(app)
	course.NewTopicRouter(app)
	course.NewLanguageRouter(app)
	city.NewCityRouter(app)
	// Start server
	log.Println("Server listening on port 8080")
	log.Fatal(app.Listen(":8080"))
}
