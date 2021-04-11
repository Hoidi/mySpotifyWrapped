package main

import (
	"fmt"
	"os"
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

	args := os.Args

	// TODO: This is pretty ugly code
	switch args[1]{
	case "all":
		printGeneral()
		printArtistUnique()
		printArtistTime()
		printTrackUnique()
		printTrackTime()
	case "general":
		printGeneral()
	case "artist":
		if args[2] == "unique" {
			printArtistUnique()
		} else if args[2] == "time" {
			printArtistTime()
		}
	case "track":
		if args[2] == "unique" {
			printTrackUnique()
		} else if args[2] == "time" {
			printTrackTime()
		}
	default:
		fmt.Print("Something went wrong")
	}

}

func printGeneral() {
	fmt.Print(TotalTimeListened()," hours of total listening time.\n\n")

	uniqueArtists, uniqueTracks := NumberOfUnique()
	fmt.Print("You listened to ", uniqueArtists, " unique artist and ",uniqueTracks, " unique tracks.\n\n")


	printHours(true)
}

func printArtistTime() {
	popArtist, popArtistListens := MostPopularArtist(false)
	fmt.Print(popArtist," is the artist you've listened to the most time. Total listening time is ", popArtistListens," hours.\n\n")

	streakArtistTime, streakArtistNumTime := HighestSteakArtist(false)
	fmt.Print(streakArtistTime, " is the artist with the longest streak in a row. The streak time was ", streakArtistNumTime, " minutes.\n\n")

	printTopN(10,"tracks",false)
}

func printArtistUnique() {
	popArtist, popArtistListens := MostPopularArtist(true)
	fmt.Print(popArtist," is the artist you've listened to the most times. Total listening times is ", popArtistListens,".\n\n")

	streakArtist, streakArtistNum := HighestSteakArtist(true)
	fmt.Print(streakArtist, " is the artist with the highest streak in a row. The streak was ", streakArtistNum,"\n\n")

	printTopN(10,"tracks",true)
}

func printTrackTime() {
	popTrack, popTrackListens := MostPopularTrack(false)
	fmt.Print(popTrack," is the track you've listened to the most time. Total listening time is ", popTrackListens," hours.\n\n")

	streakTrackTime, streakNumTime := HighestSteakTrack(false)
	fmt.Print(streakTrackTime, " is the song with the longest streak in a row. The streak time was ", streakNumTime, " minutes.\n\n")

	printTopN(10,"tracks",false)
}

func printTrackUnique() {
	popTrack, popTrackListens := MostPopularTrack(true)
	fmt.Print(popTrack," is the track you've listened to the most times. Total listening times is ", popTrackListens,".\n\n")

	streakTrack, streakTrackNum := HighestSteakTrack(true)
	fmt.Print(streakTrack, " is the song with the highest streak in a row. The streak was ", streakTrackNum,"\n\n")


	printTopN(10,"tracks",true)
}

func printTopN(n int, a string, unique bool) {
	var topList PairList
	if a == "artists" {
		topList = TopArtist(unique)
	} else {
		topList = TopTracks(unique)
	}
	fmt.Print("Your top ", n, " ", a,".\n")
	for i, item := range topList[:n] {
		if unique {
			fmt.Print(i+1,". ", item.Key, ": ", item.Value, " minutes. (", item.Value/60," hours)\n")
		} else {
			fmt.Print(i+1,". ", item.Key, ": ", item.Value, " times.\n")
		}

	}
	fmt.Print("\n")
}

func printHours(unique bool) {
	var hours [24]int = Hour(unique)

	for i,j := range hours {
		fmt.Print(i," ",j,"\n")
	}
	fmt.Print("\n")
}
