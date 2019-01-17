package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//	Get file by path
func ReadFile(path string) []byte {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	return data
}

//	Write file
func WriteFile(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Request(method, url string, data interface{}) []byte {
	if method == "GET" {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			log.Printf("HTTP Request Get Fail URL: %s, Status: %d, Data: %s", url, resp.StatusCode, ReadRequest(resp.Body))
		}
		resp_data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err.Error())
		}
		return resp_data
	} else if method == "POST" {
		return nil
	}

	return nil
}

func ReadRequest(read io.Reader) string {
	req_buffer, err := ioutil.ReadAll(read)
	if err != nil {
		log.Printf("ReadRequest: %s", err.Error())
	}
	return string(req_buffer)
}
