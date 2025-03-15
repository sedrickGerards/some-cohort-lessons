package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	// "golang.org/x/text/number"
)

type Response struct {
	Reference string `json:"reference"`
	Ussd_Code string `json:"ussd_code"`
}

var momoNumber, amount, description, ref string

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return fmt.Errorf("failed to load env file: %w", err)
		}
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return fmt.Errorf("API_KEY is not set")
	}

	// Your code here

	fmt.Println("Amount: ")
	fmt.Scanln(&amount)

	fmt.Println("Phone number: ")
	fmt.Scanln(&momoNumber)

	fmt.Println("Description: ")
	fmt.Scanln(&description)

	fmt.Println("Reference: ")
	fmt.Scanln(&ref)

	postBody, _ := json.Marshal(map[string]string{
		"amount":             amount,
		"from":               momoNumber,
		"description":        description,
		"external_reference": ref,
	})

	responseBody := bytes.NewBuffer(postBody)

	//GO HTTP post request

	resp, err := http.NewRequest(http.MethodPost, "https://demo.campay.net/api/collect/", responseBody)
	resp.Header.Set("Authorization", "Token "+apiKey)
	resp.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(resp)

	//habdling response error
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer response.Body.Close()

	// Get request

	// if err != nil{
	// 	log.Fatalln(err)
	// }

	// read response body
	var sb Response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error reading POST response body: ", err)

	}

	//unmarshal the response body into the response struct
	err = json.Unmarshal(body, &sb)
	if err != nil {
		fmt.Println("Error unmarshalling POST response")
		return err
	}

	fmt.Printf("POST Response: Reference: %s, USSD Code: %s\n", sb.Reference, sb.Ussd_Code)

	// making get request https://www.youtube.com/watch?v=EerdGm-ehJQ

	getUrl := fmt.Sprintf("https://demo.campay.net/api/transaction/%s/", sb.Reference)

	getReq, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		return fmt.Errorf("Error creating get request: %v", err)
	}

	getReq.Header.Set("Authorization", "Token "+apiKey)
	getReq.Header.Set("Content-Type", "application/json")

	time.Sleep(20 * time.Second)

	// sending the get request
	getResp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		return fmt.Errorf("error sending GEt request: %v", err)
	}
	defer getResp.Body.Close()

	//read the response body from get requestapiKey
	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		return fmt.Errorf(" error reading GET response body: %w", err)

	}

	fmt.Printf("Transaction Details: %v\n", string(getBody))

	// fmt.Println(string(body))

	// fmt.Printf("Reference:  %s\n Ussd_Code: %s\n ", sb.Reference, sb.Ussd_Code)
	return nil

}
