package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Page struct {
	Title string
	Body  []byte	
}

func loadPage(title string) (*Page, error) {
	filename := "templates/" + title + ".gohtml"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := tpl.ExecuteTemplate(w, tmpl+".gohtml", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/notFound", http.StatusFound)
	}
	title := "index"
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, title, page)
}

func hideWindowHandler(w http.ResponseWriter, r *http.Request) {
	title := "hideWindow"
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/notFound", http.StatusFound)
		return
	}
	renderTemplate(w, title, page)
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	title := "404"
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, title, page)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hideWindow", hideWindowHandler)
	http.HandleFunc("/notFound", notFoundPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
