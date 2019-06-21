package main

import (
	"log"
	"os"
	"flag"
	"net/http"
	"erpass/src/data"
)
var (
	//Installed if the erpass.key is existed.
	Installed bool
	ip string
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
		Installed = true
	} else {
		log.Println("[warning]: erpass.key is not installed.")
		Installed = false
	}
}

func restore() {
	if err := data.RestoreAssets("./", "static"); err != nil {
		log.Println("[error]: Resotre static files failed.")
	}

	if err := data.RestoreAssets("./", "templates"); err != nil {
		log.Println("[error]: Resotre template files failed.")
    }
}

func parserCommand() (*string,*string) {
	host := flag.String("host", "127.0.0.1", "HTTP listen IP address.")
	port := flag.String("port", "8080", "HTTP listen port")
	
	flag.Parse()

	return host,port
}

func main()  {
	host,port := parserCommand()

	staticFileHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileHandle))
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/gen", generateSecretKeyHandle)
	http.HandleFunc("/import", importHandle)
	log.Printf("[info]: erpass server starting. Listening %s:%s\n",*host,*port);
    if err := http.ListenAndServe(*host+":"+*port, nil); err != nil {
        log.Printf("[error]: erpass server failed to start.")
    }
}