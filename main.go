package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SongPlayJSON struct {
	When 	string		`json:"endTime"`
	Artist	string		`json:"artistName"`
	Track 	string		`json:"trackName"`
	Time	int			`json:"msPlayed"`
}

var songs []SongPlayJSON				// all streaming history
var artistsSongs map[string][]string	// all artists and their respective songs
var songUniqueListens map[string]int	// all songs and how many times they've been listened to
var songTimeListens map[string]int		// all songs and how many milli seconds they've been listened to

func main() {
	setup()

	fmt.Print(totalTimeListened()," hours of total listening time.\n")
func setup() {
	content0, _ := ioutil.ReadFile("MyData/StreamingHistory0.json")
	content1, _ := ioutil.ReadFile("MyData/StreamingHistory0.json")


	// add the two jsons together
	songJson0 := content0
	songJson1 := content1
	var songs0 []SongPlayJSON
	var songs1 []SongPlayJSON
	_ = json.Unmarshal(songJson0, &songs0)
	_ = json.Unmarshal(songJson1, &songs1)
	songs = append(songs0, songs1...)
	
	createArtistsSongs()
	createSongUniqueListens()
	createSongTimeListens()
}

func createArtistsSongs()  {
	artistsSongs = map[string][]string{}
	for _,song := range songs {
		artistsSongs[song.Artist] = append(artistsSongs[song.Artist], song.Track)
	}
	// remove unique
	for key, _ := range artistsSongs {
		artistsSongs[key] = unique(artistsSongs[key])
	}
}

func createSongUniqueListens() {
	songUniqueListens = map[string]int{}
	for _,song := range songs {
		songUniqueListens[song.Track] = songUniqueListens[song.Track] + 1 
	}
}

func createSongTimeListens() {
	songTimeListens = map[string]int{}
	for _,song := range songs {
		songTimeListens[song.Track] = songTimeListens[song.Track] + song.Time
	}
}
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}