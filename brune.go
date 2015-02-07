package main

import (
    "fmt"
	"encoding/json"
    "net/http"
	"io/ioutil"
)

var counter int64
var file []byte
var jsontype jsonobject

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><form action=\"/Blame Brune\" method=\"POST\">"+
        "<input type=\"submit\" value=\"Blame Brune\">"+
        "</form></html>")
}

func bruneHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><form action=\"/Blame Brune\" method=\"POST\">"+
        "<input type=\"submit\" value=\"Blame Brune\">"+
        "</form></html>")
	fmt.Print(jsontype)
	jsontype.Object.counter++
	data, err := json.Marshal(jsontype)
	err = ioutil.WriteFile("config.json", data, 755)
	if err != nil {
		fmt.Println("Can't write file.")
		return
	}
}

type jsonobject struct {
    Object ObjectType
}

type ObjectType struct {
	counter int64
}

func main() {
	file, e := ioutil.ReadFile("config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
		return
    }
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Brunes counted: %v\n", jsontype)
    fmt.Printf("%s\n", string(file))
    http.HandleFunc("/", handler)
	http.HandleFunc("/Blame Brune", bruneHandler)
    http.ListenAndServe(":8080", nil)
}
