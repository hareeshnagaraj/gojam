/*
This file contains the wrapper functions for the Spotify Developer API
The details of this process can be found at https://developer.spotify.com/web-api/authorization-guide/#authorization-code-flow

*/
package gojam
import (
  "fmt"
  "net/http"
  "net/url"
  "bytes"
  // "appengine/urlfetch"
  // "strconv"
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
  http.Redirect(w, r, authString, http.StatusMovedPermanently);
}
/*
Retrieving the refresh and access tokens from the Spotify API
*/
func SpotifyTokenRequest(w http.ResponseWriter,r *http.Request){
  tokenString := Values["spotify_token"];
  data := url.Values{}
  data.Add("grant_type","authorization_code");
  data.Add("client_secret",Values["spotify_client_secret"]);
  data.Add("client_id",Values["spotify_client_ID"]);  
  data.Add("redirect_uri",Values["spotify_redirect_URI"]);
  data.Add("code",Values["spotify_code"]);

  fmt.Fprintf(w, "spotifyToken url  %v\n\n",tokenString);

  client := &http.Client{};
  req, err := http.NewRequest("POST", tokenString, bytes.NewBufferString(data.Encode())); // <-- URL-encoded payload
  req.Close = true;     
  resp, err := client.Do(req);
  if err != nil {
      fmt.Fprintf(w, "error ",err);
      // panic("boom2")
  }else{
      defer resp.Body.Close(); 
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
  };
  parameters := url.Values{}
  parameters.Add("client_id",Constants["spotify_client_ID"]);
  parameters.Add("redirect_uri",Constants["spotify_redirect_URI"]);
  parameters.Add("response_type","code");
  Url.RawQuery = parameters.Encode()
  return Url.String();
}

/*
Handling the initial authentication response
*/
func StoreAuthResponse(code string){
  SetVal("spotify_code",code);
}

/*
Parsing url and returning map of values
*/
func ParseUrl(urlstring string) map[string][]string{
  values,err := url.ParseQuery(urlstring);
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