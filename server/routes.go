package server

import (
	"github.com/gigawaatt/ratebtc/repo"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	//Clients
	r.Methods("GET").Path("/api/currencies").HandlerFunc(repo.GetvalCurs)
	r.Methods("GET").Path("/api/btcusdt").HandlerFunc(repo.GetBTCUSDT)
	r.Methods("POST").Path("/api/currencies").HandlerFunc(repo.GetAllValCurs)
	r.Methods("POST").Path("/api/btcusdt").HandlerFunc(repo.GetAllBTCUSDT)
	r.Methods("GET").Path("/api/latest").HandlerFunc(repo.GetValcursWithBTC)
	r.Methods("POST").Path("/api/latest").HandlerFunc(repo.GetAllValcursWithBTC)
	r.Methods("GET").Path("/api/latest/{id:[A-z-]+}").HandlerFunc(repo.GetValute)
	return r

}
