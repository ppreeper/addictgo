package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ppreeper/addictgo/ldap"
)

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
func AllGet(c echo.Context) error {
	d := make(map[string]any)
	var args ldap.LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := "(objectClass=user)"
	sr := lconn.Search(filter, args)
	d["users"] = sr["data"]

	filter = "(objectClass=group)"
	sr = lconn.Search(filter, args)
	d["groups"] = sr["data"]

	filter = "(&(!(objectClass=user))(!(objectClass=group)))"
	sr = lconn.Search(filter, args)
	d["other"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Search Active Directory
// @Description Does a raw Active Directory search
// @Tags other
// @Accept plain
// @Param filter path string true "Search filter"
// @Produce json
// @Success 200 "OK"
// @Router /find/{filter} [get]
func FindFilterGet(c echo.Context) error {
	d := make(map[string]any)
	var args ldap.LDAPArgs

	// filter	Search filter, such as CN=da*
	filter := c.QueryParam("filter")
	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	sr := lconn.Search(filter, args)
	d["data"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Get all other objects
// @Description Pulls all non-user/group Active Directory objects
// @Tags other
// @Param _fields query []string false "Comma-delimited field names to return"
// @Param _q query []string false "Searches all fields for given string"
// @Param _start query int false "Result Index to start from"
// @Param _end query int false "Result Index to end to"
// @Produce json
// @Success 200 "OK"
// @Router /other [get]
func OtherGet(c echo.Context) error {
	d := make(map[string]any)
	var args ldap.LDAPArgs

	args.Fields = c.QueryParam("_fields")
	args.Q = c.QueryParam("_q")
	args.Start, _ = strconv.Atoi(c.QueryParam("_start"))
	args.End, _ = strconv.Atoi(c.QueryParam("_end"))

	filter := "(&(!(objectClass=user))(!(objectClass=group)))"
	sr := lconn.Search(filter, args)
	d["other"] = sr["data"]

	return c.JSON(http.StatusOK, d)
}

// @Summary Get API status
// @Description Gives the uptime and status of the API.
// @Tags other
// @Produce json
// @Success 200 "OK"
// @Router /status [get]
func StatusGet(c echo.Context) error {
	d := make(map[string]interface{})
	d["online"] = true
	elapsed := time.Since(startTime)
	// d["uptime"] = fmt.Sprintf("%0.fH%2.fm%0.1fs", elapsed.Hours(), math.Mod(elapsed.Minutes(), 60.0), math.Mod(elapsed.Seconds(), 60.0))
	d["uptime"] = fmtDuration(elapsed)
	return c.JSON(http.StatusOK, d)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
