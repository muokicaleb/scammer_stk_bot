package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ScammerStkPush(bearerToken string, targetNumber int, requestID string, pushAmount int) {

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
		"Amount":            pushAmount,
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
