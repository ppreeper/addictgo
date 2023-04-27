package addictgo

import (
	"github.com/gofiber/fiber/v2"
)

func (a *API) OU(app *fiber.App) {
	// Get all OUs
	// GET /ou
	// Pulls all Organization Units in Active Directory, with filters.
	app.Get("/ou", func(c *fiber.Ctx) error {
		qp := new(LDAPArgs)
		if err := c.QueryParser(qp); err != nil {
			return err
		}

		// var attributes []string
		// filter := "(objectClass=organizationalUnit)"

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

		// sr := a.Search(filter, attributes)
		// c.JSON(sr)

		// Javascript
		// const filters = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.ou().get(filters));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "sr"})
	})

	app.Post("/ou", func(c *fiber.Ctx) error {
		// Add an OU
		// POST /ou
		// Adds a new OU to Active Directory
		// Request Body
		// name			Name of the OU as displayed
		// Optional
		// description	Descripton of the OU
		// location		Relative AD Position separated by /

		// queries: {
		// 		"name": {
		// 				description: 'Name of the OU as displayed.',
		// 				optional: false
		// 		},
		// 		"description": "Description of the OU.",
		// 		"location": "Relative AD position, separated by /.",
		// }

		// queries:
		// 	name:
		// 	description: "Name of the OU as displayed."
		// 	optional: false
		// 	description: "Description of the OU."
		// 	location: "Relative AD position, separated by /."

		// Javascript
		// let [error, response] = await wrapAsync(ad.ou().add(req.body));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/ou/:ou", func(c *fiber.Ctx) error {
		// Get a single OU
		// GET /ou/:ou
		// Pulls a single OU
		// URL Parameters
		// ou		The OU name
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// params: {ou: params.ou},
		// queries: filterQuery

		// parameters:
		// 	ou: params.ou
		// 	queries: filterQuery

		// var attributes []string
		// filter := fmt.Sprintf("(&(objectClass=organizationalUnit)(name=%s))", c.Params("ou"))

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
		// c.JSON(sr)

		// Javascript
		// let ou = req.params.ou;
		// const filters = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.ou(ou).get(filters));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Delete("/ou/:ou", func(c *fiber.Ctx) error {
		// Remove an OU
		// DELETE /ou/:ou
		// Removes the OU from Active Directory.
		// URL Parameters
		// ou		The OU name

		// params: {ou: params.ou},

		// parameters:
		// 	ou: params.ou

		// Javascript
		// const ou = req.params.ou;
		// let [error, response] = await wrapAsync(ad.ou(ou).remove());
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/ou/:ou/exists", func(c *fiber.Ctx) error {
		// OU exists
		// GET /ou/:ou/exists
		// Returns whether a OU exists or not
		// URL Parameters
		// ou		The OU name

		// params: {ou: params.ou},

		// parameters:
		// 	ou: params.ou

		// var attributes []string
		// filter := fmt.Sprintf("(&(objectClass=organizationalUnit)(name=%s))", c.Params("ou"))
		// d := map[string]bool{"data": false}

		// sr := svr.Search(filter, attributes)
		// if len(sr) != 0 {
		// 	d["data"] = true
		// }

		// c.JSON(d)

		// Javascript
		// const ou = req.params.ou;
		// let [error, response] = await wrapAsync(ad.ou(ou).exists());
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})
}
