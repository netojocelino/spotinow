package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/zmb3/spotify"
)

var (
	redirectURL = "http://localhost:7521/callback"
	state       = "login" // TODO Gerar chave dinamica
	ch          = make( chan *spotify.Client )
)



func main () {

	if len(os.Args) < 2 {
		fmt.Println("Welcome")
		fmt.Println("Spotify CLI - Saída o que está ouvindo.")
		return
	}

	command := os.Args[1]

    auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
        spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying)
	url := auth.AuthURL(state)

    fmt.Printf("%s %s\n\n", "Acesse a URl em :: ", url)

    http.HandleFunc("/callback", func (w http.ResponseWriter, r *http.Request) {
        token, err := auth.Token(state, r)
        
        if(err != nil) {
            fmt.Printf("Erro ao gerar token: %s\n", err)
            http.Error(w, "Não foi possível retornar o token\n", http.StatusNotFound)
        } else {
            fmt.Fprintf(w, "Autorizado com sucesso.\n")
            client := auth.NewClient(token)
            ch <- &client
        }

    })
    go http.ListenAndServe(":7521", nil)

    
    client := <-ch

	switch command {
	case "now":
        currentlyPlaying, currentlyPlayingError := client.PlayerCurrentlyPlaying()
        if (currentlyPlayingError != nil) {
            fmt.Println(currentlyPlayingError)
            return            
        }
        fmt.Printf("Tocando: %s - %s", currentlyPlaying.Item.Name, currentlyPlaying.Item.Artists[0].Name)

	default:
		fmt.Println("Comando não implementado")
	}


}
