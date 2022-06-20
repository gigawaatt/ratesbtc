package main

import (
	"github.com/gigawaatt/Golang/repo"
	"github.com/gigawaatt/Golang/server"
)

func main() {

	go repo.SetValCurs()
	go repo.SetBTCUSDT()
	server.Start()

}
