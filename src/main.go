package main

import (
	"log"
	"os"
	"erpass/src/data"
)
var (
	//Installed if the erpass.key is existed.
	Installed bool
)

func init() {
	// Set log format
	log.SetFlags(log.Ldate|log.Lshortfile)

	// check release
	if _,err := os.Stat("./static"); err != nil {
		log.Println("[warning]: Releasing static files.")
		restore() // release static files
	}

	if _,err := os.Stat("./erpass.key"); err == nil {
		Installed = false
	} else {
		log.Println("[warning]: erpass.key is not installed.")
		Installed = true
	}
}

func restore() {
	if err := data.RestoreAssets("./", "static"); err != nil {
		log.Println("[error]: Resotre static files failed.")
    }
}

func main()  {
	log.Println("[info]: erpass server started");
	restore();
}