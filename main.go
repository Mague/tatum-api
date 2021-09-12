package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Mague/tatum-api/middlewares"
	"github.com/Mague/tatum-api/payloads"
	"github.com/Mague/tatum-api/responses"
)

var router *gin.Engine

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	tatumApiUrl := os.Getenv("TATUM_API_URL")
	tatumApiKey := os.Getenv("TATUM_API_KEY")
	gin.SetMode(gin.DebugMode)

	router = gin.Default()
	router.Use(middlewares.RateLimit, gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Running",
		})
	})
	router.POST("/nft/deploy/", func(ctx *gin.Context) {
		var deploy payloads.Deploy
		ctx.BindJSON(&deploy)
		/*{
		"chain": "CELO",
		"name": "My ERC721",
		"symbol": "ERC_SYMBOL",
		"fromPrivateKey": "0x05e150c73f1920ec14caa1e0b6aa09940899678051a78542840c2668ce5080c2",
		"nonce": 0,
		"feeCurrency": "CELO"
		}*/
		postBody, _ := json.Marshal(deploy)

		requestBody := bytes.NewBuffer(postBody)
		req, _ := http.NewRequest("POST", tatumApiUrl+"/v3/nft/deploy/", requestBody)

		req.Header.Add("content-type", "application/json")
		//req.Header.Add("x-testnet-type", "SOME_STRING_VALUE")
		req.Header.Add("x-api-key", tatumApiKey)

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		if res.StatusCode == 200 {
			var txId responses.TxId

			if err := json.Unmarshal(body, &txId); err != nil {
				panic(err)
			}
			fmt.Println(txId.TxId)
			ctx.JSON(http.StatusOK, txId)
		} else {
			var errorTatum responses.ErrorTatum
			if err := json.Unmarshal(body, &errorTatum); err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusConflict, errorTatum)
		}
		//fmt.Println(res)
		fmt.Println(string(body))

		/*resp, err := http.Post(
			tatumApiUrl+"/nft/deploy/",
			"application/json",
		)*/

	})
	router.Run(":" + httpPort)
}
