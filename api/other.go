package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ppreeper/addictgo/ldap"
)

// @Summary Get all other objects
// @Description Pulls all non-user/group Active Directory objects
// @Tags other
// @Produce json
// @Success 200 "OK"
// @Router /other [get]
func OtherGet(c *fiber.Ctx) error {
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

	args := new(ldap.LDAPArgs)
	if err := c.QueryParser(args); err != nil {
		return fmt.Errorf("invalid query args")
	}
	// filter := "(&(!(objectClass=user))(!(objectClass=group)))"
	d := make(map[string]interface{})
	// sr := ldap.Search(filter, args)
	// d["other"] = sr["data"]
	return c.JSON(d)
}

// @Summary Get all objects
// @Description Pulls all Active Directory objects
// @Tags other
// @Accept plain
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /all [get]
func AllGet(c *fiber.Ctx) error {
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
}

// @Summary Search Active Directory
// @Description Does a raw Active Directory search
// @Tags other
// @Accept plain
// @Param filter path string true "Search filter"
// @Produce json
// @Success 200 "OK"
// @Router /find/{filter} [get]
func FindFilterGet(c *fiber.Ctx) error {
	// Search Active Directory
	// GET /find/:filter
	// Does a raw Active Directory search

	// URL Parameters
	// filter	Search filter, such as CN=da*
	filter := c.Params("filter")
	// log.Println(filter)

	// qp := new(ldap.LDAPArgs)
	// if err := c.QueryParser(qp); err != nil {
	// 	return err
	// }
	// log.Println("fields:", qp.Fields)
	// log.Println("eq:", qp.EQ)
	// log.Println("ne:", qp.NE)
	// log.Println("lt:", qp.LT)
	// log.Println("gt:", qp.GT)
	// log.Println("gte:", qp.GTE)
	// log.Println("lte:", qp.LTE)
	// log.Println("like:", qp.Like)
	// log.Println("sort:", qp.Sort)
	// log.Println("order:", qp.Order)
	// log.Println("page:", qp.Page)
	// log.Println("limit:", qp.Limit)
	// log.Println("start:", qp.Start)
	// log.Println("end:", qp.End)
	// log.Println("q:", qp.Q)

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

	return c.JSON(map[string]interface{}{"data": filter})
}

// @Summary Get API status
// @Description Gives the uptime and status of the API
// @Tags other
// @Produce json
// @Success 200 "OK"
// @Router /status [get]
func StatusGet(c *fiber.Ctx) error {
	d := make(map[string]interface{})
	d["online"] = true
	// elapsed := time.Since(startTime)
	// d["uptime"] = fmt.Sprintf("%0.fH%0.fm%0.1fs", elapsed.Hours(), math.Mod(elapsed.Minutes(), 60.0), math.Mod(elapsed.Seconds(), 60.0))
	return c.JSON(d)
}

// func StackGet(c *fiber.Ctx) error {
// 	data, _ := json.MarshalIndent(app.Stack(), "", " ")
// 	return c.Send(data)
// }
