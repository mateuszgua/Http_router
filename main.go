package main

import (
	"errors"
	"fmt"
	"log"
	router "mateusz/http_router/router"
	"net/http"
	"os"
)

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET /")
	})
}

func asdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprint(w, "GET /asd")
		case http.MethodPost:
			fmt.Fprint(w, "POST / asd")
		default:
			fmt.Fprint(w, "Not Found")
		}
	})
}

func asdZxcHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET /asd/zxc")
	})
}

func asdZxcQweHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET /asd/zxc/qwe")
	})
}

func zxcHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET /zxc")
	})
}

func qweHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET /qwe")
	})
}

func main() {

	r := router.NewRouter()

	r.Methods(http.MethodGet).Handler(`/`, indexHandler())
	r.Methods(http.MethodGet, http.MethodPost).Handler(`/asd`, asdHandler())
	r.Methods(http.MethodGet).Handler(`/asd/zxc`, asdZxcHandler())
	r.Methods(http.MethodGet).Handler(`/asd/zxc/qwe`, asdZxcQweHandler())
	r.Methods(http.MethodGet).Handler(`/zxc`, zxcHandler())
	r.Methods(http.MethodGet).Handler(`/qwe`, qweHandler())

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	port := fmt.Sprintf(":%s", httpPort)

	log.Printf("Starting server on http://localhost%s ... \n", port)
	err := http.ListenAndServe(port, r)
	// http.ListenAndServe(":8080", r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
