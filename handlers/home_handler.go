package handlers

import "net/http"

// Home is an example of struct
type Home struct {
	Name  string
	Email string
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := serveTemplate(w, "home", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
