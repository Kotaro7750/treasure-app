package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/voyagegroup/treasure-app/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// https://firebase.google.com/docs/auth/admin/create-custom-tokens#sign_in_using_custom_tokens_on_clients
func main() {

	if len(os.Args) != 2 {
		log.Fatal("Need 1 argument but got ", len(os.Args)-1)
	}

	uid := os.Args[1]
	if len(uid) == 0 {
		log.Fatal("uid flag is missing.")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	client, err := util.InitAuthClient()
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	webapikey := os.Getenv("FIREBASE_WEB_API_KEY")
	endpoint := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", webapikey)

	if webapikey == "" {
		log.Fatalf("firebase Web API key is missing")
	}

	body := []byte(fmt.Sprintf(`
	{
		"token":"%s",
		"returnSecureToken":true
	}
	`, token))
	values := url.Values{}
	values.Set("returnSecureToken", "true")
	values.Set("token", token)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf( "error minting custom token: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBytes))
}