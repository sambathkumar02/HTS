package Logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	LogFilePath string
	LogFile     os.File
	//count       int
	LogQueue []string
}

//Method for creating a Log file in append mode and return the pointer
func (logger Logger) CreateLogFile() os.File {
	//Open file in append mode
	file, err := os.OpenFile("request.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("Unable to open Log file!")
	}
	return *file

}

//Method to be Called from HTS to Log a Request
func (logger Logger) Log(LogData string) {

	//set output of log to created file
	log.SetOutput(&logger.LogFile)
	//Append log to file
	log.Print(LogData)

	fmt.Println(LogData)
}
