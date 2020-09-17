package cli

import (
	"fmt"

	"github.com/zmb3/spotify"
)

// HandlerSpotify handler the spotify client
type HandlerSpotify struct {
	client *spotify.Client
}

// Now - shows whats playing
func (h *HandlerSpotify) Now() {
	currentlyPlaying, currentlyPlayingError := h.client.PlayerCurrentlyPlaying()
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
}

// User - shows users data
func (h *HandlerSpotify) User() {
	user, userErr := h.client.CurrentUser()
	if userErr != nil {
		fmt.Println(userErr)
		return
	}
	fmt.Printf("%s <%s> [%s]\n", user.DisplayName, user.Product, user.ID)
}
