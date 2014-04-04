package main

import (
		"fmt"
		"html/template"
		"net/http"
		//"errors"
)

func main() {

	templates = template.New("titleTest").Funcs(funcMap)
	templates = template.Must(templates.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
	
	http.HandleFunc("/", frontHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	http.ListenAndServe(":8080", nil)
	fmt.Println("test")
}


