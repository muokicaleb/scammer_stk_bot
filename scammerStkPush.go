package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type StkPushData struct {
	BusinessShortCode int    `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            int    `json:"Amount"`
	PartyA            int    `json:"PartyA"`
	PartyB            int    `json:"PartyB"`
	PhoneNumber       int    `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

func ScammerStkPush(bearerToken string, targetNumber int, requestID string) {
	

	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"
	timeStamp := GetTimeStamp()
	requestPass := GetReqPass(timeStamp)
	CallBackURL := fmt.Sprintf("%s%s", os.Getenv("CALLBACK_URL"), requestID)

	postBody, _ := json.Marshal(map[string]interface{}{
		"BusinessShortCode": os.Getenv("BUSINESS_SHORT_CODE"),
		"Password":          requestPass,
		"Timestamp":         timeStamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            1,
		"PartyA":            targetNumber,
		"PartyB":            os.Getenv("BUSINESS_SHORT_CODE"),
		"PhoneNumber":       targetNumber,
		"CallBackURL":       CallBackURL,
		"AccountReference":  "NotBob",
		"TransactionDesc":   "Payment of scamming"})
		
	payload := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
