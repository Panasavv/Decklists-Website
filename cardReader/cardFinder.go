package cardReader

import (
	"encoding/json"
	"log"
	"net/http"
)

func PioneerCardFinder(s string) bool {
	resp, err := http.Get("https://mtgjson.com/api/v5/PioneerAtomic.json")
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	data1 := generic["data"]
	cards1 := data1.(map[string]interface{})[s]
	if cards1 != nil {
		cards1 = nil
		return true
	} else {
		return false
	}

}

func ModernCardFinder(s string) bool {
	resp, err := http.Get("https://mtgjson.com/api/v5/ModernAtomic.json")
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	data1 := generic["data"]
	cards1 := data1.(map[string]interface{})[s]
	if cards1 != nil {
		cards1 = nil
		return true
	} else {
		return false
	}

}
