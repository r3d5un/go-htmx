package main

import (
	"net/http"
)

func contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contacts" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Contacts!"))

}
