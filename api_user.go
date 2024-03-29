package main

import "github.com/gofiber/fiber/v2"

// @Summary Get all users
// @Description Pulls all users in Active Directory, with filters
// @Tags users
// @Accept plain
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /user [get]
func UsersAllGet(c *fiber.Ctx) error {
	// filter := fmt.Sprintf("(&(objectClass=user)(!objectClass=computer))")
	// var attributes []string

	// fields := strings.Split(c.Query("_fields"), ",")
	// fmt.Println("fields: ", fields)
	// if len(c.Query("_fields")) > 0 {
	// 	for _, a := range fields {
	// 		attributes = append(attributes, a)
	// 	}
	// }
	// q := strings.Split(c.Query("_q"), ",")
	// fmt.Println("q: ", q)

	// q := c.Query("_q")
	// start := c.Query("_start")
	// end := c.Query("_end")
	// fmt.Println(c.Params("user"), len(fields), fields, attributes)

	// sr := svr.Search(filter, attributes)
	// fmt.Printf("%v \n", sr)

	// c.JSON(sr)
	// Javascript
	// const filter = api.parseQuery(req.query);
	// let [error, response] = await wrapAsync(ad.user().get(filter));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Add a user
// @Description Adds a new user to Active Directory
// @Tags users
// @Accept plain
// @Param userName body string true "User name"
// @Param pass body string false "Password to log in"
// @Param commonName body string false "Equivalent to the user's full name"
// @Param firstName body string false "First name"
// @Param lastName body string false "Last name"
// @Param email body string false "Email address"
// @Param title body string false "Job title"
// @Param location body string false "Relative AD folder position"
// @Produce json
// @Success 201 "Created"
// @Router /user [post]
func UserPost(c *fiber.Ctx) error {
	// TODO convert js code to go for user create
	// Javascript
	// let [error, response] = await wrapAsync(ad.user().add(req.body));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Get a user
// @Description Pulls a single user
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user} [get]
func UserUserGet(c *fiber.Ctx) error {
	// var attributes []string
	// filter := fmt.Sprintf("(&(objectClass=user)(!objectClass=computer)(sAMAccountName=%s))", c.Params("user"))

	// fields := strings.Split(c.Query("_fields"), ",")
	// if len(c.Query("_fields")) > 0 {
	// 	for _, a := range fields {
	// 		attributes = append(attributes, a)
	// 	}
	// }

	// q := c.Query("_q")
	// start := c.Query("_start")
	// end := c.Query("_end")
	// fmt.Println(c.Params("user"), len(fields), fields, attributes)

	// sr := svr.Search(filter, attributes)
	// c.JSON(sr)

	// Javascript
	// const user = req.params.user;
	// const config = api.parseQuery(req.query);
	// let [error, response] = await wrapAsync(ad.user(user).get(config));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Remove a user
// @Description Removes a user from Active Directory
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user} [delete]
func UserUserDelete(c *fiber.Ctx) error {
	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).remove());
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Update a user
// @Description Updates properties of a user
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param password body string false "Password of user"
// @Param firstName body string false "First name of user"
// @Param lastName body string false "Last name of user"
// @Param commonName body string false "Full name of user"
// @Param email body string false "Email address of user"
// @Param title body string false "Job title of user"
// @Param enabled body string false "Whether the account is enabled"
// @Param passwordExpires body string false "Whether the password should expire"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user} [put]
func UserUserPut(c *fiber.Ctx) error {
	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).remove());
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary User exists
// @Description Returns whether a user exists or not
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/exists [get]
func UserUserExistsGet(c *fiber.Ctx) error {
	// /user/{user}/exists:
	// 	get:
	// 		summary: "User exists"
	// 		description: "Returns whether a user exists or not."
	// 		parameters:
	// 		user: params.user

	// var attributes []string
	// filter := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", c.Params("user"))
	// d := map[string]bool{"data": false}

	// sr := svr.Search(filter, attributes)
	// if len(sr) != 0 {
	// 	d["data"] = true
	// }

	// c.JSON(d)

	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).exists());
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary User is a member of group
// @Description Returns whether a user is a member of a group or not
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/member-of/{group} [get]
func UserUserMemberofGroupGet(c *fiber.Ctx) error {
	// /user/{user}/member-of/{group}:
	// 	get:
	// 		summary: "User is a member of group"
	// 		description: "Returns whether a user is a member of a group or not."
	// 		parameters:
	// 		user: params.user
	// 		group: params.group

	// var attributes []string
	// filter := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%s))", c.Params("group"))
	// d := map[string]bool{"data": false}

	// gr := svr.Search(filter, attributes)
	// if len(gr) != 0 {
	// 	filter = fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s)(memberOf=%s))", c.Params("user"), gr["distinguishedName"])
	// 	sr := svr.Search(filter, attributes)
	// 	if len(sr) != 0 {
	// 		d["data"] = true
	// 	}
	// }

	// c.JSON(d)

	// Javascript
	// const user = req.params.user;
	// const group = req.params.group;
	// let [error, response] = await wrapAsync(ad.user(user).isMemberOf(group));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Authenticate a user
// @Description Runs a username and password against AD
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param password body string true "The users password"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/authenticate [post]
func UserUserAuthenticatePost(c *fiber.Ctx) error {
	// /user/{user}/authenticate:
	// 	post:
	// 		summary: "Authenticate a user"
	// 		description: "Runs a username and password against AD."
	// 		parameters:
	// 		user: params.user
	// 		queries:
	// 		password:
	// 		description: "The user's password."
	// 		optional: false

	// d := map[string]bool{"data": false}

	// r := svr.Authenticate(c.Params("user"), c.FormValue("password"))
	// c.JSON(r)

	// Javascript
	// const user = req.params.user;
	// const pass = req.body.pass || req.body.password;
	// let [error, response] = await wrapAsync(ad.user(user).authenticate(pass));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Change a password
// @Description Changes the users password
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param password body string true "The users password"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/password [put]
func UserUserPasswordPut(c *fiber.Ctx) error {
	// /user/{user}/password:
	// 	put:
	// 		summary: "Change a password"
	// 		description: "Changes the user's password"
	// 		parameters:
	// 		user: params.user
	// 		queries:
	// 		password:
	// 		description: "The user's password."
	// 		optional: false

	// Javascript
	// const user = req.params.user;
	// const pass = req.body.pass || req.body.password;
	// let [error, response] = await wrapAsync(ad.user(user).password(pass));
	// response = (!error) ? {success: true} : response;
	// error = (error) ? Object.assign({success: false}, error) : error;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Set password never expires
// @Description Sets the password to never expire
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/password-never-expires [put]
func UserUserPasswordneverexpiresPut(c *fiber.Ctx) error {
	// /user/{user}/password-never-expires:
	// 	put:
	// 		summary: "Set password never expires"
	// 		description: "Sets the password to never expire."
	// 		parameters:
	// 		user: params.user
	// 		queries: {}

	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).passwordNeverExpires());
	// response = (!error) ? {success: true} : response;
	// error = (error) ? Object.assign({success: false}, error) : error;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Set password expires
// @Description Set the password to eventually expire
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/password-expires [put]
func UserUserPasswordExpiresPut(c *fiber.Ctx) error {
	// /user/{user}/password-expires:
	// 	put:
	// 		summary: "Set password expires"
	// 		description: "Sets the password to eventually expire."
	// 		parameters:
	// 		user: params.user
	// 		queries: {}

	// Javascript
	// const user = req.params.user;
	// let [error, data] = await wrapAsync(ad.user(user).passwordExpires());
	// let response = (!error) ? {success: true} : undefined;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Enable a user
// @Description Enables the users account
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/enable [put]
func UserUserEnablePut(c *fiber.Ctx) error {
	// /user/{user}/enable:
	// 	put:
	// 		summary: "Enable a user"
	// 		description: "Enables the user's account."
	// 		parameters:
	// 		user: params.user
	// 		queries: {}

	// Javascript
	// const user = req.params.user;
	// let [error, data] = await wrapAsync(ad.user(user).enable());
	// let response = (!error) ? {success: true} : undefined;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Disable a user
// @Description Disables the users account
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/disable [put]
func UserUserDisablePut(c *fiber.Ctx) error {
	// /user/{user}/disable:
	// 	put:
	// 		summary: "Disable a user"
	// 		description: "Disables the user's account."
	// 		parameters:
	// 		user: params.user
	// 		queries: {}

	// Javascript
	// const user = req.params.user;
	// let [error, data] = await wrapAsync(ad.user(user).disable());
	// let response = (!error) ? {success: true} : undefined;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Move a user
// @Description Moves a user to a new location relative to the directory root
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param location body string true "The new position, separated by /"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/move [put]
func UserUserMovePut(c *fiber.Ctx) error {
	// /user/{user}/move:
	// 	put:
	// 		summary: "Move a user"
	// 		description: "Moves a user to a new location relative to the directory root."
	// 		parameters:
	// 		user: params.user
	// 		queries:
	// 		location:
	// 		description: "The new position, separated by '/'."

	// Javascript
	// const user = req.params.user;
	// const location = req.body.location;
	// let [error, response] = await wrapAsync(ad.user(user).move(location));
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}

// @Summary Unlock a user
// @Description Unlocks a user
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /user/{user}/unlock [put]
func UserUserUnlockPut(c *fiber.Ctx) error {
	// /user/{user}/unlock:
	// 	put:
	// 		summary: "Unlock a user"
	// 		description: "Unlocks a user."
	// 		parameters:
	// 		user: params.user

	// Javascript
	// const user = req.params.user;
	// let [error, data] = await wrapAsync(ad.unlockUser(user));
	// let response = (!error) ? {success: true} : undefined;
	// respond(res, error, response);

	return c.JSON(map[string]interface{}{"data": "OU"})
}
