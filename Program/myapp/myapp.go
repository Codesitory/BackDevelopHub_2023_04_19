package myapp

import (
	"net/http"
)

func NewHTTPHandler() http.Handler {

	router := http.NewServeMux()
	router.HandleFunc("/auth/singup", singUp)
	router.HandleFunc("/auth/login", login)
	return router
}
