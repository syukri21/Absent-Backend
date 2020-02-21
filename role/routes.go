package role

import "backend-qrcode/router"

var Routes = router.RoutePrefix{
	"/roles",
	[]router.Route{
		router.Route{
			"RolesIndex",
			"GET",
			"",
			IndexHandler,
			false,
		},
		router.Route{
			"RolesCreate",
			"POST",
			"",
			CreateHandler,
			false,
		},
		router.Route{
			"RolesShow",
			"GET",
			"/{roleId}",
			ShowHandler,
			false,
		},
	},
}
