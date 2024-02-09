// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func send() {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody("Hi there, Prateek is testing Twilio!")
	params.SetFrom("+12135663883")
	params.SetTo("+919521427560")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func main(){
	send()
}