package ansa

func GetChannelByTopic(topic string)(Channel, error){
	var ChannelErr Channel
	rss, err := getDecodedTopic(topic)
	if err != nil {
		return ChannelErr, err
	}
	return rss.Channel, nil
}

func GetFeedsByTopic(topic string) ([]Feed, error){
	var FeedErr []Feed
	var err error
	rss, err := GetChannelByTopic(topic)
	if err != nil {
		return FeedErr, nil
	}
	return rss.Feeds, err
}

func GetTopicList() ([]string){
	return TOPIC_LIST
}

func GetFeedByGuid(guid string, topic string)(Feed,error){
	var FeedErr Feed
	feeds, err := GetFeedsByTopic(topic)

	for _, feed := range feeds {
		if feed.Guid == guid {
			return feed, nil
		}
	}
	if err != nil {
		return FeedErr, err
	}
	return FeedErr, nil
}

func GetFeedByTitle(title string, topic string)(Feed, error){
	var FeedErr Feed
	feeds, err := GetFeedsByTopic(topic)

	for _, feed := range feeds {
		if feed.Title == title {
			return feed, nil
		}
	}
	if err != nil {
		return FeedErr, err
	}
	return FeedErr, nil
}
