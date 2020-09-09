package authenticate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	redirectURL  = "http://localhost:7521/callback"
	state        = "login" // TODO Gerar chave dinamica
	ch           = make(chan *spotify.Client)
	jsonFileName = "SpotiNowToken.json"
)

// GetClient using OAuth2 and return *spotify.Client
func GetClient() *spotify.Client {
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate,
		spotify.ScopeUserLibraryRead, spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadEmail)

	fOpen, fOpenErr := os.Open(jsonFileName)

	if fOpenErr != nil {

		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.Token(state, r)

			if err != nil {
				fmt.Printf("Erro ao gerar token: %s\n", err)
				http.Error(w, "Não foi possível retornar o token\n", http.StatusNotFound)
			} else {

				f, fError := os.Create(jsonFileName)

				if fError != nil {
					fmt.Printf("Ocorreu um erro ao gerar repositório do token. %s", fError.Error())
				}

				enc := json.NewEncoder(f)
				encError := enc.Encode(token)

				if encError != nil {
					fmt.Printf("Ocorreu um erro ao salvar token. %s", encError.Error())
				}
				f.Close()
				fmt.Fprintf(w, "Autorizado com sucesso.\n")
				client := auth.NewClient(token)
				ch <- &client
			}

		})
		go http.ListenAndServe(":7521", nil)

		url := auth.AuthURL(state)
		fmt.Printf("%s %s\n\n", "Acesse a URl em :: ", url)
		// go openURL(url)

	} else {

		enc := json.NewDecoder(fOpen)

		var token *oauth2.Token
		encErr := enc.Decode(&token)

		if encErr != nil {
			fmt.Printf("Erro a abrir json. %s", encErr.Error())
		}

		fOpen.Close()
		client := auth.NewClient(token)

		return &client
	}

	client := <-ch

	return client
}
