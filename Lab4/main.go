package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

// SharkAttack is a struct that represents a shark attack from the JSON file
type SharkAttack struct {
	Date     string `json:"date"`
	Year     string `json:"year"`
	Type     string `json:"type"`
	Country  string `json:"country"`
	Area     string `json:"area"`
	Location string `json:"location"`
}

// posts is a slice of SharkAttack structs
var posts []SharkAttack

// loadDataFromJSON reads the data from the JSON file and stores it in the posts slice
func loadDataFromJSON() {
	file, err := os.Open("global-shark-attack.json")
	if err != nil {
		fmt.Println("Error occured while opening the file:", err)
		return
	}
	defer file.Close()

	var data []SharkAttack
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error occured while decoding the JSON:", err)
		return
	}

	for i := 0; i < 10; i++ {
		index := rand.Intn(len(data))
		posts = append(posts, data[index])
	}
}

// getPosts returns all the posts in the posts slice
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// addPost adds a new post to the posts slice
func addPost(w http.ResponseWriter, r *http.Request) {
	var newPost SharkAttack
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

// deletePost deletes a post from the posts slice
func deletePost(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Path[len("/posts/"):]
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(posts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	deletedAttack := posts[index]
	posts = append(posts[:index], posts[index+1:]...)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedAttack)
}

// updatePost updates a post in the posts slice
func updatePost(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Path[len("/posts/"):]
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(posts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	var updatedPost SharkAttack
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts[index] = updatedPost
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

// main is the entry point of the program
func main() {
	loadDataFromJSON()

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPosts(w, r)
		} else if r.Method == http.MethodPost {
			addPost(w, r)
		}
	})

	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			deletePost(w, r)
		} else if r.Method == http.MethodPut {
			updatePost(w, r)
		}
	})

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
