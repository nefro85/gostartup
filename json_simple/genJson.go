package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JobStatus struct {
	JobName       string `json:"jobName"`
	JobStatusName string `json:"jobStatusName"`
}

type CiReport struct {
	Date string      `json:"date"`
	Jobs []JobStatus `json:"jobs"`
}

func main() {
	console := fmt.Println
	data := readJson()

	var storage map[string]interface{}
	var report CiReport

	if err := json.Unmarshal(data, &storage); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &report); err != nil {
		panic(err)
	}

	console("Map values:")
	console(storage)

	console("CiReport struct:")
	console(report)

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readJson() []byte {

	path := os.Getenv("GOPATH")
	path = path + "/src/gostartup/json_simple/"

	data, err := ioutil.ReadFile(path + "jobsStatus.json")
	checkErr(err)
	return data
}
