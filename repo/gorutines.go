package repo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SetValCurs() {
	for {
		curs := parseXml()
		jsonData, _ := json.Marshal(curs)
		fmt.Println(string(jsonData))
		valcurs2DB(curs)
		time.Sleep(time.Minute * 10)
	}
}

func valcurs2DB(curs *ValCurs) {
	db := Init()

	var dateFromDB string
	stmt := "SELECT date FROM ratestest ORDER BY ID DESC LIMIT 1"
	err := db.QueryRow(stmt).Scan(&dateFromDB)
	if err != nil {
		fmt.Println(err)
	}

	if curs.Date != dateFromDB {
		_, err = db.Exec(`INSERT INTO "ratestest" ("date") VALUES ($1)`, curs.Date)
		if err != nil {
			fmt.Println(err)
		}
		// db.Close()
		//_, err := db.Exec("INSERT INTO IF NOT EXIST ratestest (date, AUD, AZN, GBP, AMD, BYN, BGN, BRL, HUF, HKD, DKK, USD, EUR, INR, KZT, CAD, KGS, CNY, MDL, NOK, PLN, RON, XDR, SGD, TJS, TRY, TMT, UZS, UAH, CZK, SEK, CHF, ZAR, KRW, JPY) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		//	curs.Date, curs.Valute[0].Value, curs.Valute[1].Value, curs.Valute[2].Value, curs.Valute[3].Value, curs.Valute[4].Value, curs.Valute[5].Value, curs.Valute[6].Value, curs.Valute[7].Value, curs.Valute[8].Value, curs.Valute[9].Value, curs.Valute[10].Value, curs.Valute[11].Value, curs.Valute[12].Value, curs.Valute[13].Value, curs.Valute[14].Value, curs.Valute[15].Value, curs.Valute[16].Value, curs.Valute[17].Value, curs.Valute[18].Value, curs.Valute[19].Value, curs.Valute[20].Value, curs.Valute[21].Value, curs.Valute[22].Value, curs.Valute[23].Value, curs.Valute[24].Value, curs.Valute[25].Value, curs.Valute[26].Value, curs.Valute[27].Value, curs.Valute[28].Value, curs.Valute[29].Value, curs.Valute[30].Value, curs.Valute[31].Value, curs.Valute[32].Value, curs.Valute[33].Value)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// db = Init()
		for index := range curs.Valute {
			//fmt.Printf("UPDATE ratestest SET ? = '?'", curs.Valute[index].CharCode, curs.Valute[index].Value)
			//	curs.Valute[index].CharCode, curs.Valute[index].Value, curs.Date)
			strExec := "UPDATE ratestest SET " + curs.Valute[index].CharCode + " = '" + strings.Replace(curs.Valute[index].Value, ",", ".", -1) + "' WHERE date = '" + curs.Date + "'"

			fmt.Print(index)
			//_, err := db.Exec("UPDATE ratestest SET %s = %s", curs.Valute[index].CharCode, curs.Valute[index].Value)
			_, err := db.Exec(strExec)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	defer db.Close()

}

func SetBTCUSDT() {

	for {
		jsonBTC := JsonBTCUSDT{}
		req, err := http.NewRequest("GET", parseBTC, nil)
		if err != nil {
			log.Println("ACHTUNG!", err)
		}
		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			log.Fatal("ACHTUNG!", err)
		}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&jsonBTC)
		if err != nil {
			panic(err)
		}

		btcusdt2DB(jsonBTC)

		time.Sleep(time.Second * 20)
	}
}

func btcusdt2DB(jsonBTC JsonBTCUSDT) {

	average, err := strconv.ParseFloat(jsonBTC.Data.Buy, 64)
	db := Init()
	// stmt := "SELECT * FROM mytable2 WHERE  dataaverageprice = $1"
	// result, err := db.Query(stmt, average)
	var valueFromDB float64
	stmt := "SELECT dataaverageprice FROM mytable2 ORDER BY dataaverageprice DESC LIMIT 1"
	err = db.QueryRow(stmt).Scan(&valueFromDB)
	if err != nil {
		fmt.Println(err)
	}

	if average != valueFromDB {

		_, err = db.Exec(`INSERT into "mytable2" ("datatime", "dataaverageprice") VALUES ($1,$2)`,
			jsonBTC.Data.Time, average)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer db.Close()
}
