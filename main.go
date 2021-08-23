package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sambathkumar02/HTS/HTS"
	"github.com/sambathkumar02/HTS/Logger"
)

var LoggerObj = Logger.Logger{}

func main() {

	//Handling Command Line Flags Port,Homedirectory
	port := flag.String("p", "", "Port Number")
	homedir := flag.String("d", "", "Home Directory")
	//logfile := flag.String("l", "", "Log File Path")
	flag.Parse()

	//Printing Default message when port and directory not specified
	if *port == "" && *homedir == "" {
		fmt.Printf("HTS-HTTP Static Server \n OPTIONS: \n -d : Root Directory \n -p :Port Number \n USAGE: \n hts -p 80 -d /home/usernmae \n")
		os.Exit(0)
	} else if *port == "" && *homedir != "" {
		*port = "80"
		fmt.Printf("[+]Port Number Not Mentioned - Use Default Port 80!")
	} else if *port != "" && *homedir == "" {

		fmt.Printf("[+]Root Directory not Mentioned!")
		os.Exit(0)
	}

	//concatinate the connection string
	ConnectionString := ":" + *port

	//Server Multiplexer for Handling Conenctions
	mux := http.NewServeMux()

	//Configurations for Logging
	LogFileName := "request.log"
	LoggerObj.LogFilePath = LogFileName
	file, err := LoggerObj.CreateLogFile()
	if err != nil {
		fmt.Print("Opening Log File Failed!")
		os.Exit(0)
	}
	defer file.Close()
	LoggerObj.LogFile = file

	//Creating HTS server object
	hts := HTS.HTS{HomeDir: *homedir, Port: *port, LoggerObject: LoggerObj}
	hts.ParseConfig()

	//Only Handler for get all Paths

	go mux.HandleFunc("/", hts.HandleHome)
	fmt.Print("Server Listening !\nRoot Directory:", *homedir)

	//Listen and server
	log.Fatal(http.ListenAndServe(ConnectionString, mux))

}
