package repo

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
)

func rates2BTC(rates Rates, priceBTC float64) (ratesBTC Rates) {

	ratesBTC.Amd = valute2BTC(priceBTC, rates.Amd, rates.Usd)
	ratesBTC.Aud = valute2BTC(priceBTC, rates.Aud, rates.Usd)
	ratesBTC.Azn = valute2BTC(priceBTC, rates.Azn, rates.Usd)
	ratesBTC.Bgn = valute2BTC(priceBTC, rates.Bgn, rates.Usd)
	ratesBTC.Brl = valute2BTC(priceBTC, rates.Brl, rates.Usd)
	ratesBTC.Byn = valute2BTC(priceBTC, rates.Byn, rates.Usd)
	ratesBTC.Cad = valute2BTC(priceBTC, rates.Cad, rates.Usd)
	ratesBTC.Chf = valute2BTC(priceBTC, rates.Chf, rates.Usd)
	ratesBTC.Cny = valute2BTC(priceBTC, rates.Cny, rates.Usd)
	ratesBTC.Czk = valute2BTC(priceBTC, rates.Czk, rates.Usd)
	ratesBTC.Dkk = valute2BTC(priceBTC, rates.Dkk, rates.Usd)
	ratesBTC.Eur = valute2BTC(priceBTC, rates.Eur, rates.Usd)
	ratesBTC.Gbp = valute2BTC(priceBTC, rates.Gbp, rates.Usd)
	ratesBTC.Hkd = valute2BTC(priceBTC, rates.Hkd, rates.Usd)
	ratesBTC.Huf = valute2BTC(priceBTC, rates.Huf, rates.Usd)
	ratesBTC.Inr = valute2BTC(priceBTC, rates.Inr, rates.Usd)
	ratesBTC.Jpy = valute2BTC(priceBTC, rates.Jpy, rates.Usd)
	ratesBTC.Kgs = valute2BTC(priceBTC, rates.Kgs, rates.Usd)
	ratesBTC.Krw = valute2BTC(priceBTC, rates.Krw, rates.Usd)
	ratesBTC.Kzt = valute2BTC(priceBTC, rates.Kzt, rates.Usd)
	ratesBTC.Mdl = valute2BTC(priceBTC, rates.Mdl, rates.Usd)
	ratesBTC.Nok = valute2BTC(priceBTC, rates.Nok, rates.Usd)
	ratesBTC.Pln = valute2BTC(priceBTC, rates.Pln, rates.Usd)
	ratesBTC.Ron = valute2BTC(priceBTC, rates.Ron, rates.Usd)
	ratesBTC.Sek = valute2BTC(priceBTC, rates.Sek, rates.Usd)
	ratesBTC.Sgd = valute2BTC(priceBTC, rates.Sgd, rates.Usd)
	ratesBTC.Tjs = valute2BTC(priceBTC, rates.Tjs, rates.Usd)
	ratesBTC.Tmt = valute2BTC(priceBTC, rates.Tmt, rates.Usd)
	ratesBTC.Try = valute2BTC(priceBTC, rates.Try, rates.Usd)
	ratesBTC.Uah = valute2BTC(priceBTC, rates.Uah, rates.Usd)
	ratesBTC.Usd = priceBTC
	ratesBTC.Uzs = valute2BTC(priceBTC, rates.Uzs, rates.Usd)
	ratesBTC.Xdr = valute2BTC(priceBTC, rates.Xdr, rates.Usd)
	ratesBTC.Xdr = valute2BTC(priceBTC, rates.Xdr, rates.Usd)
	ratesBTC.Zar = valute2BTC(priceBTC, rates.Zar, rates.Usd)

	return ratesBTC
}

func valute2BTC(priceBTC, valute, USD float64) float64 {

	result := (priceBTC * USD) / valute
	return result

}

func parseXml() (vc *ValCurs) {
	req, err := http.NewRequest("GET", parseUrl, nil)
	if err != nil {
		log.Println("ACHTUNG!", err)
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("ACHTUNG!", err)
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	vc = &ValCurs{}
	err = decoder.Decode(&vc)
	if err != nil {
		log.Println("ACHTUNG!", err)
	}

	return vc
}

func Dump(vc *ValCurs) {
	fmt.Println("ValCurs")
	fmt.Println("--------------------------")
	for _, i := range vc.Valute {
		fmt.Println("Name:\t", i.Name)
		fmt.Println("CharCode:\t", i.CharCode)
		fmt.Println("Nominal:\t", i.Nominal)
		fmt.Println("NumCode:\t", i.NumCode)
		fmt.Println("Value:\t", i.Value)
		fmt.Println("")
	}
}
