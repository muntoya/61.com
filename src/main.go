package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func hi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range(r.Form) {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintln(w, "hello")
}

func main() {
	http.HandleFunc("/", hi)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}