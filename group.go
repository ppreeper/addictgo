package addictgo

import (
	"github.com/gofiber/fiber/v2"
)

func (a *API) Group(app *fiber.App) {
	app.Get("/group", func(c *fiber.Ctx) error {
		// Get all groups
		// GET /group
		// Pulls all groups in Active Directory, with filters.
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// verb: 'GET',
		// route: "/group",
		// title: "Get all groups",
		// description: "Pulls all groups in Active Directory, with filters.",
		// queries: filterQuery

		// /group:
		// 	get:
		// 		summary: "Get all groups"
		// 		description: "Pulls all groups in Active Directory, with filters."
		// 		queries: filterQuery

		// var attributes []string
		// filter := fmt.Sprintf("(objectClass=group)")

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
		// const config = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.group().get(config));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Post("/group", func(c *fiber.Ctx) error {
		// Add a group
		// POST /group
		// Adds a new group to Active Directory
		// Request Body
		// name			Name of the group as displayed
		// Optional
		// description	Descripton of the group
		// location		Relative AD Position separated by /

		// verb: 'POST',
		// route: "/group",
		// title: "Add a group",
		// description: "Adds a new group to Active Directory.",
		// queries: {
		// 		"name": {
		// 				description: 'Name of the Group as displayed.',
		// 				optional: false
		// 		},
		// 		"description": "Description of the Security Group.",
		// 		"location": "Relative AD position, separated by /.",
		// },

		// /group:
		// 	post:
		// 		summary: "Add a group"
		// 		description: "Adds a new group to Active Directory."
		// 		queries:
		// 		name:
		// 		description: "Name of the Group as displayed."
		// 		optional: false
		// 		description: "Description of the Security Group."
		// 		location: "Relative AD position, separated by /."

		// Javascript
		// fmt.Println(c.Route().Params)
		// let [error, response] = await wrapAsync(ad.group().add(req.body));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/group/:group", func(c *fiber.Ctx) error {
		// Get a single group
		// GET /group/:group
		// Puls a single group
		// URL Parameters
		// group	The group name
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// verb: 'GET',
		// route: "/group/:group",
		// title: "Get a single group",
		// description: "Pulls a single group.",
		// params: {group: params.group},
		// queries: filterQuery

		// /group/{group}:
		// 	get:
		// 		summary: "Get a single group"
		// 		description: "Pulls a single group."
		// 		parameters:
		// 		group: params.group
		// 		queries: filterQuery

		// var attributes []string
		// filter := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%s))", c.Params("group"))

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
		// const group = req.params.group;
		// const config = api.parseQuery(req.query);
		// let [error, response] = await wrapAsync(ad.group(group).get(config));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Delete("/group/:group", func(c *fiber.Ctx) error {
		// Remove a group
		// DELETE /group/:group
		// Removes a group from Active Directory
		// URL Parameters
		// group	The group name

		// verb: 'DELETE',
		// route: "/group/:group",
		// title: "Remove a group",
		// description: "Removes the group from Active Directory.",
		// params: {group: params.group},

		// /group/{group}:
		// 	delete:
		// 		summary: "Remove a group"
		// 		description: "Removes the group from Active Directory."
		// 		parameters:
		// 		group: params.group

		// Javascript
		// const group = req.params.group;
		// let [error, response] = await wrapAsync(ad.group(group).remove());
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/group/:group/exists", func(c *fiber.Ctx) error {
		// Group exists
		// GET /group/:group/exists
		// Returns whether a group exists or not
		// URL Parameters
		// group	The group name

		// verb: 'GET',
		// route: "/group/:group/exists",
		// title: "Group exists",
		// description: "Returns whether a group exists or not.",
		// params: {group: params.group},

		// /group/{group}/exists:
		// 	get:
		// 		summary: "Group exists"
		// 		description: "Returns whether a group exists or not."
		// 		parameters:
		// 		group: params.group

		// var attributes []string
		// filter := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%s))", c.Params("group"))
		// d := map[string]bool{"data": false}

		// sr := svr.Search(filter, attributes)
		// if len(sr) != 0 {
		// 	d["data"] = true
		// }

		// c.JSON(d)

		// Javascript
		// const group = req.params.group;
		// let [error, response] = await wrapAsync(ad.group(group).exists());
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Post("/group/:group/user/:user", func(c *fiber.Ctx) error {
		// Add user to group
		// POST /group/:group/user/:user
		// Adds a user to a group
		// URL Parameters
		// group	The group name
		// user		The user logon name

		// verb: 'POST',
		// route: "/group/:group/user/:user",
		// title: "Add user to group",
		// description: "Adds a user to a group.",
		// params: {group: params.group, user: params.user},

		// /group/{group}/user/{user}:
		// 	post:
		// 		summary: "Add user to group"
		// 		description: "Adds a user to a group."
		// 		parameters:
		// 		group: params.group
		// 		user: params.user

		// Javascript
		// const group = req.params.group;
		// const user = req.params.user;
		// let [error, response] = await wrapAsync(ad.user(user).addToGroup(group));
		// response = (!error) ? {success: true} : response;
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Delete("/group/:group/user/:user", func(c *fiber.Ctx) error {
		// Remove user from group
		// DELETE /group/:group/user/:user
		// Removes a user from a group
		// URL Parameters
		// group	The group name
		// user		The user logon name

		// verb: 'DELETE',
		// route: "/group/:group/user/:user",
		// title: "Remove user from group",
		// description: "Removes a user from a group.",
		// params: {group: params.group, user: params.user},

		// /group/{group}/user/{user}:
		// 	delete:
		// 		summary: "Remove user from group"
		// 		description: "Removes a user from a group."
		// 		parameters:
		// 		group: params.group
		// 		user: params.user

		// Javascript
		// const group = req.params.group;
		// const user = req.params.user;
		// let [error, response] = await wrapAsync(ad.user(user).removeFromGroup(group));
		// response = (!error) ? {success: true} : response;
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})
}
