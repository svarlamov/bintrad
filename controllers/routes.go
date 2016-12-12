package controllers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/controllers/api"
	"net/http"
)

var Log = config.Conf.GetLogger()

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	router = router.StrictSlash(true)

	// TODO: View routes
	//router.HandleFunc("/r/{authPathToken}", Use(api.V0_API_Activate_Registration_Link)).Methods("GET")
	//router.HandleFunc("/l/{sentFromEmailEncoded}/{tracePathToken}", Use(api.V0_API_Trace_Request)).Methods("GET")

	// API V0 Routes
	// TODO: Add in contracts routes, create a new 'my user data' route, and (as time allows) a leaderboard endpoint
	apiV0Router := router.PathPrefix("/api/v0").Subrouter()
	apiV0Router = apiV0Router.StrictSlash(true)
	apiV0Router.HandleFunc("/", Use(api.V0_API)).Methods("GET")
	apiV0Router.HandleFunc("/authenticate", Use(api.V0_API_Authenticate)).Methods("POST")
	//apiV0Router.HandleFunc("/traces", Use(api.V0_API_Init_Trace_Pixel, RequireValidTokenForAPI)).Methods("POST")

	// Ensure that the API V0 subrouter gets called
	router.PathPrefix("/api/v0/").Handler(apiV0Router)

	// TODO: Static file routes

	// Setup CSRF Protection
	csrfHandler := nosurf.New(router)
	// Exempt API routes and Static files
	exFn := func(r *http.Request) bool {
		// TODO: At some point need to figure out the full CSRF setup
		return true
	}
	csrfHandler.ExemptFunc(exFn)

	return Use(csrfHandler.ServeHTTP, EnsureSecureConnection, GetContext)
}

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
