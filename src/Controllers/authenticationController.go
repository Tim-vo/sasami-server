package authenticationService

import (
	authenticationService "github.com/Tim-vo/sasami-server/src/Services/authenticationService"
	"github.com/gorilla/mux"
)

func addAuthenticationController(router *mux.Router) {
	router.HandleFunc("/register", authenticationService.Register).Methods("POST")
}
