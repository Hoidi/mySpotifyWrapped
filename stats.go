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
