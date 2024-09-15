package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
// @Router /api/user [get]
func (lconn *LDAP) UsersAllGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := "(&(objectClass=user)(!objectClass=computer))"
	sr := lconn.Search(filter, args)
	d["users"] = sr["data"]

	return c.JSON(http.StatusOK, d)
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
// @Router /api/user [post]
func (lconn *LDAP) UserPost(c echo.Context) error {
	d := make(map[string]any)
	// TODO convert js code to go for user create
	// Javascript
	// let [error, response] = await wrapAsync(ad.user().add(req.body));
	// respond(res, error, response);

	return c.JSON(http.StatusCreated, d)
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
// @Router /api/user/{user} [get]
func (lconn *LDAP) UserUserGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := fmt.Sprintf("(&(objectClass=user)(!objectClass=computer)(sAMAccountName=%s))", c.Param("user"))
	sr := lconn.Search(filter, args)
	d["users"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Remove a user
// @Description Removes a user from Active Directory
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user} [delete]
func (lconn *LDAP) UserUserDelete(c echo.Context) error {
	d := make(map[string]any)
	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).remove());
	// respond(res, error, response);

	return c.JSON(http.StatusOK, d)
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
// @Router /api/user/{user} [put]
func (lconn *LDAP) UserUserPut(c echo.Context) error {
	d := make(map[string]any)
	// Javascript
	// const user = req.params.user;
	// let [error, response] = await wrapAsync(ad.user(user).remove());
	// respond(res, error, response);

	return c.JSON(http.StatusOK, d)
}

// @Summary User exists
// @Description Returns whether a user exists or not
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/exists [get]
func (lconn *LDAP) UserUserExistsGet(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary User is a member of group
// @Description Returns whether a user is a member of a group or not
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/member-of/{group} [get]
func (lconn *LDAP) UserUserMemberofGroupGet(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Authenticate a user
// @Description Runs a username and password against AD
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param password body string true "The users password"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/authenticate [post]
func (lconn *LDAP) UserUserAuthenticatePost(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Change a password
// @Description Changes the users password
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param password body string true "The users password"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/password [put]
func (lconn *LDAP) UserUserPasswordPut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Set password never expires
// @Description Sets the password to never expire
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/password-never-expires [put]
func (lconn *LDAP) UserUserPasswordneverexpiresPut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Set password expires
// @Description Set the password to eventually expire
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/password-expires [put]
func (lconn *LDAP) UserUserPasswordExpiresPut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Enable a user
// @Description Enables the users account
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/enable [put]
func (lconn *LDAP) UserUserEnablePut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Disable a user
// @Description Disables the users account
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/disable [put]
func (lconn *LDAP) UserUserDisablePut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Move a user
// @Description Moves a user to a new location relative to the directory root
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Param location body string true "The new position, separated by /"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/move [put]
func (lconn *LDAP) UserUserMovePut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Unlock a user
// @Description Unlocks a user
// @Tags users
// @Accept plain
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/user/{user}/unlock [put]
func (lconn *LDAP) UserUserUnlockPut(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}
