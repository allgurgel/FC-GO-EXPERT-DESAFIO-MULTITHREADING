package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ApiCep1(cep string, canal chan string) {
	req, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		return
	}

	if req.StatusCode != 200 {
		return
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	canal <- string(res)
}

func ApiCep2(cep string, canal chan string) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return
	}

	if req.StatusCode != 200 {
		return
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	canal <- string(res)
}

// Thread 1
func main() {
	var cep string = "01001-000"
	canalApi1 := make(chan string)
	canalApi2 := make(chan string)

	go ApiCep1(cep, canalApi1)
	go ApiCep2(cep, canalApi2)

	select {
	case api1 := <-canalApi1:
		fmt.Println("Dados da API: cdn.apicep.com", api1)

	case api2 := <-canalApi2:
		fmt.Println("Dados da API: viacep.com.br", api2)

	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
