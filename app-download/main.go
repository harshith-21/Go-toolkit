package main

import (
	"fmt"
	"github.com/harshith-21/toolkit"
	"log"
	"net/http"
)

func main() {
	//get routes
	mux := routes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)

	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	t := toolkit.Tools{}
	t.DownloadStaticFile(w, r, "./files", "gola.png", "golang.png")
	fmt.Println("downloading file")
}
