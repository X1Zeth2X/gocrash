package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Post Struct (Model)
type Post struct {
	ID      string  `json:"id"`
	Content string  `json:"content"`
	Title   string  `json:"title"`
	Author  *Author `json:"author"`
}

// Author struct (Model)
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Initialize posts
var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, post := range posts {
		if post.ID == params["id"] {
			json.NewEncoder(w).Encode(post)

			return
		}
	}

	json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post

	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(10000000))

	// Append post to posts
	posts = append(posts, post)
	// Encode and return post
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Set params to get book id
	params := mux.Vars(r)

	for i, post := range posts {
		if post.ID == params["id"] {
			// GO syntax for deleting items from an array
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}

	// Return the modified array
	json.NewEncoder(w).Encode(posts)
}

func setMocks() {
	posts = append(posts, Post{
		ID:      "1",
		Content: "Hello world!",
		Title:   "My first post",
		Author: &Author{
			FirstName: "Seth",
			LastName:  "Smith",
		},
	})
	posts = append(posts, Post{
		ID:      "2",
		Content: "Zimmerman in GOlang?",
		Title:   "Is it possible.",
		Author: &Author{
			FirstName: "Bor",
			LastName:  "Pyke",
		},
	})
}

func handleRequests() {
	// Initialize Router
	r := mux.NewRouter()

	// Router handlers / Endpoints
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", getPost).Methods("GET")
	r.HandleFunc("/api/post", createPost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	fmt.Println("Server running on port localhost:5000")
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}

func main() {
	setMocks()
	handleRequests()
}
