package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lastfm "github.com/austinbspencer/last.fm-go-wrapper"
)

func main() {
	hClient := http.Client{Timeout: time.Duration(1) * time.Second}
	client := lastfm.New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))

	err := client.SetUser("abspen1")

	if err != nil {
		log.Fatal(err)
	}

	// limit := lastfm.LimitOpt(2)
	// binary := lastfm.Binary(lastfm.One)
	// extend := lastfm.ExtendedOpt(binary)

	// client.UserGetRecentTracks("Abspen1", limit)
	res, err := client.UserGetRecentTracks(client.User, lastfm.ExtendedOpt(lastfm.Zero), lastfm.LimitOpt(2))

	// client.UserGetRecentTracks(client.User, lastfm.ExtendedOpt(lastfm.One), lastfm.LimitOpt(2))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
