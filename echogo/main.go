package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen")
	flag.IntVar(&port, "port", 9090, "port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		buffer := bytes.NewBuffer(nil)
		//print request method and URI
		buffer.WriteString(fmt.Sprintln(req.Method, req.RequestURI))

		buffer.WriteString(fmt.Sprintln())
		//print headers
		for k, v := range req.Header {
			buffer.WriteString(fmt.Sprintln(k+":", strings.Join(v, ";")))
		}
		buffer.WriteString(fmt.Sprintln())

		//print request body
		reqBody, readErr := ioutil.ReadAll(req.Body)
		if readErr != nil {
			buffer.WriteString(fmt.Sprintln(readErr))
		} else {
			buffer.WriteString(fmt.Sprintln(string(reqBody)))
		}

		outputData := buffer.Bytes()
		fmt.Println(string(outputData))
		w.Write(outputData)
	})

	//listen and serve
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	fmt.Println("listen err,", err)
}
