package ansa

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func isValidTopic(topic string) (bool, error) {
	if topic, ok := CHANNELS_URL[topic]; !ok {
		return false, fmt.Errorf("Invalid topic: %s", topic)
	}

	return true, nil
}

func getChannelLinkByTopic(topic string) (string, error) {
	if _, err := isValidTopic(topic); err != nil {
		return "", err
	}

	return CHANNELS_URL[topic], nil
}

func getXmlDecodedByLink(link string) (RSS, error) {
	resp, err := http.Get(link)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()

	var rss RSS
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return rss, err
	}

	return rss, nil
}

func getDecodedTopic(topic string) (RSS, error) {
	link, err := getChannelLinkByTopic(topic)
	if err != nil {
		return RSS{}, err
	}

	return getXmlDecodedByLink(link)
}
