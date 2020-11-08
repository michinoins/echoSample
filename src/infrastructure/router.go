package infrastructure

import (
	controllers "echoSample/src/interfaces/api"
	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	// Echo instance
	e := echo.New()
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

	//rest思想なのでpathは同じにしてメソッドで分ける
	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		userController.Delete(id)
		return c.String(http.StatusOK, "deleted")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
