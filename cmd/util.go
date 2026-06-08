package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Get file by path
func ReadFile(path string) []byte {
	file, err := os.Open(path) // #nosec G304 -- local dev tool reads template files by path
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() { _ = file.Close() }()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return data
}

// Write file
func WriteFile(path, content string) {
	err := os.WriteFile(path, []byte(content), 0o600)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Coustom request for api
func Request(method, url string, data any) []byte {
	switch method {
	case "GET":
		client := &http.Client{}
		request, err := http.NewRequest(method, url, nil)
		if err != nil {
			log.Fatalln(err.Error())
		}
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Chrome/7.0.517.41 Safari/534.7")

		response, err := client.Do(request)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer func() { _ = response.Body.Close() }()
		if response.StatusCode != 200 {
			log.Printf("HTTP Request Get Fail URL: %s, Status: %d, Data: %s", url, response.StatusCode, ReadRequest(response.Body))
		}
		resp_data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err.Error())
		}
		return resp_data
	case "POST":
		return nil
	}

	return nil
}

func ReadRequest(read io.Reader) string {
	req_buffer, err := io.ReadAll(read)
	if err != nil {
		log.Printf("ReadRequest: %s", err.Error())
	}
	return string(req_buffer)
}
