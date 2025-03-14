package logger

import (
	"log"
	"os"
	"fmt"
)

// Setup logger
var logFile *os.File
var err error
var Logger *log.Logger

func init(){
	logFile, err = os.OpenFile("logs.log", os.O_RDWR | os.O_APPEND, 0600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	Logger = log.New(logFile,"", log.LstdFlags)
}