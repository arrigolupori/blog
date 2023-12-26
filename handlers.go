package main

import (
	"errors"
	"html/template"
	"net/http"
	"os"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	password := r.FormValue("password")

	if password != os.Getenv("AL_BLOG_PASS") {
		err := errors.New("wrong password")
		catch(err)
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

func GetPost(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*BlogPost)
	t, _ := template.ParseFiles("templates/base.go.html", "templates/post.go.html")
	err := t.Execute(w, article)
	catch(err)
}

// func EditPost(w http.ResponseWriter, r *http.Request) {
// 	article := r.Context().Value("article").(*BlogPost)

// 	t, _ := template.ParseFiles("templates/base.html", "templates/edit.html")
// 	err := t.Execute(w, article)
// 	catch(err)
// }

// func UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	article := r.Context().Value("article").(*BlogPost)

// 	title := r.FormValue("title")
// 	content := r.FormValue("content")
// 	newPost := &BlogPost{
// 		Title:   title,
// 		Content: template.HTML(content),
// 	}
// 	fmt.Println(newPost.Content)
// 	err := dbUpdatePost(strconv.Itoa(article.ID), newPost)
// 	catch(err)
// 	http.Redirect(w, r, fmt.Sprintf("/articles/%d", article.ID), http.StatusFound)
// }

// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	article := r.Context().Value("article").(*BlogPost)
// 	err := dbDeletePost(strconv.Itoa(article.ID))
// 	catch(err)

// 	http.Redirect(w, r, "/", http.StatusFound)
// }
