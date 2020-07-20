# coffeeMaker

==================================================

This sample code helps get you started with a Go Code which Consist Make Beverages,find all ingredients which is low to make all kind of beverages and refill the ingredients.

For that you can refer func
makeBeverages() -> this will take input as json
lowIngredientsBasedOnDrink() ->this will take input as json
refillIngredients() -> this will take input as,ingrediesnt name you want to refill,quantity and input as json

What's Here
-----------

This sample includes:

* README.md - this file
* main.go - this file contains the sample Go code for all func
* logic folder - this contain all logic related code
* test folder - this contain the go test file and test input file





***************************TestFile*******************
There are three test file
1) coffeeMaker_test.go -> this take the input from file csv file(verify_coffeeMaker.csv) and validate function makeBeverages()

2) lowGredient_inBeverage_test.go -> this take the input from file csv file(verifyLowGredient_inBeverage.csv) and validate function lowIngredientsBasedOnDrink()

3)refillIngradient_test.go -> this take the input from file csv file(verifyRefillIngradient.csv) and validate function refillIngredients()

*******************How to execute Test file *******************
1) Download this package at you go root path
2) cd to directory coffeeMaker
2) go test -v test/coffeeMaker_test.go
3) go test -v test/lowGredient_inBeverage_test.go
4) go test -v test/refillIngradient_test.go

