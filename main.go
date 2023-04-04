package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//soon TM
	fmt.Println("What version of paper do you want to install? (E.G: 1.16.5)")
	var version string
	fmt.Scanln(&version)
	install("paper", version)
}

func install(software, version string) {
	fmt.Println("Where do you want to install your server? (E.G: /home/username/minecraft) (This will be created if it doesn't exist)")
	var path string
	fmt.Scanln(&path)
	fmt.Println("Installing " + software + " " + version)
	if software == "paper" || software == "Paper" {
		if version == "latest" {
			fmt.Println("Sorry, tag latest is not supported yet.")
		} else {
			fmt.Println("Downloading Paper " + version + "...")
			// Download Paper

			//get ver builds
			buildsUrl := "https://api.papermc.io/v2/projects/paper/versions/" + version
			Bres, Berr := http.Get(buildsUrl)
			if Bres.StatusCode == http.StatusOK {
				BresBytes, err := io.ReadAll(Bres.Body)
				if Berr != nil {
					log.Fatal(err)
					os.Exit(1)
				}
				defer Bres.Body.Close()
				fmt.Println(BresBytes)
				BresString := string(BresBytes)
				PbuildVer := getBuildVer(BresString)
				fmt.Println("Using build version: ", PbuildVer)

				// //get build url
				buildsUrl := "https://api.papermc.io/v2/projects/paper/versions/" + version + "/builds/" + PbuildVer
				// BuildsRes, BuildsErr := http.Get(buildsUrl)
				downloadJar(buildsUrl)
				fmt.Println("Downloaded Paper " + version + " successfully!")
				fmt.Println("Installed Paper " + version + " successfully!")
				// Install Paper
			}
		}
	}
}

func getBuildVer(jsonIn string) string {
	jsonData := []byte(jsonIn)

	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}

	builds := data["builds"].([]interface{})
	max := builds[0].(float64)
	for _, build := range builds {
		if build.(float64) > max {
			max = build.(float64)
		}
	}
	fmt.Println("Using build version: ", max)
	var float_to_int int = int(max)
	dataOut := strconv.Itoa(float_to_int)
	// dataOut := strconv.FormatFloat(max, 'E', -1, 64)
	return dataOut
}

func downloadJar(url string) {
	err := DownloadFile("paper.jar", url)
	if err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
