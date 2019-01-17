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

//	Coustom request for api
func Request(method, url string, data interface{}) []byte {
	if method == "GET" {
		client := &http.Client{}
		request, err := http.NewRequest(method, url, nil)
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Chrome/7.0.517.41 Safari/534.7")

		response, err := client.Do(request)
		if err != nil || response.StatusCode != 200 {
			log.Printf("HTTP Request Get Fail URL: %s, Status: %d, Data: %s", url, response.StatusCode, ReadRequest(response.Body))
		}
		resp_data, err := ioutil.ReadAll(response.Body)
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
