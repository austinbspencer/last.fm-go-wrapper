package main

import (
	"fmt"
	"log"

	lastfm "github.com/austinbspencer/last.fm-go-wrapper"
)

func testLastfmFunc(client *lastfm.Client) {
	res, err := client.GeoGetTopTracks("usa")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func main() {
	// hClient := http.Client{Timeout: time.Duration(1) * time.Second}
	// client := lastfm.New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))

	// client.LibraryGetArtists("abspen1", "100")
}
