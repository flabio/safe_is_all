package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/routers"
)

func main() {
	app := fiber.New()

	routers.NewCourseRouter(app)
	app.Listen(":3007")

}
