package main

import (
	"BackDevelopHub/Program/myapp"
	"net/http"
)

const DEVELOPHUB_SERVER_PORT = ":9190"

func main() {
	http.ListenAndServe(DEVELOPHUB_SERVER_PORT, myapp.NewHTTPHandler())
}
