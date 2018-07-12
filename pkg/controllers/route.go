package controllers

import "net/http"

// Route represents a path, http method, handler and whether the path requires
// authenticated access or not. Intended to be used by a server to set up handlers
type Route struct {
	Path       string
	Method     []string
	Handler    func(w http.ResponseWriter, r *http.Request)
	Restricted bool
}
