package addictgo

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	app.Get("/all", func(c *fiber.Ctx) error {
		// Get all objects
		// GET /all
		// Pulls all Active Directory objects
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// var attributes []string
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

		// filter := fmt.Sprintf("(objectClass=user)")
		// sr := svr.Search(filter, attributes)
		// d["users"] = sr["data"]

		// filter = fmt.Sprintf("(objectClass=group)")
		// sr = svr.Search(filter, attributes)
		// d["groups"] = sr["data"]

		// filter = fmt.Sprintf("(&(!(objectClass=user))(!(objectClass=group)))")
		// sr = svr.Search(filter, attributes)
		// d["other"] = sr["data"]

		// c.JSON(d)

		// Javascript
		// const config = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.all().get(config));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/find/:filter", func(c *fiber.Ctx) error {
		// Search Active Directory
		// GET /find/:filter
		// Does a raw Active Directory search

		// URL Parameters
		// filter	Search filter, such as CN=da*
		filter := c.Params("filter")
		log.Println(filter)

		qp := new(LDAPArgs)
		if err := c.QueryParser(qp); err != nil {
			return err
		}
		log.Println("fields:", qp.Fields)
		log.Println("eq:", qp.EQ)
		log.Println("ne:", qp.NE)
		log.Println("lt:", qp.LT)
		log.Println("gt:", qp.GT)
		log.Println("gte:", qp.GTE)
		log.Println("lte:", qp.LTE)
		log.Println("like:", qp.Like)
		log.Println("sort:", qp.Sort)
		log.Println("order:", qp.Order)
		log.Println("page:", qp.Page)
		log.Println("limit:", qp.Limit)
		log.Println("start:", qp.Start)
		log.Println("end:", qp.End)
		log.Println("q:", qp.Q)

		// 	var attributes []string
		// 	// d := make(map[string]interface{})
		// 	filter := fmt.Sprintf("(*)")
		// 	sr := svr.Search(filter, attributes)
		// 	c.JSON(sr)
		// Javascript
		// const filter = req.params.filter;
		// const config = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.find(filter, config));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "/find/:filter"})
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		// Get API status
		// GET /status
		// Gives the uptime and status of the API

		// operationID: apiStatus
		// responses:
		// 	"200":
		// 		description: OK

		// d := make(map[string]interface{})
		// d["online"] = true
		// elapsed := time.Since(startTime)
		// d["uptime"] = fmt.Sprintf("%0.fH%0.fm%0.1fs", elapsed.Hours(), math.Mod(elapsed.Minutes(), 60.0), math.Mod(elapsed.Seconds(), 60.0))
		// c.JSON(d)

		// Javascript
		// let uptime = new Date() - start;
		// res.send({online: true, uptime});

		return c.JSON(map[string]interface{}{"data": "status"})
	})

	// Get Prometheus Stats
	// GET /metrics
	// Gives the prometheus stats
	// Prometheus
	app.Get("/metrics", monitor.New())

	// Get Monitor Stats
	// GET /monitor
	app.Get("/monitor", monitor.New())
}
