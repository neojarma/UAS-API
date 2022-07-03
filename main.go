package main

import (
	"os"
	"uas_neoj/connection"
	"uas_neoj/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	group := e.Group("/api/v1")
	router.Router(group, db)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
