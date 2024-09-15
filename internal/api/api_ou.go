package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Get all OUs
// @Description Pulls all Organization Units in Active Directory, with filters
// @Tags ou
// @Accept plain
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /api/ou [get]
func (lconn *LDAP) OUGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))
	filter := "(objectClass=organizationalUnit)"
	sr := lconn.Search(filter, args)
	d["other"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Add an OU
// @Description Adds a new OU to Active Directory
// @Tags ou
// @Accept plain
// @Param name body string true "Name of the OU as displayed"
// @Param description body string false "Descripton of the OU"
// @Param location body string false "Relative AD Position separated by /"
// @Produce json
// @Success 201 "Created"
// @Router /api/ou [post]
func (lconn *LDAP) OUPost(c echo.Context) error {
	d := make(map[string]any)
	// queries:
	// 	name:
	// 	description: "Name of the OU as displayed."
	// 	optional: false
	// 	description: "Description of the OU."
	// 	location: "Relative AD position, separated by /."

	// Javascript
	// let [error, response] = await wrapAsync(ad.ou().add(req.body));
	// respond(res, error, response);

	return c.JSON(http.StatusCreated, d)
}

// @Summary Get a single OU
// @Description Pulls a single OU
// @Tags ou
// @Accept plain
// @Param ou path string true "The OU name"
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /api/ou/{ou} [get]
func (lconn *LDAP) OUOUGet(c echo.Context) error {
	d := make(map[string]any)
	var args LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := fmt.Sprintf("(&(objectClass=organizationalUnit)(name=%s))", c.Param("ou"))

	sr := lconn.Search(filter, args)
	d["other"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Remove an OU
// @Description Removes the OU from Active Directory
// @Tags ou
// @Accept plain
// @Param ou path string true "The OU name"
// @Produce json
// @Success 200 "OK"
// @Router /api/ou/{ou} [delete]
func (lconn *LDAP) OUOUDelete(c echo.Context) error {
	d := make(map[string]any)
	// Javascript
	// const ou = req.params.ou;
	// let [error, response] = await wrapAsync(ad.ou(ou).remove());
	// respond(res, error, response);

	return c.JSON(http.StatusOK, d)
}

// @Summary OU exists
// @Description Returns whether a OU exists or not
// @Tags ou
// @Accept plain
// @Param ou path string true "The OU name"
// @Produce json
// @Success 200 "OK"
// @Router /api/ou/{ou}/exists [get]
func (lconn *LDAP) OUOUExists(c echo.Context) error {
	d := make(map[string]any)
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

	return c.JSON(http.StatusOK, d)
}
