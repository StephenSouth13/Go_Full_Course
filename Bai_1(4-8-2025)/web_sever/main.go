package main

import (
	"fmt"
	"log"
	"net/http"
)
func fromHandler(w http.ResponseWriter, r *http.Request) {
	if err:= r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "Post request successful\n")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)
	address := r.FormValue("address")
	fmt.Fprintf(w, "Address = %s\n", address)
}
func helloHandeler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method !="Get" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, this is a web server using Go Lang")
}
func main() {
	fileSever := http.FileServer(http.Dir("./web_sever"))
	http.Handle("/", fileSever)
	http.HandleFunc("/form",fromHandler)
	http.HandleFunc("/hello", helloHandeler)	
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		
	}

	}
//D:\Code linh tinh\Go_Lang\Go_Full_Course\Bai_1(4-8-2025)\web_sever