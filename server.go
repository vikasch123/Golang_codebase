package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r* http.Request){
		fmt.Fprintf(w,"Bhai, tera server chal raha hai!")
	})


http.HandleFunc("/about",func(w http.ResponseWriter , r* http.Request){
	fmt.Fprintf(w,"Bhai, yeh mera server hai aur tune /about route use kiya hai")
})


http.HandleFunc("/query",func(w http.ResponseWriter, r* http.Request){
	name:= r.URL.Query().Get("name")

	if name =="" {
		name= "Bhai"
	}

	fmt.Fprintf(w,"Hello, %s",name)
})


	http.ListenAndServe(":8080",nil)
}
