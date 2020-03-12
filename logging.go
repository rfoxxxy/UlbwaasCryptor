package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printLog(str string) {
	if !Logging {
		return
	}
	exist := checkFileLogExist()
	if exist != true {
		_, err := createLogFile()
		if err != nil {
			fmt.Println("Unable to create log file:", err)
			return
		}
	}
	existLogFile, err := ioutil.ReadFile(FileLogName)
	if err != nil {
		fmt.Println("Unable to open log file:", err)
		return
	}
	logfile, err := os.Create(FileLogName)
	logString := fmt.Sprintf("%v\n%v", string(existLogFile), str)
	if err != nil {
		fmt.Println("Unable to open log file:", err)
		return
	}
	defer logfile.Close()
	logfile.WriteString(logString)
	return
}

func checkFileLogExist() bool {
	_, err := ioutil.ReadFile(FileLogName)
	if err != nil {
		return false
	} else {
		return true
	}
}

func createLogFile() (bool, error) {
	file, err := os.Create(FileLogName)
	if err != nil {
		return false, err
	} else {
		defer file.Close()
		file.WriteString("log")
		return true, nil
	}
}
