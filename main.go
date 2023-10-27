package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/ppreeper/addictgo/docs"

	"github.com/ppreeper/addictgo/ldap"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var startTime time.Time

var lconn ldap.LDAP

// @title AD Dict API
// @version 1.0
// @description This is the swagger spec for AD Dict
// @termsOfService http://preeper.org/addict/
// @contact.name Peter Preeper
// @contact.email ppreeper@gmail.com
// @license.name AGPL-3
// @license.url https://www.gnu.org/licenses/agpl-3.0.html
// @host localhost:3001
// @BasePath /
func main() {
	startTime = time.Now()
	// var l ldap.LDAP
	flag.StringVar(&lconn.URL, "url", LookupEnvOrString("ADDICT_URL", ""), "Address for Active Directory")
	flag.StringVar(&lconn.Username, "u", LookupEnvOrString("ADDICT_USER", ""), "BIND username")
	flag.StringVar(&lconn.Password, "p", LookupEnvOrString("ADDICT_PASS", ""), "BIND password")
	flag.StringVar(&lconn.Port, "port", LookupEnvOrString("ADDICT_PORT", "3001"), "Port to listen on")
	flag.Parse()
	if lconn.URL == "" {
		fmt.Println("no ldap url provided")
		return
	}
	if lconn.Username == "" || lconn.Password == "" {
		fmt.Println("no username or password provided")
		return
	}

	lconn.Scope = ldap.GetScope(lconn.Username)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(echoprometheus.NewMiddleware("addict"))

	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.GET("/status", StatusGet)
	app.GET("/metrics", echoprometheus.NewHandler())

	// ROUTER
	// Router > Other
	app.GET("/all", AllGet)
	app.GET("/find/:filter", FindFilterGet)
	app.GET("/other", OtherGet)

	// Router > OU
	app.GET("/ou", OUGet)
	app.POST("/ou", OUPost)
	app.GET("/ou/:ou", OUOUGet)
	app.DELETE("/ou/:ou", OUOUDelete)
	app.GET("/ou/:ou/exists", OUOUExists)

	// Router > Group
	app.GET("/group", GroupGet)
	app.POST("/group", GroupPost)
	app.GET("/group/:group", GroupGroupGet)
	app.DELETE("/group/:group", GroupGroupDelete)
	app.GET("/group/:group/exists", GroupGroupExistsGet)
	app.POST("/group/:group/user/:user", GroupGroupUserUserPost)
	app.DELETE("/group/:group/user/:user", GroupGroupUserUserDelete)

	// Router > User
	app.GET("/user", UsersAllGet)
	app.POST("/user", UserPost)
	app.GET("/user/:user", UserUserGet)
	app.DELETE("/user/:user", UserUserDelete)
	app.GET("/user/:user/exists", UserUserExistsGet)
	app.GET("/user/:user/member-of/:group", UserUserMemberofGroupGet)
	app.POST("/user/:user/authenticate", UserUserAuthenticatePost)
	app.PUT("/user/:user/password", UserUserPasswordPut)
	app.PUT("/user/:user/password-never-expires", UserUserPasswordneverexpiresPut)
	app.PUT("/user/:user/password-expires", UserUserPasswordExpiresPut)
	app.PUT("/user/:user/enable", UserUserEnablePut)
	app.PUT("/user/:user/disable", UserUserDisablePut)
	app.PUT("/user/:user/move", UserUserMovePut)
	app.PUT("/user/:user/unlock", UserUserUnlockPut)

	app.Logger.Fatal(app.Start(lconn.HostPort()))
	// app.ListenTLS(l.HostPort(), "cert.pem", "key.pem")
}

// LookupEnvOrString provides 12 Factor for string vars
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// LookupEnvOrInt provides 12 Factor for int vars
func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func encodePassword(password string) string {
	passrune := []rune(password)
	for k, v := range passrune {
		passrune[k] = v & 0xFF
	}
	return string(passrune)
}
