package storage

import (
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/Owanesh/go-ansa/pkg/fetching"
)

const baseURL string = "http://www.ansa.it/"

const (
	homepage            string = "Homepage"
	cronaca             string = "Cronaca"
	politica            string = "Politica"
	world               string = "Mondo"
	economy             string = "Economia"
	soccer              string = "Calcio"
	sport               string = "Sport"
	cinema              string = "Cinema"
	culture             string = "Cultura"
	techonology         string = "Tecnologia"
	lasthour            string = "Ultimaora"
	englishnews         string = "Englishnews"
	foto                string = "Foto"
	video               string = "Video"
	abruzzo             string = "Abruzzo"
	basilicata          string = "Basilicata"
	calabria            string = "Calabria"
	campania            string = "Campania"
	emiliaRomagna       string = "Emilia-Romagna"
	friuliVeneziaGiulia string = "Friuli-Venezia-Giulia"
	lazio               string = "Lazio"
	liguria             string = "Liguria"
	lombardia           string = "Lombardia"
	marche              string = "Marche"
	molise              string = "Molise"
	piemonte            string = "Piemonte"
	puglia              string = "Puglia"
	sardegna            string = "Sardegna"
	sicilia             string = "Sicilia"
	toscana             string = "Toscana"
	trentinoAltoAdige   string = "Trentino-Alto-Adige"
	umbria              string = "Umbria"
	valleAosta          string = "Valle-Aosta"
	veneto              string = "Veneto"
)

// Storage is the struct that implements services repositories
type Storage struct{}

// GetChannelByTopic makes an http request to the ansa topic, returning the resulting RSS
func (s *Storage) GetChannelByTopic(topic string) (*fetching.RSS, error) {

	link := s.getTopicLink(topic)
	resp, err := http.Get(link)
	if err != nil {
		return nil, errors.New("getting topic link")
	}
	rss := &fetching.RSS{}
	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return nil, err
	}
	return rss, nil
}

func (s *Storage) getTopicLink(topic string) string {

	var urls = map[string]string{
		homepage:            baseURL + "sito/ansait_rss.xml",
		cronaca:             baseURL + "sito/notizie/cronaca/cronaca_rss.xml",
		politica:            baseURL + "sito/notizie/politica/politica_rss.xml",
		world:               baseURL + "sito/notizie/mondo/mondo_rss.xml",
		economy:             baseURL + "sito/notizie/economia/economia_rss.xml",
		soccer:              baseURL + "sito/notizie/sport/calcio/calcio_rss.xml",
		sport:               baseURL + "sito/notizie/sport/sport_rss.xml",
		cinema:              baseURL + "sito/notizie/cultura/cinema/cinema_rss.xml",
		culture:             baseURL + "sito/notizie/cultura/cultura_rss.xml",
		techonology:         baseURL + "sito/notizie/tecnologia/tecnologia_rss.xml",
		lasthour:            baseURL + "sito/notizie/topnews/topnews_rss.xml",
		englishnews:         baseURL + "english/english_rss.xml",
		foto:                baseURL + "sito/photogallery/foto_rss.xml",
		video:               baseURL + "sito/videogallery/video_rss.xml",
		abruzzo:             baseURL + "abruzzo/notizie/abruzzo_rss.xml",
		basilicata:          baseURL + "basilicata/notizie/basilicata_rss.xml",
		calabria:            baseURL + "calabria/notizie/calabria_rss.xml",
		campania:            baseURL + "campania/notizie/campania_rss.xml",
		emiliaRomagna:       baseURL + "emiliaromagna/notizie/emiliaromagna_rss.xml",
		friuliVeneziaGiulia: baseURL + "friuliveneziagiulia/notizie/friuliveneziagiulia_rss.xml",
		lazio:               baseURL + "lazio/notizie/lazio_rss.xml",
		liguria:             baseURL + "liguria/notizie/liguria_rss.xml",
		lombardia:           baseURL + "lombardia/notizie/lombardia_rss.xml",
		marche:              baseURL + "marche/notizie/marche_rss.xml",
		molise:              baseURL + "molise/notizie/molise_rss.xml",
		piemonte:            baseURL + "piemonte/notizie/piemonte_rss.xml",
		puglia:              baseURL + "puglia/notizie/puglia_rss.xml",
		sardegna:            baseURL + "sardegna/notizie/sardegna_rss.xml",
		sicilia:             baseURL + "sicilia/notizie/sicilia_rss.xml",
		toscana:             baseURL + "toscana/notizie/toscana_rss.xml",
		trentinoAltoAdige:   baseURL + "trentino/notizie/trentino_rss.xml",
		umbria:              baseURL + "umbria/notizie/umbria_rss.xml",
		valleAosta:          baseURL + "valledaosta/notizie/valledaosta_rss.xml",
		veneto:              baseURL + "veneto/notizie/veneto_rss.xml",
	}

	return urls[topic]
}
