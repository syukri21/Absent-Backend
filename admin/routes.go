package admin

import (
	handler "backend-qrcode/admin/handler"
	"backend-qrcode/router"
)

var (
	student = "student"
	teacher = "teacher"
	admin   = "admin"
)

var Routes = router.RoutePrefix{
	"/students",
	[]router.Route{
		router.Route{
			Name:        "AdminsIndex",
			Method:      "GET",
			Pattern:     "",
			HandlerFunc: handler.Index,
			Protected:   true,
			Previlage:   &admin,
		},

		router.Route{
			Name:        "AdminsShow",
			Method:      "GET",
			Pattern:     "/{userId}",
			HandlerFunc: handler.Show,
			Protected:   true,
			Previlage:   &admin,
		},

		router.Route{
			Name:        "AdminRegister",
			Method:      "POST",
			Pattern:     "/register",
			HandlerFunc: handler.Register,
			Protected:   true,
			Previlage:   &admin,
		},
	},
}
