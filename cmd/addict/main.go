package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/ppreeper/addictgo/docs"
	"github.com/ppreeper/addictgo/internal/router"
)

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
	var url, username, password, port string
	// var lconn pkg.LDAP
	// lconn.StartTime = time.Now()
	// var l LDAP
	flag.StringVar(&url, "url", LookupEnvOrString("ADDICT_URL", ""), "Address for Active Directory")
	flag.StringVar(&username, "u", LookupEnvOrString("ADDICT_USER", ""), "BIND username")
	flag.StringVar(&password, "p", LookupEnvOrString("ADDICT_PASS", ""), "BIND password")
	flag.StringVar(&port, "port", LookupEnvOrString("ADDICT_PORT", "3001"), "Port to listen on")	
	// flag.StringVar(&lconn.URL, "url", LookupEnvOrString("ADDICT_URL", ""), "Address for Active Directory")
	// flag.StringVar(&lconn.Username, "u", LookupEnvOrString("ADDICT_USER", ""), "BIND username")
	// flag.StringVar(&lconn.Password, "p", LookupEnvOrString("ADDICT_PASS", ""), "BIND password")
	// flag.StringVar(&lconn.Port, "port", LookupEnvOrString("ADDICT_PORT", "3001"), "Port to listen on")
	flag.Parse()
	if url == "" {
		fmt.Println("no ldap url provided")
		return
	}
	if username == "" || password == "" {
		fmt.Println("no username or password provided")
		return
	}

	r := router.NewRouter(url, username, password, port)

	// WEB
	// web := app.Group("/web")
	// web.GET("", func(c echo.Context) error {
	// 	component := pages.Index("Hello")
	// 	return component.Render(c.Request().Context(), c.Response().Writer)
	// })

	r.Logger.Fatal(r.Start(":" + port))
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
