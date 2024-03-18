package main

import (
	"io"
	"log"
	"net/http"
)

func echoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("New request...")
	if request == nil {
		http.Error(writer, "empty request", http.StatusInternalServerError)
		return
	}

	resp, err := process(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = writer.Write([]byte("Echo back: '" + resp + "'\n"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func process(request *http.Request) (string, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return "", err
	}
	request.Body.Close()
	return string(body), nil
}

func main() {
	const port = "8081"
	log.Println("Starting server, listening on port", port)
	http.HandleFunc("/", echoHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
