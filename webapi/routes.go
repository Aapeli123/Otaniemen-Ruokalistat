package webapi

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"otaniemenruokalistat.tk/ruokalista"
)

var tmpl *template.Template

// Init loads templates and starts the webserver
func Init() {
	tmpl = template.Must(template.ParseFiles("./templates/index.go.tpl"))
	http.HandleFunc("/index", handleHTMLReq)
	http.HandleFunc("/", handleHTMLReq)

	http.HandleFunc("/json", handleJSONReq)
	fmt.Println("Server started...")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func handleJSONReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	foods, err := ruokalista.GetThisWeeksFood()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	enc := json.NewEncoder(w)
	enc.Encode(foods)
}

type data struct {
	Paivat []ruokalista.Päivä
}

func handleHTMLReq(w http.ResponseWriter, r *http.Request) {
	foods, err := ruokalista.GetThisWeeksFood()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	tmpl.Execute(w, data{Paivat: foods})
}
