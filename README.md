# mySpotifyWrapped

This is a simple tool I've written to analyze my Spotify data. You can get your data [here](https://www.spotify.com/account/privacy/). Place the `StreamingHistory0.json` and `StreamingHistory1.json` in a directory called `MyData` in the source directory.

As of now the program only prints some interesting statistics about the listening history in the console, but the ultimate goal is to make this into a web page.

## How to run:

You can run the program using the `Makefile`

### Possible make options

`general` 		- Some general information like total listening time and number of unique artists
`artistUnique` 	- Prints statistics like most popular artist in number of songs listened
`artistTime` 	- Prints statistics like most popular artist in total time listened
`trackUnique` 	- Prints statistics like most popular song in number of unique listening occasions
`trackTime` 	- Prints statistics like most popular song in total time listened
`all` 			- Prints all the possible information

