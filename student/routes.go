package student

import (
	"backend-qrcode/router"
	handler "backend-qrcode/student/handler"
)

var (
	student = "student"
	teacher = "teacher"
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
			"StudentSchedule",
			"GET",
			"/schedule/{id}",
			handler.Schedule,
			true,
			&teacher,
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
