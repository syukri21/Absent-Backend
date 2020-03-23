package course

import (
	"backend-qrcode/course/handler"
	"backend-qrcode/router"
)

var (
	student = "student"
	teacher = "teacher"
	admin   = "admin"
)

// Routes ...
var Routes = router.RoutePrefix{
	"/courses",
	[]router.Route{
		router.Route{
			"CoursesIndex",
			"GET",
			"",
			handler.Index,
			false,
			nil,
		},
		router.Route{
			"CourseCreate",
			"POST",
			"",
			handler.Create,
			true,
			&admin,
		},

		router.Route{
			"CoursesIndex",
			"DELETE",
			"/{id}",
			handler.Delete,
			false,
			nil,
		},
	},
}
