package schedule

import (
	"backend-qrcode/router"
	handler "backend-qrcode/student/handler"
)

var (
	student = "student"
	teacher = "teacher"
)

var Routes = router.RoutePrefix{
	"/schedules",
	[]router.Route{
		router.Route{
			"ScheduleIndex",
			"GET",
			"",
			handler.Index,
			true,
			&teacher,
		},
	},
}
