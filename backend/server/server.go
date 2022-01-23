package server

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

// https://gobyexample.com/http-servers
func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	fmt.Fprintf(w, "hello\n")
}

func getComments(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	queryResult := GetAllComments()
	fmt.Fprintf(w, queryResult + "\n")
}

type CommentData struct {
	Comment string	`json:"comment"`
}

func makeComment(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if req.Method == "OPTIONS" {
		return
	}

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	text, _ := io.ReadAll(req.Body)
	var comment CommentData;
	json.Unmarshal(text, &comment)

	MakeComment(comment.Comment)
}

func createHandlers() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/getComments", getComments)
	http.HandleFunc("/makeComment", makeComment)
}

func RunServer() {
	createHandlers()
	http.ListenAndServe("0.0.0.0:8090", nil)
}