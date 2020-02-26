package student

import (
	"backend-qrcode/router"
	handler "backend-qrcode/student/handler"
)

var Routes = router.RoutePrefix{
	"/students",
	[]router.Route{
		router.Route{
			"StudentsIndex",
			"GET",
			"",
			handler.Index,
			false,
			nil,
		},
		router.Route{
			"StudentShow",
			"GET",
			"/{userId}",
			handler.Show,
			true,
			nil,
		},
		router.Route{
			"StudentRegister",
			"POST",
			"/register",
			handler.Register,
			false,
			nil,
		},
	},
}
