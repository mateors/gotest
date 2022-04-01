package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/exp/errors/fmt"
)

//4. print results via call that service from main.go
func main() {

	var jsonStr = []byte(`{"last":"5"}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/api", bytes.NewBuffer(jsonStr))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	//var resMap = make(map[string]interface{})
	//json.NewDecoder(resp.Body).Decode(&resMap)
	//fmt.Println(resMap)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
