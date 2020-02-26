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
			handler.Index,
			false,
			nil,
		},
		router.Route{
			"TeachersShow",
			"GET",
			"/{userId}",
			handler.Show,
			false,
			nil,
		},
		router.Route{
			"TeacherRegister",
			"POST",
			"/register",
			handler.Register,
			false,
			nil,
		},
	},
}
