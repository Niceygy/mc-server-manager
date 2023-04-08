package main

import (
	"fmt"
	"io"
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
	subProcess := exec.Command("java -Xmx" + ram + " -jar " + name + ".jar")
	//subProcess := exec.Command("go", "run", "./helper/main.go")
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()
	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	fmt.Println("Server Starting...")
	if err = subProcess.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	} else {
		fmt.Println(subProcess.Stdout)
	}

	io.WriteString(stdin, "4\n")
	subProcess.Wait()
	fmt.Println("END")
}
