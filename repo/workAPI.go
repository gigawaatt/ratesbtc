package repo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ViewBTC struct {
	Total int     `json:"total"`
	Rates []Rates `json:"rates"`
}

type Rates struct {
	Id   uint    `json:"id"`
	Date string  `json:"date"`
	Aud  float64 `json:"AUD"`
	Azn  float64 `json:"AZN"`
	Gbp  float64 `json:"GBP"`
	Amd  float64 `json:"AMD"`
	Byn  float64 `json:"BYN"`
	Bgn  float64 `json:"BGN"`
	Brl  float64 `json:"BRL"`
	Huf  float64 `json:"HUF"`
	Hkd  float64 `json:"HKD"`
	Dkk  float64 `json:"DKK"`
	Usd  float64 `json:"USD"`
	Eur  float64 `json:"EUR"`
	Inr  float64 `json:"INR"`
	Kzt  float64 `json:"KZT"`
	Cad  float64 `json:"CAD"`
	Kgs  float64 `json:"KGS"`
	Cny  float64 `json:"CNY"`
	Mdl  float64 `json:"MDL"`
	Nok  float64 `json:"NOK"`
	Pln  float64 `json:"PLN"`
	Ron  float64 `json:"RON"`
	Xdr  float64 `json:"XDR"`
	Sgd  float64 `json:"SGD"`
	Tjs  float64 `json:"TJS"`
	Try  float64 `json:"TRY"`
	Tmt  float64 `json:"TMT"`
	Uzs  float64 `json:"UZS"`
	Uah  float64 `json:"UAH"`
	Czk  float64 `json:"CZK"`
	Sek  float64 `json:"SEK"`
	Chf  float64 `json:"CHF"`
	Zar  float64 `json:"ZAR"`
	Krw  float64 `json:"KRW"`
	Jpy  float64 `json:"JPY"`
}

type BTCUSDT struct {
	Id           uint    `json:"id"`
	Time         int64   `json:"time"`
	AveragePrice float64 `json:"value"`
}

type AllBTCUSDT struct {
	Total   uint      `json:"total"`
	BTCUSDT []BTCUSDT `json:"history"`
}

type AllRates struct {
	Total uint    `json:"total"`
	Rates []Rates `json:"history"`
}

const (
	parseUrl = "http://www.cbr.ru/scripts/XML_daily.asp"
	parseBTC = "https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT"
)

type ValCurs struct {
	//XMLName xml.Name `xml:"ValCurs"`
	Date   string   `xml:"Date,attr" json:"date"`
	Name   string   `xml:"name,attr" json:"name"`
	Valute []Valute `xml:"Valute" json:"Valute"`
}

type Valute struct {
	//XMLName  xml.Name `xml:"Valute"`
	Id       string `xml:"ID,attr"  json:"Id"`
	NumCode  int    `xml:"NumCode" json:"NumCode"`
	CharCode string `xml:"CharCode" json:"CharCode"`
	Nominal  int    `xml:"Nominal" json:"Nominal"`
	Name     string `xml:"Name" json:"Name"`
	Value    string `xml:"Value" json:"Value"`
}

type JsonBTCUSDT struct {
	Code string `json:"code"`
	Data struct {
		Time             int64  `json:"time"`
		Symbol           string `json:"symbol"`
		Buy              string `json:"buy"`
		Sell             string `json:"sell"`
		ChangeRate       string `json:"changeRate"`
		ChangePrice      string `json:"changePrice"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Vol              string `json:"vol"`
		VolValue         string `json:"volValue"`
		Last             string `json:"last"`
		AveragePrice     string `json:"averagePrice"`
		TakerFeeRate     string `json:"takerFeeRate"`
		MakerFeeRate     string `json:"makerFeeRate"`
		TakerCoefficient string `json:"takerCoefficient"`
		MakerCoefficient string `json:"makerCoefficient"`
	} `json:"data"`
}

func GetvalCurs(w http.ResponseWriter, r *http.Request) {

	var rates Rates
	db := Init()
	stmt := "SELECT * FROM ratescurs WHERE id = (SELECT MAX(id) FROM ratescurs)"
	err := db.QueryRow(stmt).Scan(&rates.Id, &rates.Date, &rates.Aud, &rates.Azn,
		&rates.Gbp, &rates.Amd, &rates.Byn, &rates.Bgn, &rates.Brl, &rates.Huf,
		&rates.Hkd, &rates.Dkk, &rates.Usd, &rates.Eur, &rates.Inr, &rates.Kzt,
		&rates.Cad, &rates.Kgs, &rates.Cny, &rates.Mdl, &rates.Nok, &rates.Pln,
		&rates.Ron, &rates.Xdr, &rates.Sgd, &rates.Tjs, &rates.Try,
		&rates.Tmt, &rates.Uzs, &rates.Uah, &rates.Czk, &rates.Sek, &rates.Chf, &rates.Zar,
		&rates.Krw, &rates.Jpy)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}

	fmt.Print(rates)
	defer db.Close()
	jsonData, _ := json.Marshal(rates)
	w.Write(jsonData)

}

func GetAllValCurs(w http.ResponseWriter, r *http.Request) {

	var rates []Rates

	db := Init()
	stmt := "SELECT * FROM ratescurs"
	result, err := db.Query(stmt)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	defer db.Close()

	var total uint = 0

	for result.Next() {
		var rate Rates
		if err := result.Scan(&rate.Id, &rate.Date, &rate.Aud, &rate.Azn,
			&rate.Gbp, &rate.Amd, &rate.Byn, &rate.Bgn, &rate.Brl, &rate.Huf,
			&rate.Hkd, &rate.Dkk, &rate.Usd, &rate.Eur, &rate.Inr, &rate.Kzt,
			&rate.Cad, &rate.Kgs, &rate.Cny, &rate.Mdl, &rate.Nok, &rate.Pln,
			&rate.Ron, &rate.Xdr, &rate.Sgd, &rate.Tjs, &rate.Try,
			&rate.Tmt, &rate.Uzs, &rate.Uah, &rate.Czk, &rate.Sek, &rate.Chf, &rate.Zar,
			&rate.Krw, &rate.Jpy); err != nil {
			fmt.Fprintf(w, "Error data : %v", err)
		}
		rates = append(rates, rate)
		total++
	}

	var allinfo AllRates
	allinfo.Total = total
	allinfo.Rates = rates

	jsonResponse, jsonError := json.Marshal(allinfo)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
		jsonError, _ := json.Marshal(jsonError)
		w.Write(jsonError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}

}

func GetBTCUSDT(w http.ResponseWriter, r *http.Request) {

	jsonBTC := Rates{}
	var jsonBTCDB BTCUSDT

	db := Init()
	stmt := "SELECT * FROM btc WHERE id = (SELECT MAX(id) FROM btc)"
	err := db.QueryRow(stmt).Scan(&jsonBTCDB.Id, &jsonBTCDB.Time, &jsonBTCDB.AveragePrice)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	log.Print(jsonBTC)
	defer db.Close()

	jsonResponse, jsonError := json.Marshal(jsonBTCDB)
	if jsonError != nil {
		log.Println("Unable to encode JSON")
		jsonError, _ := json.Marshal(jsonError)
		w.Write(jsonError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetAllBTCUSDT(w http.ResponseWriter, r *http.Request) {

	var jsonBTCDBArr []BTCUSDT
	db := Init()
	stmt := "SELECT * FROM btc"
	result, err := db.Query(stmt)
	if err != nil {
		fmt.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	defer db.Close()

	var total uint = 0
	for result.Next() {
		var jsonBTCDB BTCUSDT
		if err := result.Scan(&jsonBTCDB.Id, &jsonBTCDB.Time, &jsonBTCDB.AveragePrice); err != nil {
			fmt.Fprintf(w, "Error data : %v", err)
		}
		jsonBTCDBArr = append(jsonBTCDBArr, jsonBTCDB)
		total++
	}

	var allinfo AllBTCUSDT
	allinfo.Total = total
	allinfo.BTCUSDT = jsonBTCDBArr

	jsonResponse, jsonError := json.Marshal(allinfo)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
		jsonError, _ := json.Marshal(jsonError)
		w.Write(jsonError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetValcursWithBTC(w http.ResponseWriter, r *http.Request) {

	var rates Rates
	db := Init()

	stmt := "SELECT * FROM ratescurs WHERE id = (SELECT MAX(id) FROM ratescurs)"
	err := db.QueryRow(stmt).Scan(&rates.Id, &rates.Date, &rates.Aud, &rates.Azn,
		&rates.Gbp, &rates.Amd, &rates.Byn, &rates.Bgn, &rates.Brl, &rates.Huf,
		&rates.Hkd, &rates.Dkk, &rates.Usd, &rates.Eur, &rates.Inr, &rates.Kzt,
		&rates.Cad, &rates.Kgs, &rates.Cny, &rates.Mdl, &rates.Nok, &rates.Pln,
		&rates.Ron, &rates.Xdr, &rates.Sgd, &rates.Tjs, &rates.Try,
		&rates.Tmt, &rates.Uzs, &rates.Uah, &rates.Czk, &rates.Sek, &rates.Chf, &rates.Zar,
		&rates.Krw, &rates.Jpy)
	if err != nil {
		fmt.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}

	var jsonBTCDB BTCUSDT
	stmt = "SELECT * FROM btc WHERE id = (SELECT MAX(id) FROM btc)"
	err = db.QueryRow(stmt).Scan(&jsonBTCDB.Id, &jsonBTCDB.Time, &jsonBTCDB.AveragePrice)
	if err != nil {
		fmt.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	defer db.Close()

	res := rates2BTC(rates, jsonBTCDB.AveragePrice)
	res.Date = time.Now().String()
	res.Id = jsonBTCDB.Id
	fmt.Print(res)
	jsonData, _ := json.Marshal(res)
	w.Write(jsonData)

}

func GetAllValcursWithBTC(w http.ResponseWriter, r *http.Request) {

	var rates []Rates

	db := Init()
	stmt := "SELECT * FROM ratescurs"
	result, err := db.Query(stmt)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}

	var jsonBTCDB BTCUSDT
	stmt = "SELECT * FROM btc WHERE id = (SELECT MAX(id) FROM btc)"
	err = db.QueryRow(stmt).Scan(&jsonBTCDB.Id, &jsonBTCDB.Time, &jsonBTCDB.AveragePrice)
	if err != nil {
		fmt.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	defer db.Close()

	var total uint = 0

	for result.Next() {
		var rate Rates
		if err := result.Scan(&rate.Id, &rate.Date, &rate.Aud, &rate.Azn,
			&rate.Gbp, &rate.Amd, &rate.Byn, &rate.Bgn, &rate.Brl, &rate.Huf,
			&rate.Hkd, &rate.Dkk, &rate.Usd, &rate.Eur, &rate.Inr, &rate.Kzt,
			&rate.Cad, &rate.Kgs, &rate.Cny, &rate.Mdl, &rate.Nok, &rate.Pln,
			&rate.Ron, &rate.Xdr, &rate.Sgd, &rate.Tjs, &rate.Try,
			&rate.Tmt, &rate.Uzs, &rate.Uah, &rate.Czk, &rate.Sek, &rate.Chf, &rate.Zar,
			&rate.Krw, &rate.Jpy); err != nil {
			fmt.Fprintf(w, "Error data : %v", err)
		}
		res := rates2BTC(rate, jsonBTCDB.AveragePrice)
		res.Id = rate.Id
		res.Date = rate.Date
		rates = append(rates, res)
		total++
	}

	var allinfo AllRates
	allinfo.Total = total
	allinfo.Rates = rates

	jsonResponse, jsonError := json.Marshal(allinfo)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
		jsonError, _ := json.Marshal(jsonError)
		w.Write(jsonError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
}

func GetValute(w http.ResponseWriter, r *http.Request) {

	code, _ := mux.Vars(r)["id"]
	var value float64
	var valueUSD float64

	db := Init()
	stmt := "SELECT " + code + ", usd FROM ratescurs WHERE id = (SELECT MAX(id) FROM ratescurs)"
	err := db.QueryRow(stmt).Scan(&value, &valueUSD)
	if err != nil {
		jsonError, _ := json.Marshal(err)
		fmt.Println(jsonError)
		w.Write(jsonError)
	}

	var jsonBTCDB BTCUSDT
	stmt = "SELECT * FROM btc WHERE id = (SELECT MAX(id) FROM btc)"
	err = db.QueryRow(stmt).Scan(&jsonBTCDB.Id, &jsonBTCDB.Time, &jsonBTCDB.AveragePrice)
	if err != nil {
		fmt.Println(err)
		jsonError, _ := json.Marshal(err)
		w.Write(jsonError)
	}
	defer db.Close()

	value = valute2BTC(jsonBTCDB.AveragePrice, value, valueUSD)

	fmt.Print(value)
	jsonData, _ := json.Marshal(value)
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
