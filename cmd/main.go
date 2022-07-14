package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rByte, err := getContentAsString("https://api.waxpeer.com/v1/get-items-list?api=a7ca663f5fcacd8420419caa9192507b3b15316a379374321c4e6831f377fdf6&skip=0&sort=profit&min_price=50000&game=csgo")
	if err != nil {
		log.Fatal(err)
	}

	var final itemsList
	err = json.Unmarshal(rByte, &final)
	if err != nil {
		log.Fatal(err)
	}

	if !final.Success {
		log.Fatal("unsuccessful")
	}

	ePrice := final.Items[0].Price / 1000
	eDeals := final.Items[0].BestDeals / 1000
	eSteamPrice := final.Items[0].SteamPrice / 1000

	for _, take := range final.Items {
		fmt.Printf("\nItemID: %s\nBrand: %s\nImage: %s\nType: %s\nPrice: %d\nName: %s\nFloat: %f\nBest deals: %d\nDiscount: %d\nSteam Price: %d\n_____",
			take.ItemID, take.Brand, take.Image, take.Type, ePrice, take.Name, take.Float, eDeals, take.Discount, eSteamPrice)
	}
}

func getContentAsString(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}

type itemsList struct {
	Success bool       `json:"success"`
	Items   []itemList `json:"items"`
}

type itemList struct {
	ItemID     string  `json:"item_id"`
	Brand      string  `json:"brand"`
	Image      string  `json:"image"`
	Type       string  `json:"type"`
	Price      int     `json:"price"`
	Name       string  `json:"name"`
	Float      float64 `json:"float"`
	BestDeals  int     `json:"best_deals"`
	Discount   int     `json:"discount"`
	SteamPrice int     `json:"steam_price"`
}
