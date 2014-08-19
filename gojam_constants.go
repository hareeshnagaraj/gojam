/*

Constants contains the important key information for your API calls, this is an extremely sensitive file

*/
package gojam

/*

Values is the map that contains the relevant information for each

spotify_authorize - do not modify, as this is defined by spotify
spotify_client_ID - uniquely assigned to your application, find under 'my applications' at developer.spotify.com
spotify_client_Secret - uniquely assigned to your application, find under 'my applications' at developer.spotify.com
spotify_redirect_URI - the URI which spotify redirects to upon authentication, with the token as well

*/
var Values = map[string]string{
  "spotify_authorize":"https://accounts.spotify.com/authorize", 
  "spotify_client_ID":"b5e44b2ab18f4f6686d0a7aefd217788",
  "spotify_client_Secret":"6ee513b2a1d8459f8d8b5ed4bb7bc305",
  "spotify_redirect_URI":"http://localhost:8080/spotifyURI/",
}

func GetConstants() map[string]string{
  return Values;
};