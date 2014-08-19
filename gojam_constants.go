/*

Constants contains the important key information for your API calls, this is an extremely sensitive file

*/
package gojam
import (
  // "fmt"
)
var Values = map[string]string{
  "spotify_authorize":"https://accounts.spotify.com/authorize",
  "spotify_client_ID":"b5e44b2ab18f4f6686d0a7aefd217788",
  "spotify_client_Secret":"6ee513b2a1d8459f8d8b5ed4bb7bc305",
  "spotify_redirect_URI":"http://localhost:8080/",
}

func GetConstants() map[string]string{
  return Values;
};