package tests

import (
	"coffeeMaker/logic"
	"encoding/csv"
	"encoding/json"
	"fmt"

	"os"
	"testing"
)

func getCoffeeMakerTestCases() (CsvTests []logic.CoffeeMachine) {
	filename := "verify_coffeeMaker.csv"

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

func TestServerPossibleDrink(t *testing.T) {
	TestInput := getCoffeeMakerTestCases()
	count := 1
	for _, value := range TestInput {
		fmt.Println("TestCases ======================", count)
		logic.MakeBeverages(value)
		fmt.Println()
		count++
	}
}
