package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func send(request HttpRequest, apiKey string) string {
	url := fmt.Sprintf("%s%s", request.baseUrl, request.path)
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request.data)
	req, _ := http.NewRequest(request.method, url, buffer)
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return ""
	}

	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	return string(responseBody)
}
