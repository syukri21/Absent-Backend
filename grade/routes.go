package grade

import (
	handler "backend-qrcode/grade/handler"
	"backend-qrcode/router"
)

var (
	student = "student"
	teacher = "teacher"
)

var Routes = router.RoutePrefix{
	"/grades",
	[]router.Route{
		router.Route{
			"GradeShow",
			"GET",
			"/schedule/{id}",
			handler.ShowByScheduleID,
			true,
			&teacher,
		},
		router.Route{
			"GradeShow",
			"POST",
			"/schedule/{scheduleId}/student/{studentId}",
			handler.Create,
			true,
			&teacher,
		},
	},
}
