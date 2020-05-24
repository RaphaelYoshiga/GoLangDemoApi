package main;

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "HOMEPAGE Endpoint hit");
}

type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
}

type Articles []Article 

func allArticles(w http.ResponseWriter, r*http.Request) {
	articles := Articles {
		Article{ Title: "test", Desc: "descirpitop"},
	}

	fmt.Println("Endpoinmt hit: all articles endpoint");
	json.NewEncoder(w).Encode(articles);
}

func postArticles(w http.ResponseWriter, r*http.Request) {
	articles := Articles {
		Article{ Title: "test", Desc: "descirpitop"},
	}

	fmt.Println("Endpoinmt hit: post article endpoint");
	json.NewEncoder(w).Encode(articles);
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true);

	myRouter.HandleFunc("/", homePage);
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main(){
	handleRequests();
}