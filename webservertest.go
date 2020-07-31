package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

type Articles []Article

func fetchAllArticles(rw http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "First Article", Desc: "The first article ever.", Content: "This article doesn't contain shit."}}
	fmt.Println("All Articles retrieved!")
	json.NewEncoder(rw).Encode(articles)

}

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homepage endpoint reached!")
}

func handleRequests() {
	theRouter := mux.NewRouter().StrictSlash(true)
	theRouter.HandleFunc("/", homePage)
	theRouter.HandleFunc("/allarticles", fetchAllArticles)
	log.Fatal(http.ListenAndServe(":8082", theRouter))
}

func main() {
	handleRequests()
}
