package routes

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.RegisterUserCtrl())
	e.POST("/users/login", uctrl.LoginAuthCtrl())
	e.GET("/user", uctrl.GetUserByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/user", uctrl.UpdateUserCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.DELETE("/user", uctrl.DeleteUserCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

}