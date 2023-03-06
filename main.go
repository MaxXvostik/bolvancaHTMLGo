package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {

	fmt.Println("port :3000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
