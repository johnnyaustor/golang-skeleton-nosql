package server

import (
	"context"
	"log"

	"github.com/johnnyaustor/golang-skeleton-nosql/server/api"
	"github.com/johnnyaustor/golang-skeleton-nosql/server/configuration"
	"github.com/johnnyaustor/golang-skeleton-nosql/server/database"
	"github.com/labstack/echo/v4"
)

func init() {
	// read env
	configuration.ReadConfiguration()
}

func RunServer() {
	// Connect to Database
	db := database.ConnectDatabase()
	defer func() {
		if err := db.Connection.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
		}
	}()

	log.Println("Run Server")
	e := echo.New()

	// Initial Beans
	api.InitialBeans(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
