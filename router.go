package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/ppreeper/addictgo/api"
)

func basicAnswer(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{"data": "other"})
}

func Other(app *fiber.App) {
	app.Get("/other", basicAnswer)
	app.Get("/all", basicAnswer)
	app.Get("/find/:filter", api.FindFilterGet)
	app.Get("/status", basicAnswer)
	app.Get("/stack", basicAnswer)
	app.Get("/monitor", monitor.New(monitor.Config{Title: "ADDict Metrics Page"}))
}

func OU(app *fiber.App) {
	app.Get("/other", basicAnswer)
	app.Get("/ou", basicAnswer)
	app.Post("/ou", basicAnswer)
	app.Get("/ou/:ou", basicAnswer)
	app.Delete("/ou/:ou", basicAnswer)
	app.Get("/ou/:ou/exists", basicAnswer)
}

func Group(app *fiber.App) {
	app.Get("/group", basicAnswer)
	app.Post("/group", basicAnswer)
	app.Get("/group/:group", basicAnswer)
	app.Delete("/group/:group", basicAnswer)
	app.Get("/group/:group/exists", basicAnswer)
	app.Post("/group/:group/user/:user", basicAnswer)
	app.Delete("/group/:group/user/:user", basicAnswer)
}

func User(app *fiber.App) {
	app.Get("/user", basicAnswer)
	app.Get("/user", basicAnswer)
	app.Get("/user/:user", basicAnswer)
	app.Delete("/user/:user", basicAnswer)
	app.Get("/user/:user/exists", basicAnswer)
	app.Get("/user/:user/member-of/:group", basicAnswer)
	app.Post("/user/:user/authenticate", basicAnswer)
	app.Put("/user/:user/password", basicAnswer)
	app.Put("/user/:user/password-never-expires", basicAnswer)
	app.Put("/user/:user/password-expires", basicAnswer)
	app.Put("/user/:user/enable", basicAnswer)
	app.Put("/user/:user/disable", basicAnswer)
	app.Put("/user/:user/move", basicAnswer)
	app.Put("/user/:user/unlock", basicAnswer)
}
