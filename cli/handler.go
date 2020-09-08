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
		fmt.Printf("Tocando: %s - %s \n", currentlyPlaying.Item.Name, currentlyPlaying.Item.Artists[0].Name)

	default:
		fmt.Println("Comando não implementado")
	}
}
