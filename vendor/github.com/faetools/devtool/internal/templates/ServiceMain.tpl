package service

import (
	"github.com/labstack/echo/v4"
	"github.com/faetools/{{.Name}}/internal/httpservice"
	"net/http"
)

func RunService(
	{{if .IsHTTPService}}pubhttpsrv *echo.Echo,{{end}}
) error {
	{{if .IsHTTPService}}
	httpservice.RegisterHandlers(pubhttpsrv, httpservice.HTTPHandler{
	})
	{{end}}

	return nil
}
