package cardReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Cards struct {
	No   int
	Name string
}

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
	//data1 := [0].(map[string]interface{})["name"]
	sunolo := len(cards1.([]interface{}))

	var cardNames []string

	for i := 0; i < sunolo; i++ {
		card := cards1.([]interface{})[i]
		name := card.(map[string]interface{})["name"]
		name2 := fmt.Sprint(name)
		cardNames = append(cardNames, name2)

	}

	jsonString, _ := json.Marshal(cardNames)
	err = ioutil.WriteFile("userfile.json", jsonString, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return cardName

}
