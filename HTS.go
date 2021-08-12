package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Struct for HTS Server
type HTS struct {
	Port    string
	IP      string
	HomeDir string
	//logger  Logger
	ConfigData Config
}

//structure for getting data from configuration file
type Config struct {
	Restricted []string `json:"Restrictedroutes"` //use exact spelling used in json file
}

//Method for Identifying content type of the file with extension
func (hts HTS) GetContentType(extension string) string {
	switch extension {
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	case "js":
		return "text/javascript"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	case "jpg", "jpeg", "jfif", "pjpeg", "pjp":
		return "image/jpeg"
	case "webp":
		return "image/webp"
	case "svg":
		return "image/svg+xml"

	}
	return "text/plain"
}

//Method for Getting extension from file URL
func (hts HTS) GetExtension(path string) string {
	data := strings.Split(path, ".")
	return data[1]

}

//Method for check if the file exixts
func (hts HTS) IsFileExists(path string) (bool, string) {
	filelocation := hts.HomeDir + path
	_, err := os.Stat(filelocation)

	//if the index.html not found then use our Deafult index page--This is used to display the directoty does not have index.html file
	if err != nil && path == "/index.html" {
		return true, "Static/Default.html"
	}

	//return false when file not found
	if err != nil {
		return false, ""
	}

	//return true if the file existss
	return true, filelocation

}

//Parsing ACL.json for access control of routes
func (hts HTS) ParseConfig() {
	fmt.Println("Reading Config File")
	ConfigFile, err := os.Open("config.json")
	if err != nil {
		fmt.Print("\n Cannot Open config.json!")
	}
	defer ConfigFile.Close()
	//var data Config
	//var data map[string]interface{}
	jsondata, _ := ioutil.ReadAll(ConfigFile)
	json.Unmarshal(jsondata, &hts.ConfigData)

}

//Method for finding url in Restricted List
func (hts HTS) IsIn(query string, list []string) bool {

	for _, i := range list {
		if i == query {

			return true
		}
	}
	return false
}

//method for finding a authorized url
func (hts HTS) IsAuthorizedRoute(route string) bool {
	return hts.IsIn(route, hts.ConfigData.Restricted)
}

//Main handler Function for / route
func (hts HTS) HandleHome(response http.ResponseWriter, request *http.Request) {

	//Filtering Methods
	if request.Method != "GET" {
		http.Error(response, "Method Not Alowed", http.StatusMethodNotAllowed)
	}

	//Form Empty request Url have the value of /
	//Get path from the URL
	url := request.URL.Path

	//check for the request to root of directory if it is then apppend index.html as url string
	if url == "/" {
		url = url + "index.html"
	}

	fmt.Printf("\nMethod:%s From:%v Path:%s", request.Method, request.RemoteAddr, url)

	//check if the url authorized
	if hts.IsAuthorizedRoute(url) {
		http.Error(response, "Unauthorised", http.StatusUnauthorized)

	}
	//Get the status of file exists
	result, Location := hts.IsFileExists(url)
	//result := true
	//Location := hts.HomeDir + "index.html"
	//If file Not exists
	if !result {
		file, _ := os.Open("Static/NotFound.html")
		defer file.Close()
		file.Seek(0, 0)
		io.Copy(response, file)

		defer file.Close()
		http.Error(response, "", 404)
	} else { //If file Exists
		extension := hts.GetExtension(Location)
		//extension := "html"
		contenttype := hts.GetContentType(extension)
		response.Header().Set("Content-Type", contenttype)
		response.WriteHeader(http.StatusOK)
		file, _ := os.Open(Location)
		defer file.Close()

		//setting the cursor to start of the file
		file.Seek(0, 0)

		//Copy contents of the file to Response
		io.Copy(response, file)

	}

}
