package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/ppreeper/addictgo/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ppreeper/addictgo/ldap"
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

	app := fiber.New()
	app.Use(compress.New())
	app.Use(cors.New())

	fmt.Println(encodePassword("p4ssw04d"))

	app.Get("/swagger/*", swagger.HandlerDefault)
	OU(app)
	Group(app)
	User(app)
	Other(app)

	app.Listen(l.HostPort())
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
		fmt.Println(pbyte[i], pbyte[i]&0xFF)
		newPassword += string(pbyte[i] & 0xFF)
	}
	return newPassword
}
