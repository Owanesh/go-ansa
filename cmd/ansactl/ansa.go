package main

import (
	"fmt"

	"../.." // for local testing
	//"github.com/Owanesh/go-ansa"
)

func main() {
	url := "http://www.ansa.it/sito/notizie/sport/calcio/2018/10/27/consiglio-stato-la-serie-b-torna-a-19_47b88159-3664-4281-8922-1904014302ac.html"
	feed, err := ansa.GetFeedByGuid(url, ansa.HOMEPAGE)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("feed title: %v\n", feed.Title)
	fmt.Printf("feed GUID: %v\n", feed.Guid)
}
