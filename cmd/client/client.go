package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/IsJordanBraz/go-client-server-api/internal/dto"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Error Requesting: " + err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error Requesting: " + err.Error())
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading body: " + err.Error())
		return
	}

	var bid dto.CotacaoOutput
	err = json.Unmarshal(res, &bid)

	if err != nil {
		fmt.Println("Error Unmarshal: " + err.Error())
		return
	}

	file, err := os.Create("cotacao.txt")

	if err != nil {
		fmt.Println("Error Creating File: " + err.Error())
		return
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s \n", bid.Bid))
	if err != nil {
		fmt.Println("Error Writing File: " + err.Error())
		return
	}
}
