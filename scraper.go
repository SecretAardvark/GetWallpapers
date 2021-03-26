package main

//TODO: Give each file a unique name so you can run the script multiple times without duplicates.
//TODO: Do the http requests concurrently.
import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	for i, link := range links[0:5] {
		//Send a get request to each URL.
		url := link.URL
		fileExt := url[len(url)-3 : len(url)-0]
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("HTTP error: ", err)
		}
		defer resp.Body.Close()

		//Copy the http response into a new file.
		fileName := "wallpaper" + fmt.Sprint(i) + "." + fileExt
		out, outErr := os.Create(fileName)
		if outErr != nil {
			log.Fatal("Error creatging file: ", outErr)
		}
		defer out.Close()
		b, copyerr := io.Copy(out, resp.Body)
		if copyerr != nil {
			log.Println("Error copying file", copyerr)
		}
		fmt.Println("File size: ", b)

	}
}
