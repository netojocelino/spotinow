package cli

import (
	"fmt"

	"github.com/zmb3/spotify"
)

// HandlerCommandLineInput handler with command and do magic
func HandlerCommandLineInput(client *spotify.Client, command string) {

	switch command {
	case "now":
		currentlyPlaying, currentlyPlayingError := client.PlayerCurrentlyPlaying()
		if currentlyPlayingError != nil {
			fmt.Println(currentlyPlayingError)
			return
		}

		artistsName := "- "

		artistsList := currentlyPlaying.Item.Artists
		artistsListLen := len(artistsList)

		for indexArtist, artist := range artistsList {
			artistsName += artist.Name
			if indexArtist < (artistsListLen - 2) {
				artistsName += ", "
			} else if indexArtist < (artistsListLen - 1) {
				artistsName += " e "
			} else {
				artistsName += "."
			}
		}
		fmt.Printf("Tocando: %s %s \n", currentlyPlaying.Item.Name, artistsName)

	default:
		fmt.Println("Comando não implementado")
	}
}
