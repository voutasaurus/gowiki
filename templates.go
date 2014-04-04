package main

import (
		"fmt"
		"html/template"
		"net/http"
		//"errors"
)

var templates *template.Template

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		fmt.Println("Template " + tmpl + " cannot be rendered")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
