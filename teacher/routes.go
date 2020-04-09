package teacher

import (
	"backend-qrcode/router"
	"backend-qrcode/teacher/handler"
)

var (
	admin   = "admin"
	teacher = "teacher"
	student = "student"
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
			"TeacherRegister",
			"POST",
			"/register",
			handler.Register,
			false,
			nil,
		},
		router.Route{
			"TeacherEdit",
			"PUT",
			"",
			handler.Edit,
			true,
			&teacher,
		},
		router.Route{
			"TeachersShow",
			"GET",
			"/",
			handler.Show,
			true,
			&teacher,
		},
		router.Route{
			"TeachersShow",
			"GET",
			"/{userId}",
			handler.Show,
			false,
			nil,
		},
	},
}
