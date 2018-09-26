package ansa
import (
	"errors"
	"encoding/xml"
	"net/http"
)

func isValidTopic(topic string) (bool,error){
	check := CHANNELS_URL[topic];
	if check != "" {
		return true, nil
	} else {
		return false, errors.New("Invalid topic \"" + topic+ "\"")
	}
}

func getChannelLinkByTopic(topic string) (string,error){
	valid, err := isValidTopic(topic)
	if err != nil && !valid {
			return "", err
	}else{
		return CHANNELS_URL[topic], nil
	}
}

func getXmlDecodedByLink(link string) (RSS, error) {
	var RSSerr RSS
	resp, err := http.Get(link)
	if err != nil {
		return RSSerr, nil
	}
	defer resp.Body.Close()
	rss := RSS{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return rss, err
	}
	return rss, nil
}

func getDecodedTopic(topic string) (RSS,error) {
	var RSSerr RSS
	link, err := getChannelLinkByTopic(topic)
	if link!="" {
		return getXmlDecodedByLink(link)
	}
	return RSSerr, err
}
