package main

import (
	"backend-qrcode/absent"
	"backend-qrcode/admin"
	"backend-qrcode/course"
	"backend-qrcode/middleware"
	customRouter "backend-qrcode/router"
	"backend-qrcode/schedule"
	"backend-qrcode/socket"
	"backend-qrcode/student"
	"backend-qrcode/teacher"
	"backend-qrcode/user"

	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	//append user routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, user.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, admin.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, teacher.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, student.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, absent.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, course.Routes)
	customRouter.AppRoutes = append(customRouter.AppRoutes, schedule.Routes)

	for _, route := range customRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//check to see if route should be protected with jwt
			if r.Protected {
				handler = middleware.Middleware(r.HandlerFunc, r.Previlage)
			}

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	hub := socket.NewHub()
	go hub.Run()

	router.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})

	return router
}
