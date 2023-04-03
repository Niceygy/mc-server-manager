// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"strconv"
// )

// func install(software, version string) {
// 	fmt.Println("Installing " + software + " " + version)
// 	if software == "paper" || software == "Paper" {
// 		if version == "latest" {
// 			fmt.Println("Sorry, tag latest is not supported yet.")
// 		} else {
// 			fmt.Println("Downloading Paper " + version + "...")
// 			// Download Paper

// 			//get ver builds
// 			buildsUrl := "https://api.papermc.io/v2/projects/paper/versions/" + version
// 			Bres, Berr := http.Get(buildsUrl)
// 			if Bres.StatusCode == http.StatusOK {
// 				BresBytes, err := io.ReadAll(Bres.Body)
// 				if Berr != nil {
// 					log.Fatal(err)
// 					os.Exit(1)
// 				}
// 				defer Bres.Body.Close()
// 				BresString := string(BresBytes)
// 				PbuildVer := getBuildVer(BresString)

// 				// //get build url
// 				buildsUrl := "https://api.papermc.io/v2/projects/paper/versions/" + version + "/builds/" + PbuildVer
// 				// BuildsRes, BuildsErr := http.Get(buildsUrl)
// 				PdownloadUrl := "wget -O paper.jar " + buildsUrl
// 				downloadRes := exec.Command(PdownloadUrl)
// 				downloadRes.Run()

// 				fmt.Println("Installed Paper " + version + " successfully!")
// 				// Install Paper
// 			}
// 		}
// 	}
// }

// func getBuildVer(jsonIn string) string {
// 	jsonData := []byte(jsonIn)

// 	var data map[string]interface{}
// 	err := json.Unmarshal(jsonData, &data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	builds := data["builds"].([]interface{})
// 	max := builds[0].(float64)
// 	for _, build := range builds {
// 		if build.(float64) > max {
// 			max = build.(float64)
// 		}
// 	}
// 	fmt.Println("Using build version: ", max)
// 	dataOut := strconv.FormatFloat(max, 'E', -1, 64)
// 	return dataOut
// }
