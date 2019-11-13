package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	WhatsMyIP()
}

func WhatsMyIP() {
	resp, err := http.Get("https://ifconfig.co")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(strings.TrimSpace(string(body)))
}
