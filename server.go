package main

import (
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"strings"
)

func main() {
	ServeHTTP()
}
func ServeHTTP() {
	var addrFlag string

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var args string
		args = "jql=project+%3D+CASSANDRA+AND+fixVersion+in+%282.1.4%2C+%222.1+beta1%22%2C+%222.1+beta2%22%2C+%222.1+rc1%22%2C+%222.1+rc2%22%2C+%222.1+rc3%22%2C+%222.1+rc4%22%2C+%222.1+rc5%22%2C+%222.1+rc6%22%2C+2.1.0%2C+2.1.1%2C+2.1.2%2C+2.1.3%29+ORDER+BY+watchers+DESC&maxResults=100"

		response, err := http.Get("https://issues.apache.org/jira/rest/api/2/search?" + args)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			//fmt.Printf("%s\n", string(contents))
			w.Header().Set("Content-Type", "application/json")
			w.Write(contents)
			//w.Write([]byte("{\"hello\": \"world\"}"))
		}

	})

	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.Handle("/jira.json", http.StripPrefix("/jira.json", h))

	err := http.ListenAndServe(addrFlag, nil)
	if err != nil {
		log.Fatalf("net.http could not listen on address '%s': %s\n", addrFlag, err)
	}
}

/*
func InventoryDataHandler(w http.ResponseWriter, r *http.Request) {
	files := InventoryData("www")

	// separate logfiles by log type
	out := make(map[string][]string)
	for _, file := range files {
		base := strings.TrimSuffix(file, ".json")
		logtype := strings.Split(base, "-")[1]

		if _, ok := out[logtype]; !ok {
			out[logtype] = make([]string, 0)
		}
		out[logtype] = append(out[logtype], file)
	}

	json, err := json.Marshal(out)
	if err != nil {
		log.Printf("JSON marshal failed: %s\n", err)
		http.Error(w, fmt.Sprintf("Marshaling JSON failed: %s", err), 500)
	}

	w.Write(json)
}
*/
