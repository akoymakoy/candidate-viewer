package main


import (
    "html/template"
    "log"
	"net/http"
	"time"
    
     "github.com/gorilla/mux"
)

func parseTpl(w http.ResponseWriter, r *http.Request) {
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
		"templates/template1.tmpl",
	}

	contactTmpl, _ := template.ParseFiles(paths...)

	err := contactTmpl.ExecuteTemplate(w, "layout", b)
	if err != nil {
		panic(err)
	}
}

func main() {
	
	 r := mux.NewRouter()

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/assets/styles"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/assets/scripts"))))

	r.HandleFunc("/", parseTpl)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}