package main

import (
	"net/http"
	"log"

	"github.com/gargath/vfsgen-example/pkg/data"
)

func main() {
	fs := http.FileServer(data.Assets)
	http.Handle("/", fs)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
