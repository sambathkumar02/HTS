package Logger

import (
	"errors"
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
func (logger Logger) CreateLogFile() (os.File, error) {
	//Open file in append mode
	file, err := os.OpenFile("/etc/HTS/hts.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("Unable to open Log file!")
		return *file, errors.New("Log File Opening Failed")
	} else {

		return *file, nil

	}

}

//Method to be Called from HTS to Log a Request
func (logger Logger) Log(LogData string) {

	//set output of log to created file
	log.SetOutput(&logger.LogFile) //Keep it in same function (Scope error occours)

	//Append log to file
	log.Print(LogData)

	fmt.Print("\n", LogData)
}
