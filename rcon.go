package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func rcon(port, password, command string) {
	fmt.Println("Connecting...")
	initRcon := exec.Command("mcutils rcon localhost" + port + " " + password + " " + command)
	initRcon.Stdout = os.Stdout
	initRcon.Stderr = os.Stderr
	err := initRcon.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(initRcon.Stdout)
}
