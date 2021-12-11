package createfile

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullUrlFile string
)

func Easy(URL string) {
	fullUrlFile = URL
	// Build fileName from fullPath
	BuildFileName()

	// Create blank file
	file := CreateFile()

	// Put content on file
	PutFile(file, HttpClient())
}

func PutFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)

	CheckError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	CheckError(err)

	fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func BuildFileName() {
	fileUrl, err := url.Parse(fullUrlFile)
	CheckError(err)

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func HttpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func CreateFile() *os.File {
	file, err := os.Create(fileName)

	CheckError(err)
	return file
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
