package absent

import (
	"backend-qrcode/absent/handler"
	"backend-qrcode/router"
)

var (
	student = "student"
	teacher = "teacher"
)

var Routes = router.RoutePrefix{
	"/absents",
	[]router.Route{
		router.Route{
			"AbsentIndex",
			"GET",
			"",
			handler.Index,
			true,
			&teacher,
		},
		router.Route{
			"AbsentCreate",
			"POST",
			"",
			handler.Create,
			true,
			&student,
		},
		router.Route{
			"AbsentSetup",
			"POST",
			"/setup",
			handler.Setup,
			true,
			&teacher,
		},
	},
}
