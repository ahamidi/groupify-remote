package main

import (
	"flag"
	"log"
	"strings"

	"github.com/op/go-libspotify/spotify"
)

// API Host Flag
var apiHost = flag.String("host", "localhost:8080",
	"Hostname of Groupify API")

// Spotify Username
var username = flag.String("username", "",
	"Spotify Username")

// Spotify Password
var password = flag.String("password", "",
	"Spotify Password")

func main() {
	log.Println("Starting Groupify Remote")

	flag.Parse()

	// Create host string
	url := strings.Join([]string{"ws://", *apiHost, "/api/v1/remote"}, "")

	// Create WS Connection
	c, err := NewWSConnection(url, "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}

	// Create new portaudio writer
	audio, err := newAudioWriter()
	if err != nil {
		log.Fatal(err)
	}

	// Create Spotify Session
	session := newSpotifySession(audio)

	// Login to Spotify
	credentials := spotify.Credentials{
		Username: *username,
		Password: *password,
	}
	if err = session.Login(credentials, false); err != nil {
		log.Fatal(err)
	}

	// Get track
	track := getTrackFromURI(session, "spotify:track:4yoirlyne2EwkftLG7CpvN")

	track.Wait()
	log.Println("Track Name:", track.Name())

	// Spotify status queue
	//spotifyState := make(chan interface{})
	//go polling(spotifyState)

	// Create notification message channel
	notifications := make(chan interface{})
	go WSMessageReceiver(c, notifications)

	//go processQueue(notifications)

	//for s := range spotifyState {
	//log.Println("Spotify State:", s)
	//SendWSMessage(c, s)
	//}

	done := make(chan bool, 1)
	<-done
}
