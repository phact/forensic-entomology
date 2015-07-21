package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	ServeHTTP()
}

func makeRequest(api string, args string) http.HandlerFunc{

  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    var q string

    q = r.FormValue("q")


    if q != "" {
      //args = q
    }
    fmt.Printf(api + "?"+ args)

    response, err := http.Get(api + "?" + args)
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

  return handler

}

func ServeHTTP() {
	var addrFlag string

  addrFlag = "0.0.0.0:9999"

  hProjects := makeRequest("https://issues.apache.org/jira/rest/api/2/project","")


  var args string
  var args22 string
  var args3 string
	
  args = "jql=project+%3D+CASSANDRA+AND+fixVersion+in+%282.1.4%2C+%222.1+beta1%22%2C+%222.1+beta2%22%2C+%222.1+rc1%22%2C+%222.1+rc2%22%2C+%222.1+rc3%22%2C+%222.1+rc4%22%2C+%222.1+rc5%22%2C+%222.1+rc6%22%2C+2.1.0%2C+2.1.1%2C+2.1.2%2C+2.1.3%29+ORDER+BY+watchers+DESC&maxResults=100"


  hIssues := makeRequest("https://issues.apache.org/jira/rest/api/2/search",args)


  hVersions := makeRequest("https://issues.apache.org/jira/rest/api/2/project/CASSANDRA/",args)


	args22 = "jql=project+%3D+CASSANDRA+AND+fixVersion+in+(2.2.0,'2.2.0%20beta%201',2.2.x,'2.2.0%20rc1','2.2.0%20rc2',2.2.0,2.2.1)+ORDER+BY+watchers+DESC&maxResults=100"
	args3 = "jql=project+%3D+CASSANDRA+AND+fixVersion+in+(3.0.x,'3.0%20beta%201','3.0.0%20rc1')+ORDER+BY+watchers+DESC&maxResults=100"
  hIssues22 := makeRequest("https://issues.apache.org/jira/rest/api/2/search",args22)
  hIssues3 := makeRequest("https://issues.apache.org/jira/rest/api/2/search",args3)


	http.Handle("/", http.FileServer(http.Dir("./www/")))
	http.Handle("/jira.json", http.StripPrefix("/jira.json", hIssues))
	http.Handle("/jira22.json", http.StripPrefix("/jira22.json", hIssues22))
	http.Handle("/jira3.json", http.StripPrefix("/jira3.json", hIssues3))
	http.Handle("/projects.json", hProjects)
	http.Handle("/versions.json", hVersions)

	err := http.ListenAndServe(addrFlag, nil)
	if err != nil {
		log.Fatalf("net.http could not listen on address '%s': %s\n", addrFlag, err)
	}
}

