package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ss(res http.ResponseWriter, req *http.Request) {
	io.Copy(res, strings.NewReader("hello"))
}

func main() {
	fs := http.FileServer(http.Dir("./assets"))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			http.ServeFile(res, req, "./assets/index.html")
			return
		}
		fs.ServeHTTP(res, req)
	})
	http.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	}))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	wd, _ := os.Getwd()
	log.Println("Working directory:", wd)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
