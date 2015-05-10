package main

import (
    "fmt"
	"encoding/json"
    "net/http"
	"io/ioutil"
	"log"
)

var file []byte
var jsontype jsonobject

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><span>\"Brune Blamed\"<span><em>%v</em></html>", jsontype.Object.Counter)
    jsontype.Object.Counter++
}

func bruneHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><em>%v</em><form action=\"/Blame Brune\" method=\"POST\">"+
        "<input type=\"submit\" value=\"Blame Brune\">"+
        "</form></html>", jsontype.Object.Counter)
	data, err := json.Marshal(jsontype)
	err = ioutil.WriteFile("config.json", data, 755)
	if err != nil {
		fmt.Println("Can't write file.")
		return
	}
	log.Printf("Brunes counted: %v", jsontype)
}

type jsonobject struct {
    Object ObjectType
}

type ObjectType struct {
	Counter int64
}

func main() {
	file, e := ioutil.ReadFile("config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
    }
    json.Unmarshal(file, &jsontype)
    http.HandleFunc("/", handler)
	http.HandleFunc("/Blame Brune", bruneHandler)
    http.ListenAndServe(":8080", nil)
	fmt.Print(jsontype)
}
