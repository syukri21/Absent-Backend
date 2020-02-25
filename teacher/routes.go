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
			false,
			false,
		},
		router.Route{
			"TeachersShow",
			"GET",
			"/{userId}",
			handler.Show,
			false,
			false,
			false,
		},
		router.Route{
			"TeacherRegister",
			"POST",
			"/register",
			handler.Register,
			false,
			false,
			false,
		},
		router.Route{
			"TeacherLogin",
			"POST",
			"/login",
			handler.Login,
			false,
			false,
			false,
		},
	},
}
