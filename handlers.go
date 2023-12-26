package main

import (
	"html/template"
	"net/http"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dbGetAllArticles()
	catch(err)

	t, _ := template.ParseFiles("templates/base.html.templ", "templates/index.html.templ")
	err = t.Execute(w, articles)
	catch(err)
}
