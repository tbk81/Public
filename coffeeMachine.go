package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Ingredients struct {
	Water  int     `json:"water"`
	Milk   int     `json:"milk"`
	Coffee int     `json:"coffee"`
	Cost   float64 `json:"cost"`
}

//type Drinks struct {
//	Espresso   Ingredients `json:"espresso"`
//	Latte      Ingredients `json:"latte"`
//	Cappuccino Ingredients `json:"cappuccino"`
//}

type Machine struct {
	Menu map[string]Ingredients `json:"menu"`
}

func main() {

	byteValue, err := os.ReadFile("menu-rev.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var coffee Machine
	err = json.Unmarshal(byteValue, &coffee)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for name, drink := range coffee.Menu {
		fmt.Printf("%v\t%v%T\t%v\t%v\t%v\n", name, drink.Water, drink.Water, drink.Milk, drink.Coffee, drink.Cost)
	}

	//fmt.Println("Espresso:", coffee.Menu.Espresso)
	//fmt.Println("Latte:", coffee.Menu.Latte)
	//fmt.Println("Cappuccino:", coffee.Menu.Cappuccino)
}
