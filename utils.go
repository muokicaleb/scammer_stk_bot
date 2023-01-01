package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func BearerTokenGenerator() string {
	godotenv.Load(".env")

	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	tokenVal := fmt.Sprintf("%s:%s", os.Getenv("SAFARICOM_CONSUMER_KEY"), os.Getenv("SAFARICOM_CONSUMER_SECRET"))
	tokenValb64 := base64.StdEncoding.EncodeToString([]byte(tokenVal))
	fmt.Println(os.Getenv("SAFARICOM_CONSUMER_KEY"))
	fmt.Println(tokenValb64)

	req.Header.Add("Authorization", "Basic "+tokenValb64)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var result TokenData
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return result.AccessToken
}

func GetTimeStamp() string {
	// Load the Africa/Nairobi location
	location, err := time.LoadLocation("Africa/Nairobi")
	if err != nil {
		fmt.Println(err)
		return "timezone error"
	}

	// Get the current time in the Africa/Nairobi location
	now := time.Now().In(location)

	// Format the time using the layout "YYYYMMDDHHmmss"
	formattedTime := now.Format("20060102150405")
	fmt.Println(formattedTime)

	return formattedTime
}

func GetReqPass(timeStamp string) string {
	val := fmt.Sprintf("%s%s%s", os.Getenv("BUSINESS_SHORT_CODE"), os.Getenv("PASS_KEY"), timeStamp)
	return base64.StdEncoding.EncodeToString([]byte(val))

}
