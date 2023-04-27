package api

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all groups
// @Description Pulls all groups in Active Directory, with filters.
// @Tags group
// @Accept plain
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /group [get]
func GroupGet(c *fiber.Ctx) error {
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
}

// @Summary Add a group
// @Description Adds a new group to Active Directory
// @Tags group
// @Accept plain
// @Param name body string true "Name of the group as displayed"
// @Param description body string false "Descripton of the group"
// @Param location body string false "Relative AD Position separated by /"
// @Produce json
// @Success 201 "Created"
// @Router /group [post]
func GroupPost(c *fiber.Ctx) error {
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
}

// @Summary Get a single group
// @Description Pulls a single group
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /group/{group} [get]
func GroupGroupGet(c *fiber.Ctx) error {
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
}

// @Summary Remove a group
// @Description Removes a group from Active Directory
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /group/{group} [delete]
func GroupGroupDelete(c *fiber.Ctx) error {
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
}

// @Summary Group exists
// @Description Returns whether a group exists or not
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /group/{group}/exists [get]
func GroupGroupExistsGet(c *fiber.Ctx) error {
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
}

// @Summary Add user to group
// @Description Adds a user to a group
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /group/{group}/user/{user} [post]
func GroupGroupUserUserPost(c *fiber.Ctx) error {
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
}

// @Summary Remove user from group
// @Description Removes a user from a group
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /group/{group}/user/{user} [delete]
func GroupGroupUserUserDelete(c *fiber.Ctx) error {
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
}
