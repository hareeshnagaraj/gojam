/*
This file contains the wrapper functions for the Spotify Developer API
*/
package gojam
import (
  "fmt"
  "net/http"
  "net/url"
  // "io/ioutil"
)
var Constants = GetConstants();     //getting the object that contains our secret keys        

func RunThis(){
    fmt.Println("spotify_1 client ",Constants["spotify_client_ID"]);
}

/*
Establishing our connection with spotify and retrieving an authorization token
*/
func SpotifyAuthorizationRedirect(w http.ResponseWriter, r *http.Request){         //making the spotify api calls
  authString := GetSpotifyAuthorizationString();
  fmt.Println(authString);  
  http.Redirect(w, r, authString, http.StatusMovedPermanently);
}

/*
  Creating the authorization string
*/
func GetSpotifyAuthorizationString()string{
  var Url *url.URL
  Url, err := url.Parse(Constants["spotify_authorize"])
  if err != nil {
      panic("boom")
  };
  parameters := url.Values{}
  parameters.Add("client_id",Constants["spotify_client_ID"]);
  parameters.Add("redirect_uri",Constants["spotify_redirect_URI"]);
  parameters.Add("response_type","code");
  Url.RawQuery = parameters.Encode()
  fmt.Printf("Encoded URL is %q\n", Url.String())
  return Url.String();
}

/*
Parsing url and returning map of values
*/
func ParseUrl(a string) map[string][]string{
  values,err := url.ParseQuery(a);
  if(err != nil){
    panic("boom")
  }
  return values;
}
/*
Print values map
*/
func PrintValues(w http.ResponseWriter, values map[string][]string){
   for i, n := range values{
      fmt.Fprintf(w,"%s:%s\n", i, n)
    }
}