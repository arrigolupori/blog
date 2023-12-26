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

type Article struct {
	ID      int           `json:"id"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
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
	_, err = connect()
	catch(err)

	router.Use(ChangeMethod)

	err = http.ListenAndServe(":8005", router)
	catch(err)
}
