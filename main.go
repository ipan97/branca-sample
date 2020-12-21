package main

import (
	"fmt"
	"github.com/hako/branca"
)

func main() {
	b := branca.NewBranca("supersecretkeyyoushouldnotcommit") // This key must be exactly 32 bytes long.

	// Encode String to Branca Token.
	token, err := b.EncodeToString(`{"name":"Facebook","signature_url":"https://facebook.com"}`)
	if err != nil {
		fmt.Println(err)
	}

	//  b.SetTTL(3600) // Uncomment this to set an expiration (or ttl) of the token (in seconds).
	//token = "87y8daMzSkn7PA7JsvrTT0JUq1OhCjw9K8w2eyY99DKru9FrVKMfeXWW8yB42C7u0I6jNhOdL5ZqL" // This token will be not allowed if a ttl is set.

	// Decode Branca Token.
	message, err := b.DecodeToString(token)
	if err != nil {
		fmt.Println(err) // token is expired.
		return
	}
	fmt.Println(token)   // 87y8da....
	fmt.Println(message) // message
}
