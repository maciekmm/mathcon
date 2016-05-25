package mathcon

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = []string{
	"templates/includes/footer.tmpl",
	"templates/includes/header.tmpl",
	"templates/main.tmpl",
}

var cachedTemplates = template.Must(template.New("").ParseFiles(templates...))

func init() {
	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/contest/", contestHandler)
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if err := cachedTemplates.ExecuteTemplate(w, "main", nil); err != nil {
		fmt.Fprint(w, err)
		return
	}
	contest
}
