package routes

import (
	"api/src/controlers"
	"net/http"
)

var rotesUsers = []Route{
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controlers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controlers.SearchAllUser,
		RequiresAuthentication: false,
	}, {
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodGet,
		Function:               controlers.SearchUser,
		RequiresAuthentication: false,
	}, {
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodPut,
		Function:               controlers.UpdateUser,
		RequiresAuthentication: false,
	}, {
		URI:                    "/usuarios/{usuarioId}",
		Method:                 http.MethodDelete,
		Function:               controlers.DeleteUser,
		RequiresAuthentication: false,
	},
}
