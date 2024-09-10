package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	const PANIC = "panic"

	defer func() {
		if r := recover(); r != nil {
			if r == PANIC {
				fmt.Println(PANIC + " hai na")
			}
		}
	}()

	args := os.Args
	var port int = 5000
	var err error
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(args[1])
		if err != nil {
			panic(PANIC)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "0") // "Hello, %q\n", html.EscapeString(r.URL.Path))
		if n >= 0 {
			log.Printf("wrote %d bytes", n)
		}
		if err != nil {
			fmt.Println("Error: ", err)
		}

	})

	log.Println("Listening on localhost:" + strconv.Itoa(port))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

	fmt.Println(port)
}
