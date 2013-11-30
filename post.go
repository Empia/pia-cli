package main

import (
    "encoding/json"
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
)

type test_struct struct {
    Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    var t test_struct
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
    responseXML, _ := json.Marshal(t)
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, string(responseXML))
 //   fmt.Fprint(rw, string({"success": true, "message": "Hello!"}))
 
}

func main() {
    http.HandleFunc("/test", test)
    log.Fatal(http.ListenAndServe(":8082", nil))
}