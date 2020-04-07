package main;

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

func handleRequests(){
	http.HandleFunc("/", homePage);
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests();
}