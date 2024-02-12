package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetUniversityData(searchTerm string) *[]UniversityResponse {
	apiURL := fmt.Sprintf("http://universities.hipolabs.com/search?name=%s", searchTerm)
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error getting internet response.....\nCowardly quitting.....\n")
		os.Exit(-1)
	}
	defer response.Body.Close()
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil
	}
	universities := make([]UniversityResponse, 2)
	if err = json.Unmarshal(bodyData, &universities); err != nil {
		fmt.Println("Error - could not translate json data to struct properly")
	}
	return &universities
}
