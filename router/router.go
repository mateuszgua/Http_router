package route

import (
	my_err "mateusz/http_router/errors"
	trie "mateusz/http_router/trie"
	"net/http"
)

type Router struct {
	tree *trie.Tree
}

type route struct {
	methods []string
	path    string
	handler http.Handler
}

func NewRouter() *Router {
	return &Router{
		tree: trie.NewTree(),
	}
}

var (
	tmpRoute = &route{}
)

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}
func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	tmpRoute = &route{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := my_err.HandleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.Actions.Handler
	h.ServeHTTP(w, req)
}
