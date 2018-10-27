package ansa

type Feed struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Copyright   string `xml:"copyright"`
	Feeds       []Feed `xml:"item"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}
