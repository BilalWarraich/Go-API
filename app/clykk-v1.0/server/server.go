package server

import (
	"log"
	"net/http"

	"github.com/clykk-user-atif/config"
	"github.com/valyala/fasthttp"
)

func Start() {

	r, err := initRouter()
	if err != nil {
		log.Fatalf("[ERROR] unable to initialize router, err : %v", err)
	}

	log.Printf("server listening at %v on http...", config.HTTPAddr())
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Fatalf("Error in server: %s", err)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}
