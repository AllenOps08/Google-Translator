package main 

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter , r *http.Request){
	if err != r.ParseForm(); err != nil {
		fmt.Fprint(w, "Parseform() err: %v" , err)
		return
	}
	fmt.Fprintf(w , "POST request succesful")

    name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w , "Name = %s\n" , name)
	fmt.Fprintf(w , "Address = %s\n" , address) 
}


func helloHandler( w http.ResponseWriter , r *http.Request){
  
  if r.URL.Path != "/hello" {
	http.Error(w , "Method is not supported" , http.StatusNotFound)
	
  }

}






func main() {
	
}