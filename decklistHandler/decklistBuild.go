package decklistHandler

import (
	"Decklists-Website/cardReader"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mainboard struct {
	Quantity      []int
	Names         []string
	AllLegal      bool
	TotalQuantity int
}

type Sideboard struct {
	Quantity      []int
	Names         []string
	AllLegal      bool
	TotalQuantity int
}

type Decklist struct {
	mainboard Mainboard
	sideboard Sideboard
}

func DecklistBuild(s string) bool {
	file, err := os.Open(s)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	var decklist Decklist
	i, j := 0, 0
	decklist.mainboard.AllLegal = true
	decklist.sideboard.AllLegal = true

	for key, value := range fileLines {
		//fmt.Printf("line %v : %s \n", key, value)
		fmt.Println(key)
		if strings.Contains(value, "Name=\"") == true {

			if strings.Contains(value, "Sideboard=\"false") == true {

				s1 := strings.Split(value, "Quantity=\"")
				s2 := strings.Split(s1[1], "\" Sideb")
				s3 := strings.Split(s2[1], "Name=\"")
				s4 := strings.Split(s3[1], "\"")
				quantMain, err := strconv.Atoi(s2[0])
				if err != nil {
					// ... handle error
					panic(err)
				}

				decklist.mainboard.Quantity = append(decklist.mainboard.Quantity, quantMain)
				decklist.mainboard.Names = append(decklist.mainboard.Names, s4[0])

				tester := cardReader.PioneerCardFinder(s4[0])
				fmt.Println(tester)
				decklist.mainboard.AllLegal = tester && decklist.mainboard.AllLegal
				fmt.Println(decklist.mainboard.AllLegal)

				decklist.mainboard.TotalQuantity += decklist.mainboard.Quantity[i]
				i++

			} else {
				s1 := strings.Split(value, "Quantity=\"")
				s2 := strings.Split(s1[1], "\" Sideb")
				s3 := strings.Split(s2[1], "Name=\"")
				s4 := strings.Split(s3[1], "\"")

				quantSide, err := strconv.Atoi(s2[0])
				if err != nil {
					// ... handle error
					panic(err)
				}

				decklist.sideboard.Quantity = append(decklist.sideboard.Quantity, quantSide)
				decklist.sideboard.Names = append(decklist.sideboard.Names, s4[0])

				tester2 := cardReader.PioneerCardFinder(s4[0])
				fmt.Println(tester2)
				decklist.sideboard.AllLegal = tester2 && decklist.sideboard.AllLegal
				fmt.Println(decklist.sideboard.AllLegal)

				decklist.sideboard.AllLegal = cardReader.PioneerCardFinder(s4[0]) && decklist.sideboard.AllLegal
				decklist.sideboard.TotalQuantity += decklist.sideboard.Quantity[j]
				j++
			}
		}

	}
	fmt.Println(decklist.mainboard.AllLegal)
	fmt.Println(decklist.mainboard.TotalQuantity)
	fmt.Println(decklist.sideboard.TotalQuantity)
	if decklist.mainboard.AllLegal == true && decklist.mainboard.TotalQuantity >= 60 && decklist.sideboard.TotalQuantity <= 15 {
		return true
	} else {
		return false
	}

}
