package test

import (
	"coffeeMaker/logic"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

//read input data from csv file
func lowIngredientsBasedOnDrinkTestCases() (CsvTests []logic.CoffeeMachine) {
	filename := "verifyLowGredient_inBeverage.csv"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	header := true
	var postData logic.CoffeeMachine
	for _, column := range lines {
		if header {
			header = false
			continue
		}
		err := json.Unmarshal([]byte(column[1]), &postData)
		if err != nil {
			continue
		}
		CsvTests = append(CsvTests, postData)
	}
	//fmt.Println(CsvTests)
	return CsvTests
}

//this will give which all ingredients are missing from each beverages
func TestLowIngredientsBasedOnDrink(t *testing.T) {
	TestInput := lowIngredientsBasedOnDrinkTestCases()
	count := 1
	for _, value := range TestInput {
		fmt.Println("TestCases ======================", count)
		logic.LowIngredientsBasedOnDrink(value)
		fmt.Println()
		count++
	}
}
