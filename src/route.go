package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
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
	sk := generateSecretKey()
	ret := writeSecretKey(sk)

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
	ret := importSecretKey(secretKey)

	if (ret == true) {
		Installed = true
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func passHandle(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		r.Response.StatusCode = 500;
		return 
	}
	js := make(map[string]string)
	json.Unmarshal(result,&js)
	sk := readSecretKeyFromCookie(w,r)
	w.Write([]byte(generatePassword(js, sk)))
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.html",
	"templates/public/header.html",
	"templates/public/title.html", 
	"templates/public/footer.html")

	if err := t.Execute(w,map[string]string{});err != nil {
		log.Println("[error]: parser template error in initHandle")
	}
}

// Index page
func rootHandle(w http.ResponseWriter, r *http.Request) {
	// if ( !Installed ) {
	// 	// send to initial handle
	// 	initHandle(w,r)
	// } else {
	// 	indexHandle(w,r)
	// }
	ret := isExistSecretKeyFromCookie(w,r)
	if ret == false {
		loginHandle(w,r)
	} else {
		indexHandle(w,r)
	}
}

func docHandle(w http.ResponseWriter, r *http.Request) {
	rd,err := ioutil.ReadFile("./ReadMe.md")
	if err != nil {
		log.Println("[error]: Can not read ReadMe.md in docHandle")
	}

	w.Write(rd);
}