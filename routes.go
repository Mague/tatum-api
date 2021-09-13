package main

import (
	"github.com/Mague/tatum-api/api"
)

func initializeRoutes() {
	api.Nft{}.Load(router)
}
