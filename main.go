package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type JSON struct {
	Title  string
	Auteur string
	About  string
	Choix  string
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("template/*.html")
	port := os.Getenv("PORT")
	//the diffenrent possible path
	http.HandleFunc("/", getFormeHandler)
	http.Handle("/public/", http.FileServer(http.Dir(".")))
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// http.ListenAndServe(":8080", nil)
}
func getFormeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("In Home")
	// jsonfile, _ := ioutil.ReadFile("language.json")

	// var data []JSON
	// // unmarshall (tres sombre --> aller voir)
	// json.Unmarshal(jsonfile, &data)

	// choice := r.FormValue("Submit")
	// if choice == "choix 2" {
	// 	tpl.ExecuteTemplate(w, "index.html", data[1])
	// } else {
	// 	tpl.ExecuteTemplate(w, "index.html", data[0])
	// }
	tpl.ExecuteTemplate(w, "index.html", nil)
}
