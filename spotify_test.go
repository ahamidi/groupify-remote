package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/op/go-libspotify/spotify"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	os.Exit(m.Run())
}

func TestSpotifyPlayTrack(t *testing.T) {

	audio, _ := newAudioWriter()

	session := newSpotifySession(audio)

	credentials := spotify.Credentials{
		Username: os.Getenv("SPOTIFY_USERNAME"),
		Password: os.Getenv("SPOTIFY_PASSWORD"),
	}
	if err := session.Login(credentials, false); err != nil {
		log.Fatal(err)
	}

	uri := "spotify:track:5R0w7bVKJTeDltxIwkLpSZ"

	link, err := session.ParseLink(uri)
	if err != nil {
		log.Fatal(err)
	}
	track, err := link.Track()
	if err != nil {
		log.Fatal(err)
	}
	track.Wait()

	player := session.Player()
	if err := player.Load(track); err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	player.Play()

	assert.NotNil(t, track)
	assert.Nil(t, err)
}
