package main

import(
	"fmt"
	"net/http"
	"log"
)

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		log.Printf("%s request %s", req.RemoteAddr, req.URL)
		h.ServeHTTP(res, req)
	})
}

func main(){

	h := http.NewServeMux() //http.Handler (interface)

	h.HandleFunc("/home", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "This is home handler")
	})

	h.HandleFunc("/buy", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "This is buy handler")
	})

	h.HandleFunc("/checkout", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "This is checkout handler")
	})

	h.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "This is index handler")
	})


	//logger middleware
	h1 := logRequest(h)

	log.Println("Server running on port 5000")

	//
	err := http.ListenAndServe(":5000",h1)
	if err!= nil{
		log.Fatal(err)
	}

}