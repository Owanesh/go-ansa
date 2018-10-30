package fetching

// Feed holds information about an ansa feed
type Feed struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

// RSS is the struct to hold the ansa rss response
type RSS struct {
	Channel Channel `xml:"channel"`
}

// Channel holds a lot of feeds
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Copyright   string `xml:"copyright"`
	Feeds       []Feed `xml:"item"`
}
