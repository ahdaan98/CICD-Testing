package main

import (
	"fmt"
	"net/http"

	"github.com/enescakir/emoji"
)

const PORT = ":3000"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("CICD Testing %v",emoji.Popcorn)))
}

func main(){
	http.HandleFunc("/",HelloHandler)
	http.ListenAndServe(PORT,nil)
}