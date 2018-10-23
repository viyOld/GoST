package handlers

import (
	//"fmt"
	"html/template"
	"net/http"
	//"log"
	log "github.com/sirupsen/logrus"
)

//StartPage GET start page
func StartPage(w http.ResponseWriter, r *http.Request) {
	//logrus.Println("StartPage GET start page +++")
	log.Println("StartPage GET start page +++")
	// person := Person{ID: "1", Name: "Foo"}
	parsedTemplate, err := template.ParseFiles("web/static/index.html", "web/static/header.html",
		"web/static/footer.html", "web/static/nav.html")
	if err != nil {
		log.Println("I don`t parse static files static/index.html")
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error arsedTemplate.Execute in StartPage : ", err)
		return
	}

	//fmt.Println("Start Page: ")
	//w.Write([]byte("hello"))

}
