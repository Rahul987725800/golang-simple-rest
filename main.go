package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

var articles = Articles{
	Article{Title: "test", Desc: "test description", Content: "hello world"},
}

func allArticles(rw http.ResponseWriter, r *http.Request) {

	fmt.Printf("Endpoint Hit: All Articles Endpoint\n")
	_ = json.NewEncoder(rw).Encode(articles)
}
func testPostArticle(rw http.ResponseWriter, r *http.Request) {
	newArticle := Article{Title: "second", Desc: "test description", Content: "woohhooooooo"}
	articles = append(articles, newArticle)
	fmt.Fprintf(rw, "Artcle Added")
}
func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homepage Endpoint Hit\n")
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", myRouter))

}
func main() {
	handleRequest()
}
