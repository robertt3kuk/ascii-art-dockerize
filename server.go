package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-web/ascii"
)

type req struct {
	Asc string
}

const (
	filename = "./static/index.html"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case "/":
		mainpage(w, r)
	default:
		notfound(w, r)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/404.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusNotFound)
	t.Execute(w, nil)
}

func asciiart(w http.ResponseWriter, r *http.Request) {
	input, banner := r.FormValue("input"), r.FormValue("banner")
	str := new(req)

	str.Asc = ascii.Asciitext(input, banner)
	switch str.Asc {
	case "Incorrect template.":
		w.WriteHeader(http.StatusInternalServerError)
	case "Incorrect input.":
		w.WriteHeader(http.StatusBadRequest)

	case "Internal error.":
		w.WriteHeader(http.StatusInternalServerError)

	}

	t, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("mainpagehtml not found")
		return
	}
	err = t.Execute(w, str)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		asciiart(w, r)
		return
	}
	t, err := template.ParseFiles(filename)
	if err != nil {
		notfound(w, r)
		fmt.Println("mainpagehtml not found")
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error when executing template")
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Printf("starting ascii-art at localhost:3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
