package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/emergencycontact"
	"github.com/safe/parentesco"
	"github.com/safe/rol"
	"github.com/safe/school"
	"github.com/safe/states"
	"github.com/safe/user"
)

func main() {
	app := fiber.New()
	//router
	rol.NewRolRouter(app)
	parentesco.NewParentescoRouter(app)
	emergencycontact.NewEmergencyContactRouter(app)
	states.NewStatesRouter(app)
	user.NewUserRouter(app)
	school.NewSchoolRouter(app)
	// Start server
	log.Println("Server listening on port 8080")
	log.Fatal(app.Listen(":8080"))
}
