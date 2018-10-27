# GO meets Ansa

## Dependencies

Package github.com/julienschmidt/httprouter is used for request routing.

`go get https://github.com/julienschmidt/httprouter`

## Run application

`go run cmd/server/main.go`

And then make a request to http://localhost:8080/channels/Homepage, where Homepage is an example topic. 
You should be able to get an xml response corresponding to the RSS channel.

## Explanation

The application is divided in several parts. 

On the top level, `pkg` is where your source code resides, and cmd where your main.go goes.

Inside `fetching` package there is the logic related to 
business domain of fetching an RSS. There is nothing really inside this layer, only the models, because the app
is very simple. But as apps grow, logic tends to grow here. 

Inside `http` goes the http handler of the requests. Here you interact with the http request in terms of path params, 
query params, http status response code, etc. 

Inside `storage` goes the logic to actually retrieve the information from the database or from another API. I've hardcoded the urls for simplicity.


```
░░░░░░░░░░░░░░░░░░░░░░░░░░█▄
░▄▄▄▄▄▄░░░░░░░░░░░░░▄▄▄▄▄░░█▄
░▀▀▀▀▀███▄░░░░░░░▄███▀▀▀▀░░░█▄   This repository is
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░█   totally useless! (or maybe not)
░▄▀▀▀▀▀▄░░░░░░░░░░▄▀▀▀▀▀▄░░░░█   
█▄████▄░▀▄░░░░░░▄█░▄████▄▀▄░░█▄  But i use it
████▀▀██░▀▄░░░░▄▀▄██▀█▄▄█░█▄░░█   to learn Golang 
██▀██████░█░░░░█░████▀█▀██░█░░█
████▀▄▀█▀░█░░░░█░█████▄██▀▄▀░░█
███████▀░█░░░░░░█░█████▀░▄▀░░░█
░▀▄▄▄▄▄▀▀░░░░░░░░▀▀▄▄▄▄▀▀░░░░░█ Have pity on me ❤️
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░█
░░▓▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓▓░░░░█
░░░▓▓▓▓▓░░░░░░░░░░░░▓▓▓▓▓░░░░░█
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░█
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░█▀
░░░░░░░░░▄▄███████▄▄░░░░░░░░░█
░░░░░░░░█████████████░░░░░░░█▀
░░░░░░░░░▀█████████▀░░░░░░░█▀
░░░░░░░░░░░░░░░░░░░░░░░░░░█▀
░░░░░░░░░░░░░░░░░░░░░░░░░█▀
```