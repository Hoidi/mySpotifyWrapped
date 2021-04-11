package main

import (
	"encoding/json"
	"io/ioutil"
)

func Setup() {
	content0, _ := ioutil.ReadFile("MyData/StreamingHistory0.json")
	content1, _ := ioutil.ReadFile("MyData/StreamingHistory1.json")


	// add the two jsons together
	var songs0 []SongPlayJSON
	var songs1 []SongPlayJSON
	_ = json.Unmarshal(content0, &songs0)
	_ = json.Unmarshal(content1, &songs1)
	songs = append(songs0, songs1...)

	artistsSongs = createArtistsSongs()
	songUniqueListens = createSongUniqueListens()
	songTimeListens = createSongTimeListens()
}

func createArtistsSongs() map[string][]string {
	artistsSongs = map[string][]string{}
	for _,song := range songs {
		artistsSongs[song.Artist] = append(artistsSongs[song.Artist], song.Track)
	}
	// remove Unique
	for key, _ := range artistsSongs {
		artistsSongs[key] = Unique(artistsSongs[key])
	}
	return artistsSongs
}

func createSongUniqueListens() map[string]int {
	songUniqueListens = map[string]int{}
	for _,song := range songs {
		songUniqueListens[song.Track] = songUniqueListens[song.Track] + 1
	}
	return songUniqueListens
}

func createSongTimeListens() map[string]int {
	songTimeListens = map[string]int{}
	for _,song := range songs {
		songTimeListens[song.Track] = songTimeListens[song.Track] + song.Time
	}
	return songTimeListens
}

