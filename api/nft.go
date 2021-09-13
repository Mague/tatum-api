package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Mague/tatum-api/payloads"
	"github.com/Mague/tatum-api/responses"
	"github.com/Mague/tatum-api/utils"
)

type Nft struct {
	ctx    *gin.Context
	router *gin.Engine
}

func (this Nft) Load(engine *gin.Engine) {
	this.router = engine
	nft := this.router.Group("api/v1/nft")
	{
		nft.POST("/deploy", this.deploy)
		nft.POST("/mint", this.mint)
	}
}

func (this Nft) deploy(ctx *gin.Context) {
	var deploy payloads.Deploy
	ctx.BindJSON(&deploy)

	postBody, _ := json.Marshal(deploy)
	requestBody := bytes.NewBuffer(postBody)

	utils.RequestPost("/v3/nft/deploy/", requestBody, func(body []byte, statusCode int) {
		if statusCode == 200 {
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
	})
}

func (this Nft) mint(ctx *gin.Context) {
	var mint payloads.Mint
	ctx.BindJSON(&mint)
	postBody, _ := json.Marshal(mint)
	requestBody := bytes.NewBuffer(postBody)

	utils.RequestPost("/v3/nft/mint", requestBody, func(body []byte, statusCode int) {
		if statusCode == 200 {
			var txId responses.TxId

			if err := json.Unmarshal(body, &txId); err != nil {
				panic(err)
			}
			fmt.Println(txId.TxId)
			ctx.JSON(http.StatusOK, txId)
		} else {
			if statusCode == 400 {
				var errorTatum responses.Error400
				if err := json.Unmarshal(body, &errorTatum); err != nil {
					panic(err)
				}
				ctx.JSON(http.StatusConflict, errorTatum)
			} else {

				var errorTatum responses.ErrorTatum
				if err := json.Unmarshal(body, &errorTatum); err != nil {
					panic(err)
				}
				ctx.JSON(http.StatusConflict, errorTatum)
			}
		}
	})
}
