package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration
	routers.NewAuthRouter(app)
	routers.NewUserRouter(app)
	app.Listen(":3005")

}
