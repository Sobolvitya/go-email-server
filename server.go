package main

import (
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from goland Server!:%s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	page, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>ERROR</h1><div>%s</div>", err.Error())
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: file}, nil
}

func main() {
	log.Print("Server started...")
	http.HandleFunc("/view/", viewHandler)
	log.Print("/view/ assigned to viewHandler")
	log.Fatal(http.ListenAndServe(":8080", nil))
}