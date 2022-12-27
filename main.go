package main


import (
	"fmt"
	"log"
	"net/http"

)


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	
	if err:= http.ListenAndServe(":8080", nil); err!=nil{

		log.Fatal(err)
	}
	fmt.Println("Hello World")
}


func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet{
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello !")
}

func formHandler (w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err!= nil{
		fmt.Fprintf(w, "error in ParseForm() : %v", err)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name is  %v \n", name)
	fmt.Fprintf(w, "Address is  %v \n", address)	
}
