package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ScammerStkPush(bearerToken string, targetNumber int, requestID string, pushAmount int) map[string]interface{} {

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

	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	respBody, err := JsonStringToMap(string(body))

	if err != nil {
		fmt.Println(err)

	}
	return respBody

}

func AddResponseToStkDB(stkResponse map[string]interface{}, requestID string, pushAmount int) {
	// check if value is valid
	_, ok := stkResponse["errorCode"]

	if ok {

		fmt.Println("error")
		return

	} else {

		// insert to db
		db := ConnectToPSQL()
		stmt := fmt.Sprintf(`INSERT INTO %s (requestID, amount, status) VALUES ($1, $2, $3)`, os.Getenv("POSTGRES_DB"))

		// Use the Query or Exec function to insert the values
		_, err := db.Query(stmt, requestID, pushAmount, "pending")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		fmt.Println("Inserted values into the table")
	}

}

func UpdateStkPushDB(callbackBody map[string]interface{}, param string) {
	// check if value is valid
	db := ConnectToPSQL()

	// Create the UPDATE statement
	stmt := fmt.Sprintf(`UPDATE %s SET status = $1 WHERE requestID = $2`, os.Getenv("POSTGRES_DB"))

	// Execute the statement
	_, err := db.Exec(stmt, "completed", param)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	fmt.Println("Entry successfully updated.")
}

func GetTransactionStatus(param string) string {
	db := ConnectToPSQL()
	// Query the value from the database
	stmt := fmt.Sprintf(`SELECT status FROM %s WHERE requestID = $1`, os.Getenv("POSTGRES_DB"))

	var value string

	err := db.QueryRow(stmt, param).Scan(&value)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	defer db.Close()
	return value
}
