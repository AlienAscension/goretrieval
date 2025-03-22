package main

import (
	"net/http"
	"log"	
	"io"
	"os"
)


func downloadFile(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	
	defer resp.Body.Close() // Closes the response body

	// Create File
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close() // Close file (Important)

	// Copy from response into file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Download erfolgreich:", filename)

}

func main() {
	url := "https://www.gutenberg.org/ebooks/21000.txt.utf-8"
	filename := "21000.txt"
	downloadFile(url,filename)
}
