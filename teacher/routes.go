package teacher

import (
	"backend-qrcode/router"
	"backend-qrcode/teacher/handler"
)

var Routes = router.RoutePrefix{
	"/teachers",
	[]router.Route{
		router.Route{
			"TeachersIndex",
			"GET",
			"",
			handler.IndexHandler,
			false,
		},
		router.Route{
			"TeachersShow",
			"GET",
			"/{userId}",
			handler.ShowHandler,
			false,
		},
	},
}
