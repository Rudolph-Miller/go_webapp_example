package handlersV1

import (
	"github.com/Rudolph-Miller/go_webapp_example/models"
	"github.com/Rudolph-Miller/go_webapp_example/support"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func UserGroup(v *echo.Group) {
	g := v.Group("/users")

	g.GET("/:id", show)
}

func show(c echo.Context) error {
	cc := c.(*support.CustomContext)
	db := cc.DB
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, "INVALID ID")
	}

	password := c.QueryParam("password")
	user, err := models.FindUser(db, uint(id), password)

	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, &user)
}
