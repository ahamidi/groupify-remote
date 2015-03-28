package main

import (
	"io/ioutil"
	"log"

	"github.com/op/go-libspotify/spotify"
)

// start of script tag
const ScriptStart = "tell application \"Spotify\" to "

var commands = map[string]string{
	"state":      "player state",
	"play":       "play",
	"pause":      "pause",
	"duration":   "duration of current track",
	"name":       "name of current track",
	"album":      "album of current track",
	"id":         "id of current track",
	"artwork":    "artwork of current track",
	"vol_loud":   "set sound volume to 100",
	"vol_soft":   "set sound volume to 20",
	"vol_norm":   "set sound volume to 50",
	"set_volume": "set sound volume to ", //requires parameter
	"play_track": "play track ",          //requires parameter
	"position":   "player position",
}

func newSpotifySession(audio *audioWriter) *spotify.Session {
	appKey, err := ioutil.ReadFile("spotify_appkey.key")
	if err != nil {
		log.Fatal(err)
	}

	session, err := spotify.NewSession(&spotify.Config{
		ApplicationKey:   appKey,
		ApplicationName:  "Groupify Remote",
		CacheLocation:    "tmp",
		SettingsLocation: "tmp",
		AudioConsumer:    audio,

		// Disable playlists to make playback faster
		DisablePlaylistMetadataCache: true,
		InitiallyUnloadPlaylists:     true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return session
}

func getTrackFromURI(session *spotify.Session, uri string) *spotify.Track {
	// Parse the track
	link, err := session.ParseLink(uri)
	if err != nil {
		log.Fatal(err)
	}
	track, err := link.Track()
	if err != nil {
		log.Fatal(err)
	}

	return track
}

func callSpotify(command string, param string) string {
	//fullcmd := ScriptStart + commands[command] + param

	//out, err := exec.Command("/usr/bin/osascript", "-e", fullcmd).Output()
	//if err != nil {
	//log.Fatal(err)
	//log.Fatal(out)
	//}
	//return strings.TrimSpace(string(out))

	return "fixme"
}
