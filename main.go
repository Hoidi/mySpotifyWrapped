package main

import (
	"fmt"
)

type SongPlayJSON struct {
	When   string `json:"endTime"`
	Artist string `json:"artistName"`
	Track  string `json:"trackName"`
	Time   int    `json:"msPlayed"`
}

var songs []SongPlayJSON				// all streaming history
var artistsSongs map[string][]string	// all artists and their respective songs
var songUniqueListens map[string]int	// all songs and how many times they've been listened to
var songTimeListens map[string]int		// all songs and how many milli seconds they've been listened to

func main() {
	Setup()

	fmt.Print(TotalTimeListened()," hours of total listening time.\n\n")

	uniqueArtists, uniqueTracks := NumberOfUnique()
	fmt.Print("You listened to ", uniqueArtists, " unique artist and ",uniqueTracks, " unique tracks.\n\n")

	popArtist, popArtistListens := MostPopularArtist(true)
	fmt.Print(popArtist," is the artist you've listened to the most times. Total listening times is ", popArtistListens,".\n")
	popArtist, popArtistListens = MostPopularArtist(false)
	fmt.Print(popArtist," is the artist you've listened to the most time. Total listening time is ", popArtistListens," hours.\n\n")

	popTrack, popTrackListens := MostPopularTrack(true)
	fmt.Print(popTrack," is the track you've listened to the most times. Total listening times is ", popTrackListens,".\n")
	popTrack, popTrackListens = MostPopularTrack(false)
	fmt.Print(popTrack," is the track you've listened to the most time. Total listening time is ", popTrackListens," hours.\n\n")

	streakArtist, streakNum := HighestSteakTrack(true)
	fmt.Print(streakArtist, " is the song with the highest streak in a row. The streak was ", streakNum,"\n")
	streakTrackTime, streakNumTime := HighestSteakTrack(false)
	fmt.Print(streakTrackTime, " is the song with the longest streak in a row. The streak time was ", streakNumTime, " minutes.\n\n")

	streakArtist, streakArtistNum := HighestSteakArtist(true)
	fmt.Print(streakArtist, " is the song with the highest streak in a row. The streak was ", streakArtistNum,"\n")
	streakArtistTime, streakArtistNumTime := HighestSteakArtist(false)
	fmt.Print(streakArtistTime, " is the song with the longest streak in a row. The streak time was ", streakArtistNumTime, " minutes.\n\n")

	printTopNTracks(10,TopTracks(false))
}


func printTopNTracks(n int, topTracks PairList) {
	fmt.Print("Your top ", n, " tracks.\n")
	for i, track := range topTracks[:n] {
		fmt.Print(i+1,". ", track.Key, ": ", track.Value, " times.\n")
	}
}
