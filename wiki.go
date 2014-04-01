package main

import (
		"fmt"
		"html/template"
		"io/ioutil"
		"net/http"
		"regexp"
		"strings"
		// "html"
		//"errors"
)

type Page struct {
	Title string
	Body []byte
}

func titleDisplay(title string) string {
	return strings.Replace(title, "_", " ", -1)
}

func repl(semlink []byte) []byte {
	m := len(semlink)
	pageName := string(semlink[1:m-1])
	address := "/view/" + strings.Replace(pageName, " ", "_", -1)
	link := "<a href=\"" + address + "\">" + pageName + "</a>"
	return []byte(link)
}

func parseBody(body []byte) []byte {

	safe := template.HTMLEscapeString(string(body))

	safe = strings.Replace(safe, "\n", "<br>", -1)
	
	semlink := regexp.MustCompile(`\[[0-9A-Za-z_ ]+\]`) // note the space
	
	return semlink.ReplaceAllFunc([]byte(safe), repl)

}

func show(body []byte) template.HTML {
	return template.HTML(string(body))
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func frontHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/Front_Page", http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


var funcMap = template.FuncMap{
			"titleFmt":titleDisplay,
			"parse":parseBody,
			"show":show,
	}

var templates *template.Template // = template.Must(template.New("titleTest").Funcs(funcMap).ParseFiles("tmpl/edit.html", "tmpl/view.html"))

//var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
//.Funcs(funcMap))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		fmt.Println("Template " + tmpl + " cannot be rendered")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9_]+$")

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return //err = errors.New("Invalid Page Title")
		}
		fn(w, r, title)
	}
}

func main() {

	templates = template.New("titleTest").Funcs(funcMap)
	templates = template.Must(templates.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
	
	http.HandleFunc("/", frontHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8080", nil)
	fmt.Println("test")
}

/*	
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
*/

