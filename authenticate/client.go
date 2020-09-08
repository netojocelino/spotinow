package authenticate

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify"
)

var (
	redirectURL = "http://localhost:7521/callback"
	state       = "login" // TODO Gerar chave dinamica
	ch          = make(chan *spotify.Client)
)

// Authenticate using OAuth2 and return *spotify.Client
func GetClient() *spotify.Client {
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Token(state, r)

		if err != nil {
			fmt.Printf("Erro ao gerar token: %s\n", err)
			http.Error(w, "Não foi possível retornar o token\n", http.StatusNotFound)
		} else {
			fmt.Fprintf(w, "Autorizado com sucesso.\n")
			client := auth.NewClient(token)
			ch <- &client
		}

	})
	go http.ListenAndServe(":7521", nil)

	url := auth.AuthURL(state)
	fmt.Printf("%s %s\n\n", "Acesse a URl em :: ", url)
	// go openURL(url)

	client := <-ch

	return client
}
