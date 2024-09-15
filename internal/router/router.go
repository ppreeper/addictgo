package router

import (
	"embed"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ppreeper/addictgo/internal/api"
	"github.com/ppreeper/addictgo/internal/web"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//go:embed static
var staticFiles embed.FS

func NewRouter(url, username, password, port string) *echo.Echo {
	lconn := &api.LDAP{
		URL:      url,
		Username: username,
		Password: password,
		Port:     port,
	}
	lconn.Scope = api.GetScope(lconn.Username)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(echoprometheus.NewMiddleware("addict"))
	// embed /static/ files
	app.StaticFS("/", staticFiles)

	// API
	api := app.Group("/api")
	api.GET("/swagger/*", echoSwagger.WrapHandler)
	api.GET("/status", lconn.StatusGet)
	api.GET("/metrics", echoprometheus.NewHandler())

	// API > Other
	api.GET("/all", lconn.AllGet)
	api.GET("/find/:filter", lconn.FindFilterGet)
	api.GET("/other", lconn.OtherGet)

	// API > OU
	api.GET("/ou", lconn.OUGet)
	api.POST("/ou", lconn.OUPost)
	api.GET("/ou/:ou", lconn.OUOUGet)
	api.DELETE("/ou/:ou", lconn.OUOUDelete)
	api.GET("/ou/:ou/exists", lconn.OUOUExists)

	// API > Group
	api.GET("/group", lconn.GroupGet)
	api.POST("/group", lconn.GroupPost)
	api.GET("/group/:group", lconn.GroupGroupGet)
	api.DELETE("/group/:group", lconn.GroupGroupDelete)
	api.GET("/group/:group/exists", lconn.GroupGroupExistsGet)
	api.POST("/group/:group/user/:user", lconn.GroupGroupUserUserPost)
	api.DELETE("/group/:group/user/:user", lconn.GroupGroupUserUserDelete)

	// API > User
	api.GET("/user", lconn.UsersAllGet)
	api.POST("/user", lconn.UserPost)
	api.GET("/user/:user", lconn.UserUserGet)
	api.DELETE("/user/:user", lconn.UserUserDelete)
	api.GET("/user/:user/exists", lconn.UserUserExistsGet)
	api.GET("/user/:user/member-of/:group", lconn.UserUserMemberofGroupGet)
	api.POST("/user/:user/authenticate", lconn.UserUserAuthenticatePost)
	api.PUT("/user/:user/password", lconn.UserUserPasswordPut)
	api.PUT("/user/:user/password-never-expires", lconn.UserUserPasswordneverexpiresPut)
	api.PUT("/user/:user/password-expires", lconn.UserUserPasswordExpiresPut)
	api.PUT("/user/:user/enable", lconn.UserUserEnablePut)
	api.PUT("/user/:user/disable", lconn.UserUserDisablePut)
	api.PUT("/user/:user/move", lconn.UserUserMovePut)
	api.PUT("/user/:user/unlock", lconn.UserUserUnlockPut)

	// WEB
	app.GET("/", web.Index)
	app.GET("/login", web.Login)

	return app
}
