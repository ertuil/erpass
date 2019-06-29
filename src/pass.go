package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"io"
	"strings"
	"io/ioutil"  
	"strconv"
	"net/http"
	"crypto/sha256"
	"encoding/binary"
	"golang.org/x/crypto/pbkdf2"
)

const(
	filename = "erpass.key"
)

func isExistSecretKeyFromCookie(w http.ResponseWriter,r *http.Request) bool{
	_, err := r.Cookie("sk")
	if err != nil {
		return false
	}
	return true
}

func readSecretKeyFromCookie(w http.ResponseWriter,r *http.Request) string{
	cookie, err := r.Cookie("sk")
	if err != nil {
		return ""
	} 
	return cookie.Value
}

func writeSecretKeyFromCookie(w http.ResponseWriter,r *http.Request, sk string, maxage  int) {
	cookie := http.Cookie{Name: "sk", Value: sk, Path: "/", MaxAge: maxage}
    http.SetCookie(w, &cookie)
}

func deleteSecretKeyFromCookie(w http.ResponseWriter,r *http.Request) {
	cookie := http.Cookie{Name: "sk", Path: "/", MaxAge: -1}
    http.SetCookie(w, &cookie)
}

func generateSecretKey() string {
	t := time.Now().UnixNano()
	h := sha256.New()
	io.WriteString(h, strconv.FormatInt(t, 10)) 
	token := fmt.Sprintf("%x", h.Sum(nil)) 
	log.Println("[warning]:Gernerate a secret key:",token);
	return token
}

func  checkSecretKey(sk string) bool {
	sk = strings.ToLower(sk)
	if len(sk) != 64 {
		return false
	} 
	for _,c := range(sk) {
		if  ((c >= '0'  && c <='9') ||  ( c >= 'a' && c <='f')) {
			continue
		} else {
			return false
		}
	}
	return true
}

func importSecretKey(sk string) bool {
	sk = strings.ToLower(sk)
	ret := checkSecretKey(sk)
	if ret == false {
		return false 
	}
	return writeSecretKey(sk);
}

func writeSecretKey(sk string) bool {

	file, err := os.Create(filename);

	if err != nil {
		log.Println("[error]:Can not open file %v",filename)
		return false
	}
	if _, err := io.WriteString(file,sk); err != nil {
		log.Println("[error]:Can not write file %v",filename)
		return false
	}
	file.Close()
	return true
}

func readSecretKey() string{
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("[error]:Can not read file ",filename)
		return ""
	}
	return string(b)
}

func generatePassword(info map[string]string, sk string) string {

	var salt []byte
	if sk != "" {
		salt = []byte(sk)
	} else  {
		salt = []byte(readSecretKey())
	}
	
	lru := []string{
		"abcdefjhigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"abcdefjhigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{}<>,.",
		"0123456789",
		}
    pwd := []byte(info["mk"] + info["account"] + info["count"])
    iterations := 15000
    digest := sha256.New

    dk := pbkdf2.Key(pwd, salt, iterations, 64, digest)

	//str := base64.URLEncoding.EncodeToString(dk)
	st,err := strconv.Atoi(info["style"])
	if err != nil {
		st = 0
	}

	length,err := strconv.Atoi(info["length"])
	if err != nil {
		length = 16
	}

	str := remapPassword(lru[st],dk,length)
    return str
}

func remapPassword(charset string, raw []byte, length int) string {
	rawInt := make([]uint16,0)
	ansByte := make([]byte,0)

	for i := 0; i+1 < len(raw); i = i+2 {
		tmp := raw[i:i+2]
		rawInt = append(rawInt, binary.BigEndian.Uint16(tmp))
	}
	for _,cc := range(rawInt) {
		idx := int(cc) % len(charset)
		ansByte = append(ansByte, charset[idx])
	}

	return string(ansByte)[0:length]
}