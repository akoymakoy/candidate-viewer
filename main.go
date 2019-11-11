// Package main starts the simple server on port and serves HTML,
// CSS, and JavaScript to clients.
package main

import (
	"candidate-viewer/candidateutil"
	"candidate-viewer/database"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// templates parses the specified templates and caches the parsed results
// to help speed up response times.
var templates = template.Must(template.ParseFiles("./templates/base.html", "./templates/body.html"))

// logging is middleware for wrapping any handler we want to track response
// times for and to see what resources are requested.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

// index is the handler responsible for rending the index page for the site.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int
		var err error
		if r.Method == http.MethodPost {
			id,err = strconv.Atoi(r.FormValue("Id"))
			log.Print("POST")
		}else{
			id =1
			log.Print("REQUEST INDEX")
		}


		db,err := database.Connect()
		c := candidateutil.Get(db,id)

		b := struct {
			Title        template.HTML
			Id			 int
			Name,Email   string
		}{
			Title:        template.HTML("Business &verbar; Landing"),
			Id: 		c.Id,
			Name:       c.Name,
			Email:       c.Email,
		}
		templates.ExecuteTemplate(w, "base", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// public serves static assets such as CSS and JavaScript to clients.
func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

//candidate view
func profilepage(design int) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int
		var err error
		if r.Method == http.MethodPost {
			id,err = strconv.Atoi(r.FormValue("Id"))
			log.Print("POST")
		}else{
			id =1
			log.Print("REQUEST VIEW1")
		}
		db,err := database.Connect()
		c := candidateutil.Get(db,id)

		b := struct {
			Title        template.HTML
			Id			 int
			Name,Email   string
		}{
			Title:        template.HTML("Business &verbar; Landing"),
			Id: 		c.Id,
			Name:       c.Name,
			Email:       c.Email,
		}
		if design == 1{
			log.Print("GOING TO TEMPLATE 1")
			templates.ExecuteTemplate(w, "template1", &b)
		}else{
			log.Print("GOING TO TEMPLATE 2")
			templates.ExecuteTemplate(w, "template2", &b)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/public/", logging(public()))
	mux.Handle("/view1/", logging(profilepage(1)))
	mux.Handle("/view2/", logging(profilepage(2)))
	//mux.Handle("/", logging(index()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	http.ListenAndServe(addr, mux)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start simple server: %v\n", err)
	}


}
