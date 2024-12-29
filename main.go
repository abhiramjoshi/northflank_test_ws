package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func incrementCounterReset(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter = 0
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func printSecret(w http.ResponseWriter, r *http.Request) {
	env, exists := os.LookupEnv("PRINT_SECRET")
	if exists {
		fmt.Fprintf(w, string(env))
	} else {
		fmt.Fprintf(w, "No secret set")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/increment/reset", incrementCounterReset)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/secret", printSecret)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
