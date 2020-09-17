package cli

import (
	"fmt"

	"github.com/netojocelino/spotinow/authenticate"
	"github.com/zmb3/spotify"
)

// CommandHandler handler a command from the CLI
type CommandHandler interface {
	HandleCommand()
}

// PlayingNow show whats playing in current moment
type PlayingNow struct {
	client *spotify.Client
}

// UserProfile show current user data
type UserProfile struct {
	client *spotify.Client
}

// HandleCommand make something
func (h *PlayingNow) HandleCommand() {
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

// HandleCommand make something
func (h *UserProfile) HandleCommand() {
	user, userErr := h.client.CurrentUser()
	if userErr != nil {
		fmt.Println(userErr)
		return
	}
	fmt.Printf("%s <%s> [%s]\n", user.DisplayName, user.Product, user.ID)
}

// HandlerCommandLineInput handler with command and do magic
func HandlerCommandLineInput(command string) {

	client := authenticate.GetClient()

	switch command {
	case "now":
		playingNow := PlayingNow{
			client: client,
		}
		playingNow.HandleCommand()

	case "user":
		currentUser := UserProfile{
			client: client,
		}
		currentUser.HandleCommand()
	default:
		fmt.Println("Comando não implementado")
	}
}
