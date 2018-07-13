package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// FOR LATER USE

// type execute struct {
// 	CommandName string
// }

// type writeExecute struct {
// 	CommandName   string
// 	ActualCommand string
// }

// type readRead struct {
// 	CommandName string
// }

// type writeRead struct {
// 	CommandName   string
// 	ActualCommand string
// }

func splitAndRunCommand(commandName string) {

	commandArray := strings.Split(commandName, " ")
	cmd := exec.Command(commandArray[0], commandArray[1:]...)

	if commandArray[0] == "sh" {
		output, err := cmd.Output()
		checkError(err)
		fmt.Println(string(output))
	} else {
		err := cmd.Run()
		if err != nil {
			fmt.Println("error: ")
			fmt.Println(err)
		}
	}
}

// FOR ENUM USE IOTA, using commandtype string for now.

// FOR LATER USE

// func addNewCommand(commandName string, actualCommand string) {
// 	session, err := mgo.Dial("127.0.0.1:27017")
// 	checkError(err)

// 	// If adding an execute command.
// 	mainC := strings.Split(commandName, " ")

// 	if mainC[0] == "execute" {
// 		currentWorker := session.DB("test1").C("execute1")

// 		commands := writeExecute{CommandName: commandName, ActualCommand: actualCommand}

// 		err := currentWorker.Insert(&commands)
// 		checkError(err)

// 		fmt.Println("Command: " + actualCommand + "inserted, mapped as: " + commandName)

// 	}
// }
