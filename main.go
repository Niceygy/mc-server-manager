package main

import (
	"fmt"
)

func main() {
	//soon TM
	fmt.Println("What version of paper do you want to install? (E.G: 1.16.5)")
	var version string
	fmt.Scanln(&version)
	fmt.Println("What software do you want to install? (E.G: paper, fabric)")
	var software string
	fmt.Scanln(&software)
	install(software, version)
}

func install(software, version string) {
	fmt.Println("Where do you want to install your server? (E.G: /home/username/minecraft) (This will be created if it doesn't exist)")
	var path string
	fmt.Scanln(&path)
	fmt.Println("Installing " + software + " " + version)
	if software == "paper" || software == "Paper" {
		paperInstall(version)
	} else if software == "fabric" || software == "Fabric" {
		fabricInstall(version)
	} else {
		fmt.Println("Sorry, " + software + " is not supported yet.")
	}
}
