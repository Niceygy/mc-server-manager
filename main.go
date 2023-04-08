package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "install" {
		install()
	} else if os.Args[1] == "run" {
		fmt.Println("Starting server...")
		startServer("./server", "2G")
	}
}

func install() {
	fmt.Println("What version of paper do you want to install? (E.G: 1.16.5)")
	var version string
	fmt.Scanln(&version)
	fmt.Println("What software do you want to install? (E.G: paper, fabric)")
	var software string
	fmt.Scanln(&software)
	// fmt.Println("Where do you want to install your server? (E.G: /home/username/minecraft) (This will be created if it doesn't exist)")
	// var path string
	// fmt.Scanln(&path)
	fmt.Println("Installing " + software + " " + version)
	if software == "paper" || software == "Paper" {
		paperInstall(version)
	} else if software == "fabric" || software == "Fabric" {
		fabricInstall(version)
	} else {
		fmt.Println("Sorry, " + software + " is not supported yet.")
	}
}
