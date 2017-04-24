package http

import (
	"net/http"
	"testing"
	"time"

	"github.com/martinz0/mux"
)

func TestHTTP(t *testing.T) {
	r := NewRouter()

	r.Index("static/index.html")
	r.Static("static")

	r.Group("/lor/v1", func(r *router) {
		r.Get("/users/:uid", testHandler)
	})
	ListenAndServe(":8080", r)
}

func testHandler(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	time.Sleep(5e9)
	w.Write([]byte(ps.Get("uid")))
}
