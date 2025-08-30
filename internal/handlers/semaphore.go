package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/api/projects", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer 2unxfxg7ikkv3pacj5s6vomri1lcavyqty9xr1x5ny8=")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Resposta:", string(body))

}
