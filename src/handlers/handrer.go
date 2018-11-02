package handlers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//StartPage GET start page
func StartPage(w http.ResponseWriter, r *http.Request) {
	log.Println("StartPage GET start page handler")
	// person := Person{ID: "1", Name: "Foo"}
	parsedTemplate, err := template.ParseFiles("web/static/index.html", "web/static/header.html",
		"web/static/footer.html", "web/static/nav.html")
	if err != nil {
		log.Println("I don`t parse static files static/index.html")
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsedTemplate.Execute in StartPage : ", err)
		return
	}

	//fmt.Println("Start Page: ")
	//w.Write([]byte("hello"))

}

//LoginPage GET login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginPage GET start page handler")
	parsedTemplate, err := template.ParseFiles("web/static/login.html", "web/static/header.html",
		"web/static/footer.html", "web/static/nav.html")
	if err != nil {
		log.Println("I don`t parse static files static/login.html")
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsedTemplate.Execute in LoginPage : ", err)
		return
	}
}

//PostLogin POST Login
func PostLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Login POST Login")
	r.ParseForm()
	nik := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("Nik = ", nik)
	log.Println("Password = ", password)
	//fmt.Println("Пользователь:", r.Form["username"])
	//fmt.Println("Пароль:", r.Form["password"])
	http.Redirect(w, r, "/", 301)
}
