package main

import "net/http"

func RenderError(w http.ResponseWriter, status int, errorTemplate string) {
	w.WriteHeader(status)
	err := tpl.ExecuteTemplate(w, errorTemplate, nil)
	if err != nil {
		http.Error(w, http.StatusText(status), status)
	}
}
