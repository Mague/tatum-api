package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func RequestPost(endPoint string, requestBody *bytes.Buffer, fn func([]byte, int)) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Siempree lee el archivo")
	}

	tatumApiUrl := os.Getenv("TATUM_API_URL")
	tatumApiKey := os.Getenv("TATUM_API_KEY")

	req, _ := http.NewRequest("POST", tatumApiUrl+endPoint, requestBody)

	req.Header.Add("content-type", "application/json")
	//req.Header.Add("x-testnet-type", "SOME_STRING_VALUE")
	req.Header.Add("x-api-key", tatumApiKey)

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	fn(body, res.StatusCode)
}
