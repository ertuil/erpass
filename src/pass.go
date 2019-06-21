package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"io"
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