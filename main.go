package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	var l ldap.LDAP
	flag.StringVar(&l.URL, "url", LookupEnvOrString("ADDICT_URL", ""), "Address for Active Directory")
	flag.StringVar(&l.Username, "u", LookupEnvOrString("ADDICT_USER", ""), "BIND username")
	flag.StringVar(&l.Password, "p", LookupEnvOrString("ADDICT_PASS", ""), "BIND password")
	flag.StringVar(&l.Port, "port", LookupEnvOrString("ADDICT_PORT", "3001"), "Port to listen on")
	flag.Parse()
	if l.URL == "" {
		fmt.Println("no ldap url provided")
		return
	}
	if l.Username == "" || l.Password == "" {
		fmt.Println("no username or password provided")
		return
	}

	l.Scope = ldap.GetScope(l.Username)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(echoprometheus.NewMiddleware("addict"))

	fmt.Println(encodePassword("p4ssw04d"))

	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.GET("/status", StatusGet)
	app.GET("/metrics", echoprometheus.NewHandler())

	// ROUTER
	// Router > Other
	app.GET("/other", basicAnswer)
	app.GET("/all", AllGet)
	app.GET("/find/:filter", FindFilterGet)
	app.GET("/stack", basicAnswer)

	// Router > OU
	app.GET("/ou", basicAnswer)
	app.POST("/ou", basicAnswer)
	app.GET("/ou/:ou", basicAnswer)
	app.DELETE("/ou/:ou", basicAnswer)
	app.GET("/ou/:ou/exists", basicAnswer)

	// Router > Group
	app.GET("/group", basicAnswer)
	app.POST("/group", basicAnswer)
	app.GET("/group/:group", basicAnswer)
	app.DELETE("/group/:group", basicAnswer)
	app.GET("/group/:group/exists", basicAnswer)
	app.POST("/group/:group/user/:user", basicAnswer)
	app.DELETE("/group/:group/user/:user", basicAnswer)

	// Router > User
	app.GET("/user", basicAnswer)
	app.GET("/user", basicAnswer)
	app.GET("/user/:user", basicAnswer)
	app.DELETE("/user/:user", basicAnswer)
	app.GET("/user/:user/exists", basicAnswer)
	app.GET("/user/:user/member-of/:group", basicAnswer)
	app.POST("/user/:user/authenticate", basicAnswer)
	app.PUT("/user/:user/password", basicAnswer)
	app.PUT("/user/:user/password-never-expires", basicAnswer)
	app.PUT("/user/:user/password-expires", basicAnswer)
	app.PUT("/user/:user/enable", basicAnswer)
	app.PUT("/user/:user/disable", basicAnswer)
	app.PUT("/user/:user/move", basicAnswer)
	app.PUT("/user/:user/unlock", basicAnswer)

	app.Logger.Fatal(app.Start(l.HostPort()))
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
	newPassword := ""
	password = "\"" + password + "\""
	pbyte := []rune(password)
	for i := 0; i < len(password); i++ {
		// fmt.Println(pbyte[i], pbyte[i]&0xFF)
		newPassword += string(pbyte[i] & 0xFF)
	}
	return newPassword
}

func encodePassword2(password string) string {
	passrune := []rune(password)
	for k, v := range passrune {
		passrune[k] = v & 0xFF
	}
	return string(passrune)
}

func basicAnswer(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": "other"})
}
