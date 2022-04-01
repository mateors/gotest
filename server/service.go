package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/errors/fmt"
)

var data = make(map[string]interface{})

func init() {

	data["data"] = map[string]interface{}{
		"projects": map[string]interface{}{
			"nodes": []map[string]interface{}{
				{"name": "hcs_utils", "description": "", "forksCount": 1},
				{"name": "K", "description": nil, "forksCount": 1},
				{"name": "Heroes of Wesnoth", "description": nil, "forksCount": 5},
				{"name": "Leiningen", "description": "", "forksCount": 1},
				{"name": "TearDownWalls", "description": nil, "forksCount": 5},
			},
		},
	}

}

//2a. service layer
func main() {

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	LOG_FILE := filepath.Join(workingDirectory, "log.txt")
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	//log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

//2b. service layer return
func nodeCounter(nodes []map[string]interface{}, last int) map[string]interface{} {

	var outputMap = make(map[string]interface{})
	var names string
	var total, c int

	for _, row := range nodes {

		names += fmt.Sprintf("%s,", row["name"])
		forksCount, err := strconv.Atoi(fmt.Sprintf("%v", row["forksCount"]))
		if err != nil {
			log.Println(err)
		}
		total += forksCount
		c++
		if c == last {
			break
		}
	}
	outputMap["names"] = strings.TrimRight(names, ",")
	outputMap["total_forks"] = total
	return outputMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, `API endpoint http://localhost:8080/api`)
}

//2b. service layer return
func apiHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if rec := recover(); rec != nil {
			log.Println("apiHandler panicking >>", rec)
		}
	}()

	var req = make(map[string]interface{})
	var errorFound bool

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorFound = true
		req["error"] = 1
		req["message"] = `Bad request, Please send a json body request, for example { "last": 2 }`
	}

	if !errorFound {

		//fmt.Println("request:", req)
		if last, isOk := req["last"]; isOk {

			lastn, err := strconv.Atoi(fmt.Sprintf("%v", last))
			if err != nil {
				req["error"] = 2
				req["message"] = "invalid number"
				log.Printf("ERROR %v is not a number", last)
				writeResponse(w, &req)
				return
			}
			if lastn < 0 {
				req["error"] = 3
				req["message"] = "wrong number"
				log.Printf("ERROR %v is a negative number", last)
				writeResponse(w, &req)
				return
			}
			projects := data["data"].(map[string]interface{})["projects"].(map[string]interface{})
			nodes := projects["nodes"].([]map[string]interface{})
			req = nodeCounter(nodes, lastn)
		}
	}
	writeResponse(w, &req)
}

func writeResponse(w http.ResponseWriter, req interface{}) {

	bs, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `%s`, bs)
}
