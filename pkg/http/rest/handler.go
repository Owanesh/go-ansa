package rest

import (
	"encoding/xml"
	"net/http"

	"github.com/Owanesh/go-ansa/pkg/fetching"

	"github.com/julienschmidt/httprouter"
)

// Handler is where the router paths are setup
func Handler(service fetching.Service) http.Handler {
	router := httprouter.New()

	router.GET("/channels/:topic", getFeeds(service))
	return router
}

func getFeeds(service fetching.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		topic := p.ByName("topic")
		channel, err := service.GetChannelByTopic(topic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(channel)
	}
}
