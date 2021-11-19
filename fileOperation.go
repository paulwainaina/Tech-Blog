package main

import (
	"encoding/json"
	"io/ioutil"
)

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}
type Blog struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	Url         string `json:"Url"`
}

func ReadDataFromFile(title string) ([]byte, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetJsonData(title string, _interface interface{}) {
	jsondata, err := ReadDataFromFile(title)
	if err != nil {
		_interface = nil
	}
	json.Unmarshal(jsondata, _interface)
}
