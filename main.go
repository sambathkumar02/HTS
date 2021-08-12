package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

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

	//Creating HTS server object
	hts := HTS{HomeDir: *homedir, Port: *port}
	hts.ParseConfig()

	//Only Handler for get all Paths

	go mux.HandleFunc("/", hts.HandleHome)
	fmt.Print("Server Listening !\nRoot Directory:", *homedir)

	//Listen and server
	log.Fatal(http.ListenAndServe(ConnectionString, mux))

}
