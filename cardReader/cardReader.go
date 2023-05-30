package cardReader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getCardName() {
	resp, err := http.Get("https://mtgjson.com/api/v5/ONE.json")
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	//Getting the date from meta
	//	imerominia := generic["meta"].(map[string]interface{})["date"]
	//	fmt.Println(imerominia)

	data1 := generic["data"]
	cards1 := data1.(map[string]interface{})["cards"]
	firstCard := cards1.([]interface{})[0]
	cardName := firstCard.(map[string]interface{})["name"]

	fmt.Println(cardName)

	//cardName := generic["data"].(map[string]interface{})["cards"].([]interface{})[0]
}
