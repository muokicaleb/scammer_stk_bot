package pkg

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	return formattedTime
}

func GetReqPass(timeStamp string) string {
	val := fmt.Sprintf("%s%s%s", os.Getenv("BUSINESS_SHORT_CODE"), os.Getenv("PASS_KEY"), timeStamp)
	return base64.StdEncoding.EncodeToString([]byte(val))

}

func JsonStringToMap(jsonString string) (map[string]interface{}, error) {
	var data map[string]interface{}

	// Use json.Unmarshal to parse the JSON string
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err)
	}

	return data, err

}

func ConnectToPSQL() *sql.DB {

	// Set the values for the connection string
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dburl := os.Getenv("POSTGRES_URL")
	sslmode := "disable"

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=%s", user, password, dburl, dbname, sslmode)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// handle the error
	}

	// Create the table
	tableStmt := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			requestID TEXT,
			amount INT,
			status TEXT
		); 
	`, os.Getenv("POSTGRES_DB"))
	_, err = db.Exec(tableStmt)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("Table successfully created.")

	return db

}
