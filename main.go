package main

import (
	"os"

	"github.com/iboware/location-history/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DEFAULT_SERVER_LISTEN_ADDR = ":8080"

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	locationServer := server.NewLocationServer()
	server.RegisterHandlers(e, locationServer)

	var address string
	var ok bool
	if address, ok = os.LookupEnv("HISTORY_SERVER_LISTEN_ADDR"); !ok {
		address = DEFAULT_SERVER_LISTEN_ADDR
	}

	e.Logger.Fatal(e.Start(address))
}
