package ansa

const BASE_URL string = "http://www.ansa.it/"

const HOMEPAGE string = "Homepage"
const CRONACA string = "Cronaca"
const POLITICA string = "Politica"
const WORLD string = "Mondo"
const ECONOMY string = "Economia"
const SOCCER string = "Calcio"
const SPORT string = "Sport"
const CINEMA string = "Cinema"
const CULTURE string = "Cultura"
const TECHONOLOGY string = "Tecnologia"
const LASTHOUR string = "Ultimaora"
const ENGLISHNEWS string = "Englishnews"
const FOTO string = "Foto"
const VIDEO string = "Video"
const ABRUZZO string = "Abruzzo"
const BASILICATA string = "Basilicata"
const CALABRIA string = "Calabria"
const CAMPANIA string = "Campania"
const EMILIA_ROMAGNA string = "Emilia-Romagna"
const FRIULI_VENEZIA_GIULIA string = "Friuli-Venezia-Giulia"
const LAZIO string = "Lazio"
const LIGURIA string = "Liguria"
const LOMBARDIA string = "Lombardia"
const MARCHE string = "Marche"
const MOLISE string = "Molise"
const PIEMONTE string = "Piemonte"
const PUGLIA string = "Puglia"
const SARDEGNA string = "Sardegna"
const SICILIA string = "Sicilia"
const TOSCANA string = "Toscana"
const TRENTINO_ALTO_ADIGE string = "Trentino-Alto-Adige"
const UMBRIA string = "Umbria"
const VALLE_AOSTA string = "Valle-Aosta"
const VENETO string = "Veneto"

var TOPIC_LIST = []string{
	HOMEPAGE,
	CRONACA,
	POLITICA,
	WORLD,
	ECONOMY,
	SOCCER,
	SPORT,
	CINEMA,
	CULTURE,
	TECHONOLOGY,
	LASTHOUR,
	ENGLISHNEWS,
	FOTO,
	VIDEO,
	ABRUZZO,
	BASILICATA,
	CALABRIA,
	CAMPANIA,
	EMILIA_ROMAGNA,
	FRIULI_VENEZIA_GIULIA,
	LAZIO,
	LIGURIA,
	LOMBARDIA,
	MARCHE,
	MOLISE,
	PIEMONTE,
	PUGLIA,
	SARDEGNA,
	SICILIA,
	TOSCANA,
	TRENTINO_ALTO_ADIGE,
	UMBRIA,
	VALLE_AOSTA,
	VENETO}

var CHANNELS_URL = map[string]string{
	HOMEPAGE:              BASE_URL + "sito/ansait_rss.xml",
	CRONACA:               BASE_URL + "sito/notizie/cronaca/cronaca_rss.xml",
	POLITICA:              BASE_URL + "sito/notizie/politica/politica_rss.xml",
	WORLD:                 BASE_URL + "sito/notizie/mondo/mondo_rss.xml",
	ECONOMY:               BASE_URL + "sito/notizie/economia/economia_rss.xml",
	SOCCER:                BASE_URL + "sito/notizie/sport/calcio/calcio_rss.xml",
	SPORT:                 BASE_URL + "sito/notizie/sport/sport_rss.xml",
	CINEMA:                BASE_URL + "sito/notizie/cultura/cinema/cinema_rss.xml",
	CULTURE:               BASE_URL + "sito/notizie/cultura/cultura_rss.xml",
	TECHONOLOGY:           BASE_URL + "sito/notizie/tecnologia/tecnologia_rss.xml",
	LASTHOUR:              BASE_URL + "sito/notizie/topnews/topnews_rss.xml",
	ENGLISHNEWS:           BASE_URL + "english/english_rss.xml",
	FOTO:                  BASE_URL + "sito/photogallery/foto_rss.xml",
	VIDEO:                 BASE_URL + "sito/videogallery/video_rss.xml",
	ABRUZZO:               BASE_URL + "abruzzo/notizie/abruzzo_rss.xml",
	BASILICATA:            BASE_URL + "basilicata/notizie/basilicata_rss.xml",
	CALABRIA:              BASE_URL + "calabria/notizie/calabria_rss.xml",
	CAMPANIA:              BASE_URL + "campania/notizie/campania_rss.xml",
	EMILIA_ROMAGNA:        BASE_URL + "emiliaromagna/notizie/emiliaromagna_rss.xml",
	FRIULI_VENEZIA_GIULIA: BASE_URL + "friuliveneziagiulia/notizie/friuliveneziagiulia_rss.xml",
	LAZIO:                 BASE_URL + "lazio/notizie/lazio_rss.xml",
	LIGURIA:               BASE_URL + "liguria/notizie/liguria_rss.xml",
	LOMBARDIA:             BASE_URL + "lombardia/notizie/lombardia_rss.xml",
	MARCHE:                BASE_URL + "marche/notizie/marche_rss.xml",
	MOLISE:                BASE_URL + "molise/notizie/molise_rss.xml",
	PIEMONTE:              BASE_URL + "piemonte/notizie/piemonte_rss.xml",
	PUGLIA:                BASE_URL + "puglia/notizie/puglia_rss.xml",
	SARDEGNA:              BASE_URL + "sardegna/notizie/sardegna_rss.xml",
	SICILIA:               BASE_URL + "sicilia/notizie/sicilia_rss.xml",
	TOSCANA:               BASE_URL + "toscana/notizie/toscana_rss.xml",
	TRENTINO_ALTO_ADIGE:   BASE_URL + "trentino/notizie/trentino_rss.xml",
	UMBRIA:                BASE_URL + "umbria/notizie/umbria_rss.xml",
	VALLE_AOSTA:           BASE_URL + "valledaosta/notizie/valledaosta_rss.xml",
	VENETO:                BASE_URL + "veneto/notizie/veneto_rss.xml"}
