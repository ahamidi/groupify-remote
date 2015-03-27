package main

import (
	"log"
	"strconv"
	"time"
)

// polling sleep time
const sleepTime = time.Second / 2

func polling(s chan interface{}) {
	playerState := getPlayerState()
	track := getCurrentTrackID()
	timeLeft := int(getTimeLeft())
	getNextSong := true
	//inner for loop variables
	var (
		currentPlayerState string
		currentTimeLeft    int
		currentTrack       string
	)
	log.Println("starting player state: ", playerState)
	for {
		time.Sleep(sleepTime)
		//check player state
		currentPlayerState = getPlayerState()
		currentTimeLeft = int(getTimeLeft())
		currentTrack = getCurrentTrackID()
		// log.Println(currentTrack)
		if playerState != currentPlayerState {
			message := NotificationMessage{"player_" + currentPlayerState, "", currentTrack}
			s <- message
			playerState = currentPlayerState
			// log.Println("player state changed: ", currentPlayerState)
		}
		if currentTrack != track {
			if !getNextSong {
				s <- NotificationMessage{"track_ending", track, currentTrack}
				s <- NotificationMessage{"track_start", nextTrack, nextTrack}
				getNextSong = true
				nextTrack := getNextTrack()
				if nextTrack != "" {
					setCurrentTrack(nextTrack)
					track = nextTrack
				} else {
					track = currentTrack
				}
				setNextTrack("")
			}
		}
		//check player duration - is track over
		if currentTimeLeft != timeLeft {
			// log.Println("New Time : ", currentTimeLeft)
			timeLeft = currentTimeLeft
			message := NotificationMessage{"time_left", strconv.Itoa(timeLeft), currentTrack}
			s <- message
			if timeLeft < 30 && getNextSong { //lock out period
				getNextSong = false
				message := NotificationMessage{"get_next_track", track, currentTrack}
				s <- message
			}
		}

	}
}
