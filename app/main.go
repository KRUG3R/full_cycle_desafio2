package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	cep := "01153000"
	go GetResponse("https://brasilapi.com.br/api/cep/v1/"+cep, &waitGroup)
	go GetResponse("https://viacep.com.br/ws/"+cep+"/json/", &waitGroup)

	waitGroup.Wait()
}

func GetResponse(endpoint string, wg *sync.WaitGroup) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Retorno da API:" + endpoint)
	fmt.Println(string(body))

	wg.Done()
}
