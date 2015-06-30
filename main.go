package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
	"broker/app"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
	DEFAULT_HOST = "localhost"
)

var index = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/index.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
  index.Execute(w, nil)
}

// /user/new
func userNewHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("newUser")
	template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/user/new.html",
	)).Execute(w, nil)
}
// /user
func userEntityHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("handle")
	switch req.Method {
	case "POST":
    userName := req.FormValue("userName")
    password := req.FormValue("password")
    app.NewUser(userName, password)

    template.Must(template.ParseFiles(
      "templates/_base.html",
      "templates/user/index.html",
    )).Execute(w, app.GetUserList())
	case "GET":
		template.Must(template.ParseFiles(
			"templates/_base.html",
			"templates/user/index.html",
		)).Execute(w, app.GetUserList())
	case "PUT":
	case "DELETE":
	}
}
// /user/*
func userItemHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("handle")
	switch req.Method {
	case "POST":
	case "GET":
		template.Must(template.ParseFiles(
			"templates/_base.html",
			"templates/user/index.html",
		)).Execute(w, app.GetUserList())
	case "PUT":
	case "DELETE":
	}
}
func main() {
	var port string
	if port = os.Getenv("VCAP_APP_PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	var host string
	if host = os.Getenv("VCAP_APP_HOST"); len(host) == 0 {
		host = DEFAULT_HOST
	}

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/user", userEntityHandler)
	http.HandleFunc("/user/new", userNewHandler)
	http.HandleFunc("/user/", userItemHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Printf("Starting app on %+v:%+v\n", host, port)
	http.ListenAndServe(host+":"+port, nil)
}
