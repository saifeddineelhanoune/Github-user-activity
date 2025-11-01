# GitHub User Activity Tracker

This is a simple Go application that fetches and displays the recent public activity for a specific GitHub user.

## Description

The program sends a GET request to the GitHub API's `/users/{username}/events` endpoint to retrieve a list of recent public events performed by the user `saifeddineelhanoune`. It then parses the JSON response and prints out the type of each event and the repository where the event occurred.

This serves as a basic example of how to interact with a JSON API in Go, handle HTTP requests, and unmarshal JSON data into Go structs.

## How It Works

1.  **HTTP GET Request**: It uses the `net/http` package to make a GET request to `https://api.github.com/users/saifeddineelhanoune/events`.
2.  **Read Response Body**: The `io` package is used to read the entire response body.
3.  **JSON Unmarshaling**: The `encoding/json` package unmarshals the incoming JSON array into a slice of `Event` structs. The `Event` struct is defined to match the structure of the data we care about from the API response.
4.  **Print Results**: The program iterates through the slice of events and prints the `Type` of the event and the `Name` of the repository for each event to the console.

## How to Run

To run this program, you need to have Go installed on your machine.

1.  Clone the repository or save the `main.go` file.
2.  Open your terminal and navigate to the directory containing `main.go`.
3.  Run the following command:

    ```bash
    go mod init <mod-name>
    go mod tidy
    go run main.go
    ```

### Project Url:
https://github.com/saifeddineelhanoune/Github-user-activity

