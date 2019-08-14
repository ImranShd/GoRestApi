package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Homepage")
	fmt.Println("Endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Articles All")

	json.NewEncoder(w).Encode(articles)

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	json.Unmarshal(reqBody, &article)

	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
	//hrllocoocmw

}

func deleteArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: delete Article")
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}

}

func updateArticles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	for index, art := range articles {
		if art.Id == id {

			reqBody, _ := ioutil.ReadAll(r.Body)
			fmt.Fprintf(w, "%+v", string(reqBody))
			fmt.Println("Endpoint hit: update Article")
			var article Article
			json.Unmarshal(reqBody, &article)

			articles[index] = article
			json.NewEncoder(w).Encode(article)
		}
	}

}

func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticles)
	myRouter.HandleFunc("/article/{id}", deleteArticles).Methods("DELETE")
	myRouter.HandleFunc("/articleu/{id}", updateArticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	a := Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"}
	b := Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"}

	articles = append(articles, a)
	articles = append(articles, b)

	handleRequests()
}
