package teacher

import (
	"backend-qrcode/router"
	teacherHandler "backend-qrcode/teacher/handler"
)

var Routes = router.RoutePrefix{
	"/teachers",
	[]router.Route{
		router.Route{
			"UsersIndex",
			"GET",
			"",
			teacherHandler.IndexHandler,
			false,
		},
	},
}
