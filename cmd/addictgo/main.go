package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ppreeper/addictgo"
)

func main() {
	var addictURL string
	var addictUser string
	var addictPass string
	var port string
	flag.StringVar(&addictURL, "url", LookupEnvOrString("ADDICT_URL", ""), "Address for Active Directory")
	flag.StringVar(&addictUser, "u", LookupEnvOrString("ADDICT_USER", ""), "BIND username")
	flag.StringVar(&addictPass, "p", LookupEnvOrString("ADDICT_PASS", ""), "BIND password")
	flag.StringVar(&port, "port", LookupEnvOrString("ADDICT_PORT", "3001"), "Port to listen on")
	flag.Parse()
	if addictURL == "" {
		fmt.Println("no ldap url provided")
		return
	}
	if addictUser == "" || addictPass == "" {
		fmt.Println("no username or password provided")
		return
	}

	var a = addictgo.API{
		URL:      addictURL,
		Username: addictUser,
		Password: addictPass,
		Scope:    addictgo.GetScope(addictUser),
		Port:     port,
	}

	app := fiber.New()
	app.Use(compress.New())
	app.Use(cors.New())

	fmt.Println(encodePassword("p4ssw04d"))

	a.Other(app)
	a.OU(app)
	a.Group(app)
	a.User(app)

	app.Listen(a.HostPort())
}

//LookupEnvOrString provides 12 Factor for string vars
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

//LookupEnvOrInt provides 12 Factor for int vars
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
		fmt.Println(pbyte[i]&0xFF, (pbyte[i]>>8)&0xFF)
		newPassword += string(pbyte[i] & 0xFF)
	}
	return newPassword
}
