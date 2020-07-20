package test

import (
	"coffeeMaker/logic"
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

//csv file struct
type refillData struct {
	value       logic.CoffeeMachine
	quantity    int
	ingredients string
}

//read refill input data from csv file
func refillIngredientsTestCases() (CsvTests []refillData) {
	filename := "verifyRefillIngradient.csv"

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
		err := json.Unmarshal([]byte(column[2]), &postData)
		if err != nil {
			continue
		}
		val, _ := strconv.Atoi(column[1])
		data := refillData{
			ingredients: column[0],
			quantity:    val,
			value:       postData,
		}
		CsvTests = append(CsvTests, data)
	}
	return CsvTests
}

//@input will be what ingredients you want to add and in what quantity
//@output will print what was earlier and what is updated one
func TestRefillIngredients(t *testing.T) {
	TestInput := refillIngredientsTestCases()
	for _, value := range TestInput {
		logic.RefillIngredients(value.ingredients, value.quantity, value.value)
	}
}
