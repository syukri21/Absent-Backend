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
		},
		router.Route{
			"StudentShow",
			"GET",
			"/{userId}",
			handler.Show,
			false,
		},
	},
}
