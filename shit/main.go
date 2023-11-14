package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
	"strconv"
	"math"
)

type Article struct {
	Response float64 `json:"Response"`
}

func Radius(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key,_ := strconv.ParseFloat(vars["radius"], 32)
	var answer Article = Article{Response: math.Pi * float64(key) * float64(key)}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(answer)
}

func main(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/r/{radius}", Radius)
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}