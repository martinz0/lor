package http

import (
	"net/http"
	spath "path"

	"github.com/martinz0/mux"
)

type router struct {
	*mux.Mux

	group string
}

func NewRouter() *router {
	m := mux.New()
	return &router{
		Mux: m,
	}
}

func (r *router) Index(file string) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request, ps mux.Params) {
		http.ServeFile(w, r, file)
	})
}

func (r *router) Static(dir string) {
	r.Get(pathJoin(dir, "*"), func(w http.ResponseWriter, r *http.Request, ps mux.Params) {
		http.StripPrefix(pathJoin("/", dir), http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
	})
}

func (r *router) Group(path string, f func(r *router)) {
	nr := &router{
		Mux:   r.Mux,
		group: path,
	}
	f(nr)
}

func (r *router) Get(path string, handler mux.Handler) {
	r.Handle("GET", pathJoin(r.group, path), handler)
}

func (r *router) Post(path string, handler mux.Handler) {
	r.Handle("POST", pathJoin(r.group, path), handler)
}

func (r *router) Put(path string, handler mux.Handler) {
	r.Handle("PUT", pathJoin(r.group, path), handler)
}

func (r *router) Delete(path string, handler mux.Handler) {
	r.Handle("DELETE", pathJoin(r.group, path), handler)
}

func pathJoin(p1, p2 string) string {
	return spath.Join(p1, p2)
}
