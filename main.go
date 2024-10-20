package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}
	postArticleHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting article...\n")
	}
	ArticleListHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "article List\n")
	}
	OneArticleHandler := func(w http.ResponseWriter, r *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("article No.%d\n", articleID)
		io.WriteString(w, resString)
	}
	ArticleNiceHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "posting Nice...\n")	
	}
	CommentHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", ArticleListHandler)
	http.HandleFunc("/article/1", OneArticleHandler)
	http.HandleFunc("/article/nice", ArticleNiceHandler)
	http.HandleFunc("/comment", CommentHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

