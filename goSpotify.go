/*
This file contains the wrapper functions for the Spotify Developer API
*/
package gojam
import (
  "fmt"
  "net/http"
  "net/url"
  "io/ioutil"
)
var Constants = GetConstants();     //getting the object that contains our secret keys        

func RunThis(){
    fmt.Println("spotify_1 client ",Constants["spotify_client_ID"]);
}

/*
Establishing our connection with spotify and retrieving an authorization token
*/
func ConnectSpotify(){         //making the spotify api calls
  authString := GetSpotifyAuthorizationString();
  resp, err := http.Get(authString);
  if err != nil {
    // handle error
    fmt.Println("error");
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)      //body is returned as a byte array
  if err == nil {
    fmt.Println("ConnectSpotify success: ",string(body))
  }else{
    fmt.Println("ConnectSpotify error")
  }
}

/*
  Creating the authorization string
*/
func GetSpotifyAuthorizationString()string{
  var Url *url.URL
  Url, err := url.Parse(Constants["spotify_authorize"])
  if err != nil {
      panic("boom")
  }
  parameters := url.Values{}
  parameters.Add("client_id",Constants["spotify_client_ID"]);
  parameters.Add("response_type","code");
  parameters.Add("redirect_uri",Constants["spotify_redirect_URI"]);
  Url.RawQuery = parameters.Encode()
  fmt.Printf("Encoded URL is %q\n", Url.String())
  return Url.String();
}