package cli

import (
	"fmt"

	"github.com/netojocelino/spotinow/authenticate"
)

// HandlerCommandLineInput handler with command and do magic
func HandlerCommandLineInput(command string) {

	client := authenticate.GetClient()

	handlerfy := HandlerSpotify{
		client: client,
	}

	switch command {
	case "now":
		handlerfy.Now()

	case "user":
		handlerfy.User()
	default:
		fmt.Println("Comando não implementado")
	}
}
