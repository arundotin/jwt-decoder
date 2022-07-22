package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	accessToken := grabAccessToken()
	headerAndPayLoad := getHeaderAndPayLoadFrom(accessToken)

	printDecodedHeaderAndPayload(headerAndPayLoad)

}

func printDecodedHeaderAndPayload(headerAndPayLoad []string) {

	headerDecoded, err := base64.RawStdEncoding.DecodeString(headerAndPayLoad[0])

	if err != nil {
		fmt.Println("Not a base64 encoded string ", headerAndPayLoad[0])
		os.Exit(1)
	}

	var prettyJSONHeader bytes.Buffer

	json.Indent(&prettyJSONHeader, headerDecoded, "", "\t")

	fmt.Println("------------------------------")
	fmt.Println("Token Header")
	fmt.Println("------------------------------")

	fmt.Println(string(prettyJSONHeader.Bytes()))
	fmt.Println("------------------------------")

	payloadDecoded, err := base64.RawStdEncoding.DecodeString(headerAndPayLoad[1])

	if err != nil {
		fmt.Println("Not a base64 encoded string ", headerAndPayLoad[1])
		os.Exit(1)
	}

	var prettyJSONPayload bytes.Buffer

	json.Indent(&prettyJSONPayload, payloadDecoded, "", "\t")

	fmt.Println("------------------------------")
	fmt.Println("Token payload")
	fmt.Println("------------------------------")

	fmt.Println(string(prettyJSONPayload.Bytes()))

}

func getHeaderAndPayLoadFrom(accessToken string) []string {
	splittedAccessToken := strings.Split(accessToken, ".")

	if len(splittedAccessToken) != 3 {
		fmt.Println("Invalid Access Token")
		os.Exit(1)
	}

	return splittedAccessToken[0:2]
}

func grabAccessToken() string {
	accessToken := flag.String("accessToken", "", "Gets the Decoded Access Token")

	flag.Parse()

	if *accessToken == "" {
		fmt.Println("--accessToken flag is mandatory")
		os.Exit(1)
	}

	return *accessToken

}
