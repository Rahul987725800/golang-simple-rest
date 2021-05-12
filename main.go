package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(rw http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "test", Desc: "test description", Content: "hello world"},
	}
	fmt.Printf("Endpoint Hit: All Articles Endpoint")
	_ = json.NewEncoder(rw).Encode(articles)
}
func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homepage Endpoint Hit")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
func main() {
	handleRequest()
}
