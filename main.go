package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
	        <!DOCTYPE html>
	        <html>
	        <head>
	            <title>My Go Web Server</title>
	            <style>
				body {
					background-color: navy;
					color: white;
					font-family: Georgia, serif;
					padding: 20px;
					text-align: center;
				}
			</style>
		</head>
		<body>
		<h1> "Hello from Go Web server!" </h1>
		</body>
		</html>
	`)
	// fmt.Fprintf(w, "Hello from Go Web server!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
