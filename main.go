package main

import (
	"html/template"
	"log"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var router *chi.Mux
var db *sql.DB

type BlogPost struct {
	ID      int           `json:"id"`
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
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = connect()
	catch(err)

	router.Use(ChangeMethod)
	router.Get("/", GetAllPosts)

	router.Route("/new", func(r chi.Router) {
		r.Get("/", NewPost)
		r.Post("/", CreatePost)
	})

	// router.Route("/{postID}", func(r chi.Router) {
	// 	r.Use(PostCtx)
	// })

	err = http.ListenAndServe(":8005", router)
	catch(err)
}
