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

func ChangeMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := r.PostFormValue(`_method`); method {
			case http.MethodPut:
				fallthrough
			case http.MethodPatch:
				fallthrough
			case http.MethodDelete:
				r.Method = method
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func connect() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", "./data.sqlite")

	if err != nil {
		return nil, err
	}

	sqlStmt := `create table if not exists articles (id integer not null primary key autoincrement, title text, content text);`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// func dbCreateArticle(db *sql.DB, article *Article) error {
// 	query, err := db.Prepare(`insert into articles(title, content) values (?, ?)`)

// 	if err != nil {
// 		return err
// 	}

// 	defer query.Close()

// 	_, err = query.Exec(article.Title, article.Content)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }



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
