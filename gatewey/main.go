package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/safe/course"
	"github.com/safe/emergencycontact"
	"github.com/safe/parentesco"
	"github.com/safe/rol"
	"github.com/safe/school"
	"github.com/safe/states"
	"github.com/safe/user"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/", // Specify allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true, // Allow credentials
	}))
	//router
	rol.NewRolRouter(app)
	parentesco.NewParentescoRouter(app)
	emergencycontact.NewEmergencyContactRouter(app)
	states.NewStatesRouter(app)
	user.NewUserRouter(app)
	school.NewSchoolRouter(app)
	course.NewCourseRouter(app)
	// Start server
	log.Println("Server listening on port 8080")
	log.Fatal(app.Listen(":8080"))
}
