package user

import (
	"backend-qrcode/router"
	"backend-qrcode/user/handler"
)

var Routes = router.RoutePrefix{
	"",
	[]router.Route{
		router.Route{
			"UserLogin",
			"POST",
			"/login",
			handler.Login,
			false,
			nil,
		},
	},
}
