package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

var router *chi.Mux
var db *sql.DB

type BlogPost struct {
	ID      int           `json:"id"`
	Slug    string        `json:"slug"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.go.html", "templates/new.go.html")
	err := t.Execute(w, nil)
	catch(err)
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error
	db, err = connect()
	catch(err)
}

func main() {
	// AL_BLOG_PASS := os.Getenv("AL_BLOG_PASS")
	// fmt.Println("Blog pass:" + AL_BLOG_PASS)

	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	_, err = connect()
	catch(err)

	router.Use(ChangeMethod)
	router.Get("/", GetAllPosts)

	router.Route("/new", func(r chi.Router) {
		r.Get("/", NewPost)
		r.Post("/", CreatePost)
	})

	router.Route("/{postSlug}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Get("/", GetPost)
		// r.Put("/", UpdatePost)
		// r.Delete("/", DeletePost)
		// r.Get("/edit", EditPost)
	})

	err = http.ListenAndServe(":8080", router)
	catch(err)
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
