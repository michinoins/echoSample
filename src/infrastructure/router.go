package infrastructure

import (
	"echoSample/src/auth"
	controllers "echoSample/src/interfaces/api"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.String(http.StatusOK, "created")
	})

	//rest思想なので同じパスでメソッド名を分ける
	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		userController.Delete(id)
		return c.String(http.StatusOK, "deleted")
	})

	// Login router
	e.POST("/login", func(c echo.Context) error {
		return auth.Login(c)
	})

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/sample", func(c echo.Context) error {
		auth.Restricted(c)
		return c.String(http.StatusOK, "passToken")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
