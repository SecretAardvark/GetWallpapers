package main

//TODO: Do the http requests concurrently.
import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/painhardcore/go-reddit"
)

func main() {
	//parse the page to find the image urls
	ctx := context.Background()
	rclient := reddit.NoAuthClient
	links, err := rclient.GetHotLinks(ctx, "wallpapers")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Retrieved top 5 links.")

	for _, link := range links[0:5] {
		//Send a get request to each URL.
		url := link.URL
		split := strings.Split(url, "/")
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("HTTP error: ", err)
		}
		defer resp.Body.Close()

		//Copy the http response into a new file.
		fileName := split[len(split)-1]
		fmt.Println(split[len(split)-1])
		out, outErr := os.Create(fileName)
		if outErr != nil {
			log.Fatal("Error creating file: ", outErr)
		}
		defer out.Close()
		b, copyerr := io.Copy(out, resp.Body)
		if copyerr != nil {
			log.Println("Error copying file", copyerr)
		}
		fmt.Println("File size: ", b)

	}
}
