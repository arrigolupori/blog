package main

import (
	"html/template"
	"net/http"
	"os"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	password := r.FormValue("password")

	if password != os.Getenv("AL_BLOG_PASS") {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	post := &BlogPost{
		Title:   title,
		Content: template.HTML(content),
	}

	err := dbCreatePost(post)
	catch(err)

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := dbGetAllPosts()
	catch(err)

	t, _ := template.ParseFiles("templates/base.go.html", "templates/index.go.html")
	err = t.Execute(w, posts)
	catch(err)
}
