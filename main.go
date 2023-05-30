package main

import (
	"Decklists-Website/cardReader"
	"Decklists-Website/decklistHandler"

	"fmt"
)

func main() {
	fmt.Println(cardReader.GetCardName())
	fmt.Println(cardReader.PioneerCardFinder("Legion Loyalist"))
	fmt.Println(cardReader.PioneerCardFinder("Tarmogoyf"))
	fmt.Println(cardReader.ModernCardFinder("Tarmogoyf"))
	fmt.Println(cardReader.ModernCardFinder("Brainstorm"))
	fmt.Println(cardReader.ModernCardFinder("Preordain"))
	decklistHandler.DecklistIsOK("decklistHandler/Deck - Rakdos Midrange.dek")
}
