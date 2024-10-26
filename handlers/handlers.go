package handlers

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting article...\n")
}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int  

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	resString := fmt.Sprintf("Article List (page: %d)\n", page)
	io.WriteString(w, resString)
}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}
func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "posting Nice...\n")	
}
func CommentHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "posting Comment...\n")
}