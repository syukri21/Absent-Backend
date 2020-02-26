package absent

import (
	"backend-qrcode/absent/handler"
	"backend-qrcode/router"
)

var student = "student"

var Routes = router.RoutePrefix{
	"/absents",
	[]router.Route{
		router.Route{
			"AbsentCreate",
			"POST",
			"",
			handler.Create,
			true,
			&student,
		},
	},
}
