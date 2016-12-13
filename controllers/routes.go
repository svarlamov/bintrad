package controllers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/controllers/api"
	"github.com/svarlamov/bintrad/controllers/views"
	"net/http"
)

var Log = config.Conf.GetLogger()

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	router = router.StrictSlash(true)

	router.HandleFunc("/", Use(views.V0_VIEWS_Index)).Methods("GET")
	router.HandleFunc("/tradingDesk", Use(views.V0_VIEWS_Trading_Desk, RequireValidTokenForView)).Methods("GET")

	// API V0 Routes
	// TODO: Add in leaderboard endpoint
	apiV0Router := router.PathPrefix("/api/v0").Subrouter()
	apiV0Router = apiV0Router.StrictSlash(true)
	apiV0Router.HandleFunc("/", Use(api.V0_API, GetContext)).Methods("GET")
	apiV0Router.HandleFunc("/authenticate", Use(api.V0_API_Authenticate, GetContext)).Methods("POST")
	apiV0Router.HandleFunc("/contracts/sessions", Use(api.V0_API_Start_Contract_Session, RequireValidTokenForAPI, GetContext)).Methods("POST")
	apiV0Router.HandleFunc("/contracts/sessions/{sessionId}", Use(api.V0_API_Finalise_Contract_Session, RequireValidTokenForAPI, GetContext)).Methods("POST")
	apiV0Router.HandleFunc("/users/me", Use(api.V0_API_Get_My_User_Data, RequireValidTokenForAPI, GetContext)).Methods("POST")
	//apiV0Router.HandleFunc("/leaderboard", Use(api.V0_API_Finalise_Contract_Session, RequireValidTokenForAPI, GetContext)).Methods("POST")

	// Ensure that the API V0 subrouter gets called
	router.PathPrefix("/api/v0/").Handler(apiV0Router)

	staticFS := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(staticFS).Methods("GET")

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
