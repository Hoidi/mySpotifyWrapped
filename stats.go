package main

func TotalTimeListened() float64 {
	var time int
	for _,song := range songs {
		time+= song.Time
	}
	return float64(((time / 1000)/60)/60)
}

func MostPopularArtist(unique bool) (string,int) {
	var mostPop string
	var listens int
	for artist,songList := range artistsSongs {
		artistsListens := 0
		for _, song := range songList {
			if unique {
				artistsListens += songUniqueListens[song]
			} else {
				artistsListens += songTimeListens[song]
			}

		}
		if artistsListens > listens {
			listens = artistsListens
			mostPop = artist
		}
	}

	if unique {
		return mostPop, listens
	} else {
		return mostPop, ((listens/1000)/60)/60 // return time in hours
	}
}

func MostPopularTrack(unique bool) (string,int) {
	var mostPop string
	var listening int

	if unique {
		for key,value := range songUniqueListens {
			if value > listening {
				listening = value
				mostPop = key
			}
		}
	} else {
		for key,value := range songTimeListens {
			if value > listening {
				listening = value
				mostPop = key
			}
		}
	}


	if unique {
		return mostPop, listening
	} else {
		return mostPop, ((listening/1000)/60)/60 // return time in hours
	}
}

func HighestSteakTrack(unique bool) (string, int) {
	var highestStreakNum int
	var highestStreakTrack string
	var currentStreakNum int
	var currentStreakTrack string
	for _,song := range songs {
		if song.Track == currentStreakTrack {
			if unique {
				currentStreakNum += 1
			} else {
				currentStreakNum += song.Time
			}
		} else {
			currentStreakNum = 0
		}

		if currentStreakNum > highestStreakNum {
			highestStreakNum = currentStreakNum
			highestStreakTrack = currentStreakTrack
		}

		currentStreakTrack = song.Track
	}

	if unique {
		return highestStreakTrack, highestStreakNum
	} else {
		return highestStreakTrack, (highestStreakNum/1000)/60 // return time in minutes
	}
}


// TODO: Lots of code duplication in HighestSteakArtist and HighestSteakTrack. Merge them somehow.
func HighestSteakArtist(unique bool) (string, int) {
	var highestStreakNum int
	var highestStreakArtist string
	var currentStreakNum int
	var currentStreakArtist string
	for _,song := range songs {
		if song.Artist == currentStreakArtist {
			if unique {
				currentStreakNum += 1
			} else {
				currentStreakNum += song.Time
			}
		} else {
			currentStreakNum = 0
		}

		if currentStreakNum > highestStreakNum {
			highestStreakNum = currentStreakNum
			highestStreakArtist = currentStreakArtist
		}

		currentStreakArtist = song.Artist
	}

	if unique {
		return highestStreakArtist, highestStreakNum
	} else {
		return highestStreakArtist, (highestStreakNum/1000)/60 // return time in minutes
	}
}