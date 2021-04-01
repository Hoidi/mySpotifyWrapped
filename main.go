package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type SongPlayJSON struct {
	When 	string		`json:"endTime"`
	Artist	string		`json:"artistName"`
	Track 	string		`json:"trackName"`
	Time	int			`json:"msPlayed"`
}

func main() {
	content, err := ioutil.ReadFile("MyData/StreamingHistory0.json")
	if err != nil {
		log.Fatal(err)
	}

	songJson := content
	var songs []SongPlayJSON
	_ = json.Unmarshal(songJson, &songs)

	for i := 0; i < len(songs); i++ {
		fmt.Print(songs[i].Track + "\n")
	}
}
