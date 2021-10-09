package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is a struct that represents a user in our application
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Post is a struct that represents a single post
type Post struct {
	Author    User   `json:"author"`
	Caption   string `json:"caption"`
	ImageUrl  string `json:"imageUrl"`
	TimeStamp string `json:"time Stamp"`
}

var allposts []Post = []Post{}

var allusers []User = []User{}

func main() {
	router := mux.NewRouter()

	// Create User
	router.HandleFunc("/users", createUser).Methods("POST")

	// get user
	router.HandleFunc("/users/{id}", getUser).Methods("GET")

	// Create Post
	router.HandleFunc("/posts", createPost).Methods("POST")

	// get post
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	// all posts
	router.HandleFunc("/posts/users/{id}", getAllPosts).Methods("GET")

	http.ListenAndServe(":5000", router)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	allusers = append(allusers, newUser)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allusers)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	allposts = append(allposts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(allposts)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(allusers) {
		w.WriteHeader(404)
		w.Write([]byte("No user found with specified ID"))
		return
	}

	user := allusers[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(allposts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := allposts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allposts)
}
