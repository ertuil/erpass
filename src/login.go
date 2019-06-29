package main

import (
	"net/http"
	"strconv"
)

var (
	maxageTable = []int{1800, 3600, 10800, 43200, 86400, 
		259200, 605800, 1296000, 2592000, 5184000,
		10368000, 31536000, 3122064000}
)

func newlogHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sk := r.Form["sk"][0]
	id, err := strconv.Atoi(r.Form["maxage"][0])
	if err != nil {
		id = 0
	}
	maxage := maxageTable[ id ]
	if checkSecretKey(sk) {
		writeSecretKeyFromCookie(w,r,sk,maxage)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func newgenHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(generateSecretKey()))
}

func newExitHandle(w http.ResponseWriter, r *http.Request) {
	deleteSecretKeyFromCookie(w,r)
	http.Redirect(w, r, "/", http.StatusFound)
}