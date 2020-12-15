package main

import (
	"training-go/sesi11_assignment/controllers"
	"fmt"
	"net/http"
)

type NewMux struct{}

func (*NewMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		controllers.Index(w, r)
	case "/show":
		controllers.ViewDetail(w, r)
	case "/delete":
		controllers.DeleteUser(w, r)
	case "/add":
		switch r.Method {
		case "GET":
			controllers.ViewAddUser(w, r)
		case "POST":
			controllers.AddUser(w, r)
		}
	case "/edit":
		switch r.Method {
		case "GET":
			controllers.ViewEditUser(w, r)
		case "POST":
			controllers.EditUser(w, r)
		}
	default:
		http.NotFound(w, r)
		return
	}
}

func main() {
	mux := &NewMux{}
	fmt.Println("Listen http port :9000")
	http.ListenAndServe(":9000", mux)
}
