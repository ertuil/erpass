package main

import (
	"log"
	"net/http"
	"html/template"
)

// The Handle for initial page
func initHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/install.html",
	"templates/public/header.html",
	"templates/public/title.html", 
	"templates/public/footer.html")

	if err := t.Execute(w,map[string]string{});err != nil {
		log.Println("[error]: parser template error in initHandle")
	}
}

// Main page
func indexHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html",
								"templates/public/header.html",
								"templates/public/title.html", 
								"templates/public/footer.html")

	if err := t.Execute(w,map[string]string{});err != nil {
		log.Println("[error]: parser template error in initHandle")
	}
}

// generate a new secret key page
func generateSecretKeyHandle(w http.ResponseWriter, r *http.Request) {
	if(Installed)  {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	ret := generateSecretKey();
	if (ret == true) {
		Installed = true
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// import a key handle
func importHandle(w http.ResponseWriter, r *http.Request) {
	if(Installed)  {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	
	r.ParseForm()
	secretKey := r.Form["secret"][0]
	ret := writeSecretKey(secretKey)

	if (ret == true) {
		Installed = true
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Index page
func rootHandle(w http.ResponseWriter, r *http.Request) {
	if ( !Installed ) {
		// send to initial handle
		initHandle(w,r)
	} else {
		indexHandle(w,r)
	}
}