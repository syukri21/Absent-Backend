package course

import (
	"backend-qrcode/course/handler"
	"backend-qrcode/router"
)

var (
	student = "student"
	teacher = "teacher"
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
			&teacher,
		},
	},
}
