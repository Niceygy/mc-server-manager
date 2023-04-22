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

func startServer(name, ram string) {
	fmt.Println("Starting server...")
	subProcess := exec.Command("bash -c java -Xmx" + ram + "-jar " + name + ".jar nogui")
	out, err := subProcess.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
