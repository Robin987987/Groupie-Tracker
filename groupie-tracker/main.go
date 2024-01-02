package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var pages *template.Template

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.Handle("/css/", http.FileServer(http.Dir("./templates")))
	mux.Handle("/script/", http.FileServer(http.Dir("./templates")))
	fmt.Println("Server is running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	getApi()
	if r.URL.Path == "/" {
		if pages.Execute(w, theArtists) != nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			return
		}
	}
	if r.URL.Path == "/{Id:}" {

		id, err := strconv.Atoi(r.URL.Path[:1])
		if err != nil {
			fmt.Println("error")
		}
		page := template.Must(template.ParseFiles("templates/artist.html"))
		if page.Execute(w, theArtists[id-1]) != nil {
			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			return
		}
	}

}

func init() {
	pages = template.Must(template.ParseGlob("templates/*.html"))
}
