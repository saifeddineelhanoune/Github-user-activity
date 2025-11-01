package main

import ( 
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
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

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		panic(err)
	}

	for _, event := range events {
		fmt.Printf("Event Type: %s, Repo: %s\n", event.Type, event.Repo.Name)
	}
}