package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"flag"
	"net/http"
	"erpass/src/data"
)
var (
	//Installed if the erpass.key is existed.
	Installed bool
	ip string
)

func initServer() {

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
	
	if err := data.RestoreAssets("./", "ReadMe.md"); err != nil {
		log.Println("[error]: Resotre readme files failed.")
    }
}

func getUsage(){
	st := `Erpass: A simple but security password generator.
Usage: erpass [-h][-d true][-host <host>][-port <port>][-log <logfile>][-g account]

	--d    run app as a daemon with -d=true or -d true.
	--host string
		  HTTP listen IP address. (default "127.0.0.1")
	--log string
		  Log file (default "erpass.log")
	--port string
		  HTTP listen port (default "8080")
	-g  	generate a password in terminal.
			Example:  erpass -g github.com
	-h/--help Show this help.
	`
	fmt.Println(st)
}

func parserCommand() (*string,*string,*string, bool,*string) {
	flag.Usage = getUsage
	
	host := flag.String("host", "127.0.0.1", "HTTP listen IP address.")
	port := flag.String("port", "8080", "HTTP listen port")
	log := flag.String("log", "erpass.log", "Log file")
	daemon := flag.Bool("d", false, "run app as a daemon with -d=true or -d true.")

	ac := flag.String("g", "", "Generate the password in CLI, follewed the account")

	if !flag.Parsed() {
		flag.Parse()
	}

	return host,port,log,*daemon,ac
}

func setLogFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
    	f = os.Stdout
	}
	log.SetOutput(f)
	// Set log format
	log.SetFlags(log.Ldate)
    return f
}

func main()  {
	host,port,logfile,daemon,ac := parserCommand()

	if *ac != "" {
		generateCLI(*ac)
		os.Exit(0)
	}

	f := setLogFile(*logfile)
	defer f.Close()

	if daemon == true {
		cmd := exec.Command(os.Args[0], flag.Args()[1:]...)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		daemon = false
		os.Exit(0)
	}

	initServer()

	staticFileHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileHandle))
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/doc", docHandle)
	http.HandleFunc("/gen", generateSecretKeyHandle)
	http.HandleFunc("/import", importHandle)
	http.HandleFunc("/pass", passHandle)
	log.Printf("[info]: erpass server starting. Listening %s:%s\n",*host,*port);
    if err := http.ListenAndServe(*host+":"+*port, nil); err != nil {
        log.Printf("[error]: erpass server failed to start.")
    }
}