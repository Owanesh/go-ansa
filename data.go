package ansa

// Feed contains information about the RSS feed
type Feed struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

// Channel contains information about the channel. It contains zero or more feeds.
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Copyright   string `xml:"copyright"`
	Feeds       []Feed `xml:"item"`
}

// RSS contains the Channel.
type RSS struct {
	Channel Channel `xml:"channel"`
}
