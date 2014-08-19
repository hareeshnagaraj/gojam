package gojam

import (
  "fmt"
  "net/http"
)

func init() {
    http.HandleFunc("/spotify/authorization", spotifyAuthorizationHandler)   //Getting the user's spotify authorization tokens
    http.HandleFunc("/spotify/URI", spotifyURIHandler)             //Handling the URI redirect

}

/*
Redirecting to spotify after generating the authentication string basd on the values on gojam_constants
*/
func spotifyAuthorizationHandler(w http.ResponseWriter, r *http.Request) {  
    SpotifyAuthorizationRedirect(w,r);
}

/*
Handling the response tokens from spotify's redirect
*/
func spotifyURIHandler(w http.ResponseWriter, r *http.Request) {  
    fmt.Println("SPOTIFY URI, %v","ss");
}