package main

import (
	"github.com/gigawaatt/ratebtc/repo"
	"github.com/gigawaatt/ratebtc/server"
)

func main() {

	go repo.SetValCurs()
	go repo.SetBTCUSDT()
	server.Start()

}
