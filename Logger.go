package main

import "time"

type Logger struct {
	TempDict    map[time.Time][]string
	LogFilePath string
	//count       int
}

func (logger Logger) Log() {

}
