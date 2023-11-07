package main

import (
	"html/template"
	"log/slog"
	"net/http"
)

func contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contacts" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		slog.Error("unable to read template", "error", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		slog.Error("unable to execute template", "error", err)
		http.Error(w, "Internal Server Error", 500)
	}
}
