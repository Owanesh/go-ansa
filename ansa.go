package ansa

import (
	"fmt"
)

// GetChannelByTopic returns the channel corresponding to the specified topic, or error when it could not be found.
func GetChannelByTopic(topic string) (Channel, error) {
	rss, err := getDecodedTopic(topic)
	if err != nil {
		return Channel{}, err
	}

	return rss.Channel, nil
}

// GetFeedsByTopic returns the feed list corresponding to the specified topic, or error when it could not be found.
func GetFeedsByTopic(topic string) ([]Feed, error) {
	rss, err := GetChannelByTopic(topic)
	if err != nil {
		return nil, err
	}

	return rss.Feeds, err
}

// GetTopicList returns the topic list.
func GetTopicList() []string {
	return TOPIC_LIST
}

// GetFeedByGuid returns the feed given the specified guid and topic, or error when it could not be found.
func GetFeedByGuid(guid string, topic string) (Feed, error) {
	feeds, err := GetFeedsByTopic(topic)
	if err != nil {
		return Feed{}, err
	}

	for _, feed := range feeds {
		if feed.Guid == guid {
			return feed, nil
		}
	}

	return Feed{}, fmt.Errorf("GUID %s has no matching topic %s", guid, topic)
}

// GetFeedByTitle returns the feed given the specified title and topic, or error when it could not be found.
func GetFeedByTitle(title string, topic string) (Feed, error) {
	feeds, err := GetFeedsByTopic(topic)
	if err != nil {
		return Feed{}, err
	}

	for _, feed := range feeds {
		if feed.Title == title {
			return feed, nil
		}
	}

	return Feed{}, fmt.Errorf("Title %s has no matching topic %s", title, topic)
}
