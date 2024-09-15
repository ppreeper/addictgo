package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
// @Router /api/group [get]
func (lconn *LDAP) GroupGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := "(objectClass=group)"
	sr := lconn.Search(filter, args)
	d["groups"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Add a group
// @Description Adds a new group to Active Directory
// @Tags group
// @Param cn formData string true "Name of the group as displayed"
// @Param description formData string false "Descripton of the group"
// @Param location formData string true "Relative AD Position separated by / (ex Users or Users/Office)"
// @Success 201 "Created"
// @Router /api/group [post]
func (lconn *LDAP) GroupPost(c echo.Context) error {
	d := make(map[string]any)

	cn := c.FormValue("cn")
	description := c.FormValue("description")
	location := c.FormValue("location")
	if cn == "" || location == "" {
		return fmt.Errorf("cn and location are mandatory fields")
	}

	dn := "CN=" + cn
	if location != "" {
		dn += ",CN=" + location
	}
	attrs := Attrs{
		{Type: "objectClass", Vals: []string{"group"}},
		{Type: "CN", Vals: []string{cn}},
		{Type: "description", Vals: []string{description}},
	}
	lconn.AddRecord(dn, attrs)

	return c.JSON(http.StatusCreated, d)
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
// @Router /api/group/{group} [get]
func (lconn *LDAP) GroupGroupGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := fmt.Sprintf("(&(objectClass=group)(sAMAccountName=%s))", c.Param("group"))
	sr := lconn.Search(filter, args)
	d["groups"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Remove a group
// @Description Removes a group from Active Directory
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /api/group/{group} [delete]
func (lconn *LDAP) GroupGroupDelete(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Group exists
// @Description Returns whether a group exists or not
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Produce json
// @Success 200 "OK"
// @Router /api/group/{group}/exists [get]
func (lconn *LDAP) GroupGroupExistsGet(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Add user to group
// @Description Adds a user to a group
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/group/{group}/user/{user} [post]
func (lconn *LDAP) GroupGroupUserUserPost(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}

// @Summary Remove user from group
// @Description Removes a user from a group
// @Tags group
// @Accept plain
// @Param group path string true "The group name"
// @Param user path string true "The user logon name"
// @Produce json
// @Success 200 "OK"
// @Router /api/group/{group}/user/{user} [delete]
func (lconn *LDAP) GroupGroupUserUserDelete(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}
