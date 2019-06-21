package main

import (
	"net/http"
)

func initHandle(w http.ResponseWriter, r *http.Request) {
	
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	if ( !Installed ) {
		// send to initial handle
		initHandle(w,r)
	} 
}