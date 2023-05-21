package routes

import (
	"api/src/controller"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controller.CreateUser,
		Authentication: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Function:       controller.RetrieveUsers,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controller.RetrieveUser,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controller.UpdateUser,
		Authentication: false,
	},
	{
		Uri:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controller.DeleteUser,
		Authentication: false,
	},
}
