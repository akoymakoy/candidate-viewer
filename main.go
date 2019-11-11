package main


import (
	"candidate-viewer/candidateutil"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)


func CandidateProfileHandler(w http.ResponseWriter, r *http.Request) {
	b := struct {
			Title        template.HTML
			Id,Javaskill			 int
			Name,Email   string
			Candidate candidateutil.Candidate
		}{
			Title:        template.HTML("HRDB | Candidate Profile"),
			Id: 		1,
			Javaskill: 100,
			Name:       "Mark Miraflor",
			Email:       "mark.miraflor@luunax.com",
			Candidate: candidateutil.Get(1),
		}
		
	// Files are provided as a slice of strings.
	paths := []string{
		"templates/profile.tmpl",
	}

	contactTmpl, _ := template.ParseFiles(paths...)

	err := contactTmpl.ExecuteTemplate(w, "layout", b)
	if err != nil {
		panic(err)
	}
}

func CandidateListHandler(w http.ResponseWriter, r *http.Request) {
	
	b := struct {
			Title        template.HTML
			Id,Javaskill			 int
			Name,Email   string
		}{
			Title:        template.HTML("Business &verbar; Landing"),
			Id: 		1,
			Javaskill: 100,
			Name:       "Mark Miraflor",
			Email:       "mark.miraflor@luunax.com",
		}	
		
	// Files are provided as a slice of strings.
	paths := []string{
		"templates/template2.tmpl",
	}

	contactTmpl, _ := template.ParseFiles(paths...)

	err := contactTmpl.ExecuteTemplate(w, "candidatelistlayout", b)
	if err != nil {
		panic(err)
	}
}


func main() {
	
	r := mux.NewRouter()
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/assets/styles"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/assets/scripts"))))
	r.HandleFunc("/candidate/{candidateid}", CandidateProfileHandler).Methods("GET")
	r.HandleFunc("/candidatelist", CandidateListHandler).Methods("GET")
	
    http.Handle("/", r)
    
    

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}