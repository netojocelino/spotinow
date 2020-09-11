package cli

import (
	"fmt"

	"github.com/netojocelino/spotinow/authenticate"
)

// HandlerCommandLineInput handler with command and do magic
func HandlerCommandLineInput(command string) {

	client := authenticate.GetClient()

	switch command {
	case "now":
		currentlyPlaying, currentlyPlayingError := client.PlayerCurrentlyPlaying()
		if (currentlyPlayingError != nil) || (currentlyPlaying.Item == nil) {
			fmt.Printf("%s\n", "Não foi identificado nenhuma música sendo reproduzida.")
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

	case "user":
		user, userErr := client.CurrentUser()
		if userErr != nil {
			fmt.Println(userErr)
			return
		}
		fmt.Printf("%s <%s> [%s]\n", user.DisplayName, user.Product, user.ID)
	default:
		fmt.Println("Comando não implementado")
	}
}
