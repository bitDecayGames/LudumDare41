package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/bitDecayGames/LudumDare41/server/routes"

	"github.com/gorilla/mux"
)

// TODO Turn timeout logic

const (
	// Networking
	port = 8080
)

var rts *routes.Routes

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	rts = routes.InitRoutes(r)

	log.Printf("Server started on %s", host)

	err := http.ListenAndServe(host, r)
	if err != nil {
		log.Panic(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignore websocket route.
		reqURI := strings.Split(r.RequestURI, "?")[0]
		if reqURI == routes.PubSubRoute {
			log.Printf("Skipping logging for %s", routes.PubSubRoute)
			next.ServeHTTP(w, r)
			return
		}

		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		responseDump, err := httputil.DumpResponse(rec.Result(), true)
		if err != nil {
			log.Println(err)
			return
		}

		// we copy the captured response headers to our new response
		w.WriteHeader(rec.Code)
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		data := rec.Body.Bytes()
		w.Write(data)

		log.Printf("%s\n\nRESPONSE\n%s", requestDump, responseDump)
	})
}
