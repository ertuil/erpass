package main

import (
	"log"
	"erpass/src/data"
)

func init() {
	log.SetFlags(log.Ldate|log.Lshortfile)
}

func restore() {
	if err := data.RestoreAssets("./", "static"); err != nil {
		log.Println("[error]: Resotre static files failed.");
    }
}

func main()  {
	log.Println("[info]: erpass server started");
	restore();
}