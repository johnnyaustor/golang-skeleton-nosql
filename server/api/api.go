package api

import (
	"github.com/johnnyaustor/golang-skeleton-nosql/server/api/user"
	"github.com/johnnyaustor/golang-skeleton-nosql/server/database"
	"github.com/labstack/echo/v4"
)

func InitialBeans(e *echo.Echo, db *database.DB) {
	user.RegisterHandlers(e, user.NewService(user.NewRepository(db)))
}
