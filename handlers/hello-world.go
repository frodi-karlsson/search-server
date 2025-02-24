package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const tmplPath = "../templates/hello-world.html"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if err := AssertMethod(r, http.MethodGet); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		err = fmt.Errorf("error parsing template: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		err = fmt.Errorf("error executing template: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
