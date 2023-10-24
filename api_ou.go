package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppreeper/addictgo/ldap"
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
// @Router /ou [get]
func OUGet(c echo.Context) error {
	qp := new(ldap.LDAPArgs)
	log.Println(qp)
	// if err := c.QueryParam(qp); err != nil {
	// 	return err
	// }

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

	return c.JSON(http.StatusOK, map[string]interface{}{"data": "sr"})
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
// @Router /ou [post]
func OUPost(c echo.Context) error {
	// queries:
	// 	name:
	// 	description: "Name of the OU as displayed."
	// 	optional: false
	// 	description: "Description of the OU."
	// 	location: "Relative AD position, separated by /."

	// Javascript
	// let [error, response] = await wrapAsync(ad.ou().add(req.body));
	// respond(res, error, response);

	return c.JSON(http.StatusOK, map[string]interface{}{"data": "OU"})
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
// @Router /ou/{ou} [get]
func OUOUGet(c echo.Context) error {
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

	return c.JSON(http.StatusOK, map[string]interface{}{"data": "OU"})
}

// @Summary Remove an OU
// @Description Removes the OU from Active Directory
// @Tags ou
// @Accept plain
// @Param ou path string true "The OU name"
// @Produce json
// @Success 200 "OK"
// @Router /ou/{ou} [delete]
func OUOUDelete(c echo.Context) error {
	// Javascript
	// const ou = req.params.ou;
	// let [error, response] = await wrapAsync(ad.ou(ou).remove());
	// respond(res, error, response);

	return c.JSON(http.StatusOK, map[string]interface{}{"data": "OU"})
}

// @Summary OU exists
// @Description Returns whether a OU exists or not
// @Tags ou
// @Accept plain
// @Param ou path string true "The OU name"
// @Produce json
// @Success 200 "OK"
// @Router /ou/{ou}/exists [get]
func OUOUExists(c echo.Context) error {
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

	return c.JSON(http.StatusOK, map[string]interface{}{"data": "OU"})
}
