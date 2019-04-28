package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"

)
func GenerateID()string{
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x",b)
}
type Post struct{
	Id string
	Title string
	Content string
}

func NewPost (id, title, content string)*Post{
	return &Post{id, title, content}
}

var posts map[string]*Post


func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Println(posts)

	t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	id := r.FormValue("id")
	post, found := posts[id]
	if !found{
		http.NotFound(w,r)
	}



	t.ExecuteTemplate(w, "write", post)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *Post

	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else{
		id = GenerateID()
		post := NewPost(id, title, content)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id==""{
		http.NotFound(w,r)
	}
	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("listening on port :3000" )
	posts = make(map [string] *Post, 0)

	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SavePost", savePostHandler)


	http.ListenAndServe(":3000", nil)
}
