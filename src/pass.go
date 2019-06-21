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
		log.Println("[error]:Can not read file %v",filename)
	}
	return string(b)
}