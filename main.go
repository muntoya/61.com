package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	//"61.com/utils"
)


func hi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Method)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["user"])
	for k, v := range(r.Form) {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	html := `<html>
	<head>
	<title></title>
	</head>
	<body>
	<form action="/login" method="post">
	    用户名:<input type="text" name="username">
	    密码:<input type="password" name="password">
	    <input type="submit" value="登陆">
	</form>
	</body>
	</html>`
	fmt.Fprintln(w, html)
}

func main() {

	http.HandleFunc("/", hi)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
