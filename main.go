package main

import (
	"coffeeMaker/logic"
)

//Make diffrenet kind of beverages
func makeBeverages(inputData logic.CoffeeMachine) {
	logic.MakeBeverages(inputData)
	return
}

//find which ingredients are low for all kind of beverages
func lowIngredientsBasedOnDrink(inputData logic.CoffeeMachine) {
	logic.LowIngredientsBasedOnDrink(inputData)
	return
}

//Refill ingredients
func refillIngredients(ingredients string, quantity int, inputData logic.CoffeeMachine) {
	logic.RefillIngredients(ingredients, quantity, inputData)
	return
}

//Code is Commentes but we can test from main.go also to pass as json input
// func main() {
// }
