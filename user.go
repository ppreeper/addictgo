package addictgo

import "github.com/gofiber/fiber/v2"

func (a *API) User(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		// Get all users
		// GET /user
		// Pulls all users in Active Directory, with filters
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// verb: 'GET',
		// route: "/user",
		// title: "Get all users",
		// description: "Pulls all users in Active Directory, with filters.",
		// queries: filterQuery

		// /users:
		// 	get:
		// 	summary: Get all users
		// 	description: Pulls all users in Active Directory, with filters.
		// 	responses:
		// 		"200":
		// 		description: A JSON array of user names
		// 		content:
		// 		application/json:
		// 			schema:
		// 			type: array
		// 				items:
		// 				type: string

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
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		// Add a user
		// POST /user
		// Adds a new user to Active Directory
		// Request Body
		// name			The user logon name
		// Optional
		// description	Descripton of the user
		// location		Relative AD Position separated by /
		// POST /user

		// verb: 'POST',
		// route: "/user",
		// title: "Add a user",
		// description: "Adds a new user to Active Directory.",
		// queries: {
		// 		"commonName": {
		// 				description: 'Equivalent to the user\'s full name.',
		// 				optional: false
		// 		},
		// 		"userName": {
		// 				description: 'User name.',
		// 				optional: false
		// 		},
		// 		"pass": {
		// 				description: 'Password to log in.',
		// 				optional: false
		// 		},
		// 		"firstName": "First name",
		// 		"lastName": 'Last name.',
		// 		"email": 'Email address.',
		// 		"title": 'Job title.',
		// 		"location": 'Relative AD folder position.',
		// }

		// /users:
		// 	post:
		// 	summary: "Add a user"
		// 	description: "Adds a new user to Active Directory."
		// 	requestBody:
		// 		content:
		// 		application/json:
		// 			schema:
		// 			type: object
		// 				properties:
		// 				commonName:
		// 					type: string
		// 					description: "Equivalent to the user's full name."
		// 				userName:
		// 					type: string
		// 					description: "User name."
		// 				pass:
		// 					type: string
		// 					description: "Password to log in."
		// 				firstName:
		// 					type: string
		// 					description: "First name"
		// 				lastName:
		// 					type: string
		// 					description: "Last name."
		// 				email:
		// 					type: string
		// 					description: "Email address."
		// 				title:
		// 					type: string
		// 					description: "Job title."
		// 				location:
		// 					type: string
		// 					description: "Relative AD folder position."
		// 	responses:
		// 		"201":
		// 		description: Created

		// TODO convert js code to go for user create
		// Javascript
		// let [error, response] = await wrapAsync(ad.user().add(req.body));
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/user/:user", func(c *fiber.Ctx) error {
		// Get a user
		// GET /user/:user
		// Puls a single user
		// URL Parameters
		// user		The user logon name
		// Query Parameters
		// _fields	Comma-delimited field names to return
		// _q		Searches all fields for given string
		// _start	Result Index to start from
		// _end		Result Index to end to

		// verb: 'GET',
		// route: "/user/:user",
		// title: "Get a single user",
		// description: "Pulls a single user.",
		// params: {user: params.user},
		// queries: filterQuery

		// /user/{user}:
		// 	get:
		// 		summary: "Get a single user"
		// 		description: "Pulls a single user."
		// 		parameters:
		// 		- in: path
		// 		name: user
		// 		description: username
		// 		schema:
		// 		type: string
		// 		required: true
		// 		queries: filterQuery
		// 		responses:
		// 		"200":
		// 		description: Successful operation
		// 		content:
		// 		application/json:
		// 		schema:
		// 		type: object

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
	})

	app.Delete("/user/:user", func(c *fiber.Ctx) error {
		// Remove a user
		// DELETE /user/:user
		// Removes a user from Active Directory
		// URL Parameters
		// user		The user logon name

		// verb: 'DELETE',
		// 	route: "/user/:user",
		// 	title: "Remove a user",
		// 	description: "Removes a user from Active Directory.",
		// 	params: {user: params.user}

		// /user/{user}:
		// 	put:
		// 		summary: "Update a user"
		// 		description: "Updates properties of a user."
		// 		parameters:
		// 		user: params.user
		// 		queries:
		// 		firstName: "First name of user."
		// 		lastName: "Last name of user."
		// 		commonName: "Full name of user."
		// 		email: "Email address of user."
		// 		title: "Job title of user."
		// 		password: "Password of user."
		// 		userName: "Username of user."
		// 		enabled: "Whether the account is enabled."
		// 		passwordExpires: "Whether the password should expire."
		// 		delete:
		// 		summary: "Remove a user"
		// 		description: "Removes a user from Active Directory."
		// 		parameters:
		// 		user: params.user

		// Javascript
		// const user = req.params.user;
		// let [error, response] = await wrapAsync(ad.user(user).remove());
		// respond(res, error, response);

		return c.JSON(map[string]interface{}{"data": "OU"})
	})

	app.Get("/user/:user/exists", func(c *fiber.Ctx) error {
		// User exists
		// GET /user/:user/exists
		// Returns whether a user exists or not
		// URL Parameters
		// user		The user logon name

		// verb: 'GET',
		// route: "/user/:user/exists",
		// title: "User exists",
		// description: "Returns whether a user exists or not.",
		// params: {user: params.user},

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
	})

	app.Get("/user/:user/member-of/:group", func(c *fiber.Ctx) error {
		// User is a member of group
		// GET /user/:user/member-of/:group
		// Returns whether a user is a member of a group or not
		// URL Parameters
		// user		The user logon name
		// group	The group name

		// verb: 'GET',
		// route: "/user/:user/member-of/:group",
		// title: "User is a member of group",
		// description: "Returns whether a user is a member of a group or not.",
		// params: {user: params.user, group: params.group},

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
	})

	app.Post("/user/:user/authenticate", func(c *fiber.Ctx) error {
		// Authenticate a user
		// POST /user/:user/authenticate
		// Runs a username and password against AD
		// URL Parameters
		// user		The user logon name
		// Request Body
		// password	The users password

		// verb: 'POST',
		// route: "/user/:user/authenticate",
		// title: "Authenticate a user",
		// description: "Runs a username and password against AD.",
		// params: {user: params.user},
		// queries: {
		// 		password: {
		// 				description: 'The user\'s password.',
		// 				optional: false
		// 		}
		// }

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
	})

	app.Put("/user/:user/password", func(c *fiber.Ctx) error {
		// Change a password
		// PUT /user/:user/password
		// Changes the users password
		// URL Parameters
		// user		The user logon name
		// Request Body
		// password	The users password

		// verb: 'PUT',
		// route: "/user/:user/password",
		// title: "Change a password",
		// description: "Changes the user's password",
		// params: {user: params.user},
		// queries: {
		// 		password: {
		// 				description: 'The user\'s password.',
		// 				optional: false
		// 		}
		// }

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
	})

	app.Put("/user/:user/password-never-expires", func(c *fiber.Ctx) error {
		// Set password never expires
		// PUT /user/:user/password-never-expires
		// Sets the password to never expire
		// URL Parameters
		// user		The user logon name

		// verb: 'PUT',
		// route: "/user/:user/password-never-expires",
		// title: "Set password never expires",
		// description: "Sets the password to never expire.",
		// params: {user: params.user},
		// queries: {}

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
	})

	app.Put("/user/:user/password-expires", func(c *fiber.Ctx) error {
		// Set password expires
		// PUT /user/:user/password-expires
		// Set the password to eventually expire
		// URL Parameters
		// user		The user logon name

		// verb: 'PUT',
		// route: "/user/:user/password-expires",
		// title: "Set password expires",
		// description: "Sets the password to eventually expire.",
		// params: {user: params.user},
		// queries: {}

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
	})

	app.Put("/user/:user/enable", func(c *fiber.Ctx) error {
		// Enable a user
		// PUT /user/:user/enable
		// Enables a users account
		// URL Parameters
		// user		The user logon name

		// verb: 'PUT',
		// route: "/user/:user/enable",
		// title: "Enable a user",
		// description: "Enables the user's account.",
		// params: {user: params.user},
		// queries: {}

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
	})

	app.Put("/user/:user/disable", func(c *fiber.Ctx) error {
		// Disable a user
		// PUT /user/:user/disable
		// Disables the users account
		// URL Parameters
		// user		The user logon name

		// verb: 'PUT',
		// route: "/user/:user/disable",
		// title: "Disable a user",
		// description: "Disables the user's account.",
		// params: {user: params.user},
		// queries: {}

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
	})

	app.Put("/user/:user/move", func(c *fiber.Ctx) error {
		// Move a user
		// PUT /user/:user/move
		// Moves a user to a new location relative to the directory root
		// URL Parameters
		// user		The user logon name
		// Request Body
		// location	The new position, separated by '/'

		// verb: 'PUT',
		// route: "/user/:user/move",
		// title: "Move a user",
		// description: "Moves a user to a new location relative to the directory root.",
		// params: {user: params.user},
		// queries: {
		// 		location: {
		// 				description: "The new position, separated by '/'.",
		// 				optional: false
		// 		}
		// }

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
	})

	app.Put("/user/:user/unlock", func(c *fiber.Ctx) error {
		// Unlock a user
		// PUT /user/:user/unlock
		// Unlocks a user
		// URL Parameters
		// user		The user logon name

		// verb: 'PUT',
		// 	route: "/user/:user/unlock",
		// 	title: "Unlock a user",
		// 	description: "Unlocks a user.",
		// 	params: {user: params.user}

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
	})
}
