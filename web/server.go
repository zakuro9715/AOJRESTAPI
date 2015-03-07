package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	logFormat = "%v[%v] %v"
  UsersPath = "/users/"
  ProblemsPath = "/problems/"
)

func Log(handler http.Handler) http.Handler {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.LstdFlags)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		logger.Printf(logFormat, r.URL, r.Method, r.RemoteAddr)
		fmt.Print(&buf)
	})
}

func registerHandlers() {
	http.HandleFunc(UsersPath, usersHandler)
	http.HandleFunc(ProblemsPath, problemsHandler)
}

func RunServer(host, port string) {
	registerHandlers()
	http.ListenAndServe(host+":"+port, Log(http.DefaultServeMux))
}
