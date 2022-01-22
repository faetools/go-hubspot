package httpservice

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HTTPHandler handles http requests.
type HTTPHandler struct{}

// UserAdd adds a new user.
// (POST /user).
func (h HTTPHandler) UserAdd(c echo.Context) error {
	// ctx := c.Request().Context()
	// l := log.WithContext(ctx)

	return c.NoContent(http.StatusNotImplemented)
}

// UserGet gets a user by ID.
// (GET /user/{id}).
func (h HTTPHandler) UserGet(c echo.Context, userID string) error {
	// ctx := c.Request().Context()
	// l := log.WithContext(ctx)

	return c.NoContent(http.StatusNotImplemented)
}

type UserRecord struct {
	Identifier string

	FirstName string
	LastName  string
	Email     string

	DateCreated  time.Time
	DateModified time.Time
}
