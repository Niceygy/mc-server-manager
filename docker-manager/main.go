package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	log.Println("Started!")
	var serverVer = getServerVersion()
	var serverSoft = getServerSoftware()
	if serverSoft == "paper" {
		paperInstall(serverVer)
	} else if serverSoft == "fabric" {
		fabricInstall(serverVer)
	} else {
		log.Println("Sorry, this server software is not supported yet.")
	}
}

func getServerVersion() string {
	f, err := os.Open("/var/mc/server/serverVersion")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func getServerSoftware() string {
	f, err := os.Open("/var/mc/server/serverSoftware")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func paperInstall(version string) {

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
			buildsUrl := "https://api.papermc.io/v2/projects/paper/versions/" + version + "/builds/" + PbuildVer + "/downloads/paper-" + version + "-" + PbuildVer + ".jar"
			// BuildsRes, BuildsErr := http.Get(buildsUrl)
			downloadJar(buildsUrl)
			fmt.Println("Downloaded Paper " + version + " successfully!")
			fmt.Println("Installed Paper " + version + " successfully!")
			// Install Paper
		}
	}
}

func fabricInstall(version string) {
	fmt.Println("Using default fabric loader (0.14.9) & installer version (0.11.2). If you want to use a different version, please use the fabric installer.")
	fmt.Println("Downloading Fabric " + version + "...")
	// Download Fabric
	FabricServerCmd := "https://meta.fabricmc.net/v2/versions/loader/" + version + "/0.14.19/0.11.2/server/jar"
	downloadJar(FabricServerCmd)

}

func downloadJar(url string) {
	err := DownloadFile("server.jar", url)
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
