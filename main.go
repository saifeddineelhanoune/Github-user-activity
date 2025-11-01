package main

import ( 
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type JsonResponse struct {
	CurrentUrl string `json:"current_url"`
}

func main() {
	resp, err := http.Get("https://api.github.com/users/saifeddineelhanoune/events")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data JsonResponse
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	fmt.Println(data.CurrentUrl)
}