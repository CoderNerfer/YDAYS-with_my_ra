package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// JSON sstruct pour recup les element du json
type JSON struct {
	Title          string
	Link1          string
	Link2          string
	Link3          string
	Link4          string
	Link5          string
	Home           string
	Email          string
	Facebook       string
	Popup          string
	Projet         string
	Prensation     string
	Span1          string
	Span2          string
	Span3          string
	Span4          string
	About1         string
	About2         string
	About3         string
	About4         string
	About5         string
	About6         string
	About7         string
	About8         string
	About9         string
	About10        string
	About11        string
	Classinfo      string
	Classanim      string
	Classchefe     string
	Classestratcom string
	Classda        string
	Clem           string
	Mathis         string
	Steph          string
	Lilian         string
	Ju             string
	Lele           string
	Mel1           string
	Mel2           string
	Lou1           string
	Lou2           string
	Bd             string
	Titre          string
	Archive        string
	Resume1        string
	Resume2        string
	Resume3        string
	Resume4        string
	Resume5        string
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("template/*.html")
	port := os.Getenv("PORT")
	// the diffenrent possible path
	http.HandleFunc("/", getFormeHandler)
	http.HandleFunc("/mentions_legales", getMentionLegaleHandler)
	http.Handle("/public/", http.FileServer(http.Dir(".")))
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// http.ListenAndServe(":8080", nil)
}
func getFormeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Home")
	jsonfile, _ := ioutil.ReadFile("language.json")

	var data []JSON
	// unmarshall (tres sombre --> aller voir)
	json.Unmarshal(jsonfile, &data)
	// fmt.Println(data)
	choice := r.FormValue("submit")
	if choice == "EN" {
		tpl.ExecuteTemplate(w, "index.html", data[1])
	} else {
		tpl.ExecuteTemplate(w, "index.html", data[0])
	}
	// fmt.Println(r.FormValue("submit"))
	// tpl.ExecuteTemplate(w, "index.html", nil)
}

func getMentionLegaleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Mentions Legales")

	tpl.ExecuteTemplate(w, "mentions_legales.html", nil)
}
