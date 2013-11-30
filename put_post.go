package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "net/url"
)

func postExample() {
        values := make(url.Values)
        values.Set("email", "anything@stathat.com")
        values.Set("stat", "messages sent - female to male")
        values.Set("count", "1")
        
        r, err := http.PostForm("http://localhost:8082/test", url.Values{"{\"parameters\": [\"stan\": 1]}": {"\"that\""}})
        if err != nil {
            log.Printf("error posting stat to stathat: %s", err)
            return
        }
        body, _ := ioutil.ReadAll(r.Body)
        r.Body.Close()
        log.Printf("stathat post result body: %s", body)
}

func main() {
    postExample()
    log.Printf("done")
}