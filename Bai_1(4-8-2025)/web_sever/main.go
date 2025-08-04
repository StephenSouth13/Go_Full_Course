package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	fileSever := http.FileServer(http.Dir("./web_sever"))
	http.Handle("/", fileSever)
	http.HandleFunc("/form",from)
	http.HandleFunc("/hello", helloHandeler)	
	fmt.Println("Server is running on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		
	}

	}
