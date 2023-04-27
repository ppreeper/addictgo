package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (a *API) Other(app *fiber.App) {
	app.Get("/other", func(c *fiber.Ctx) error {
		// Get all other objects
		// GET /other
		// Pulls all non-user/group Active Directory objects
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// var attributes []string
		// filter := fmt.Sprintf("(&(!(objectClass=user))(!(objectClass=group)))")
		// d := make(map[string]interface{})

		// fields := strings.Split(c.Query("_fields"), ",")
		// if len(c.Query("_fields")) > 0 {
		// 	for _, a := range fields {
		// 		attributes = append(attributes, a)
		// 	}
		// }

		// q := c.Query("_q")
		// start := c.Query("_start")
		// end := c.Query("_end")
		// fmt.Println(c.Params("user"), fields, q, start, end)

		// sr := svr.Search(filter, attributes)
		// d["other"] = sr["data"]
		// c.JSON(d)

		// Javascript
		// const config = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.other().get(config));
		// respond(res, error, response);

		args := new(LDAPArgs)
		if err := c.QueryParser(args); err != nil {
			return fmt.Errorf("invalid query args")
		}
		filter := "(&(!(objectClass=user))(!(objectClass=group)))"
		d := make(map[string]interface{})
		sr := a.Search(filter, args)
		d["other"] = sr["data"]
		return c.JSON(d)
	})
}

func (a *API) OU(app *fiber.App) {

}

func (a *API) Group(app *fiber.App) {

}

func (a *API) User(app *fiber.App) {

}
