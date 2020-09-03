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

	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate, spotify.ScopeUserLibraryRead)
	url := auth.AuthURL(state)

    fmt.Printf("%s %s", "Acesse a URl em :: ", url)
    fmt.Println("")

    http.HandleFunc("/callback", func (w http.ResponseWriter, r *http.Request) {
        fmt.Println("Recebido! :D") 
        fmt.Fprintf(w, "Autorizado com sucesso.")
    
    })
    go http.ListenAndServe(":7521", nil)

    fmt.Println("Esperando Spotify Client")
    
    client := <-ch
	fmt.Println(client)


	switch command {
	case "now":
		fmt.Println("O que está tocando agora: ")

	default:
		fmt.Println("Comando não implementado")
	}


}
