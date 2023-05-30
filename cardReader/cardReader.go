package cardReader

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetCardName() interface{} {
	resp, err := http.Get("https://mtgjson.com/api/v5/ONE.json")
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	data1 := generic["data"]
	cards1 := data1.(map[string]interface{})["cards"]
	firstCard := cards1.([]interface{})[0]
	cardName := firstCard.(map[string]interface{})["name"]

	return cardName

}
