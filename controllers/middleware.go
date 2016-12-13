package controllers

import (
	"fmt"
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/context"
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
	"strings"
)

func GetContext(handler http.Handler) http.HandlerFunc {
	// Set the context here
	return func(w http.ResponseWriter, r *http.Request) {
		var passItOn = func() {
			handler.ServeHTTP(w, r)
			context.ClearRequest(r)
		}
		var handleNonEmptyToken = func(value string) {
			token := models.AccessToken{Token: value}
			err := token.FindByToken()
			if err != nil {
				passItOn()
			} else {
				context.SetToken(r, token)
				passItOn()
			}
		}

		if c, err := r.Cookie(config.Conf.TokenCookieName); err == nil {
			if c.Value != "" {
				handleNonEmptyToken(c.Value)
			} else {
				passItOn()
			}
		} else if ah := r.Header.Get("Authorization"); ah != "" {
			if len(ah) > 6 && strings.ToUpper(ah[0:7]) == "BEARER " {
				val := ah[7:]
				if val != "" {
					handleNonEmptyToken(val)
				} else {
					passItOn()
				}
			} else {
				passItOn()
			}
		} else {
			passItOn()
		}
	}
}

func RequireValidTokenForAPI(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			utils.JSONDetailed(w, utils.APIResponse{Message: "Unauthorized", Debug: "Invalid or missing access token header/cookie"}, http.StatusUnauthorized)
		}
		token, err := context.GetToken(r)
		if err != nil {
			respondUnauthorized()
			return
		}
		if !(token.IsValid()) {
			respondUnauthorized()
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}

func RequireValidTokenForView(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var respondUnauthorized = func() {
			utils.TemporaryRedirect(w, r, "/")
		}
		token, err := context.GetToken(r)
		if err != nil {
			respondUnauthorized()
			return
		}
		if !(token.IsValid()) {
			respondUnauthorized()
			return
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}

func EnsureSecureConnection(handler http.Handler) http.HandlerFunc {
	// Assumes header names from AWS ELB setup, and should pass silently in a dev environment due to the port checks
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Header.Get("X-Forwarded-Port") == "80" || r.Header.Get("X-Forwarded-Port") == "443") && r.Header.Get("X-Forwarded-Proto") == "http" {
			http.Redirect(w, r, fmt.Sprintf("%s%s", config.Conf.Home, r.URL.Path), http.StatusMovedPermanently)
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}
