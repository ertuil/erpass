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
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

const(
	filename = "erpass.key"
)

func generateSecretKey() bool {
	t := time.Now().UnixNano()
	h := sha256.New()
	io.WriteString(h, strconv.FormatInt(t, 10)) 
	token := fmt.Sprintf("%x", h.Sum(nil)) 
	log.Println("[warning]:Gernerate a secret key:",token);
	return writeSecretKey(token)
}

func importSecretKey(sk string) bool {
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
	return true;
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

func generatePassword(info map[string]string) string {
    pwd := []byte(info["mk"] + info["account"] + info["count"])
    salt := []byte(readSecretKey())
    iterations := 15000
    digest := sha256.New

    dk := pbkdf2.Key(pwd, salt, iterations, 64, digest)

	str := base64.URLEncoding.EncodeToString(dk)
	
    return str
}