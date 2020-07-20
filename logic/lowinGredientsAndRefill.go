package logic

import (
	"fmt"
	"strconv"
)

//find which gredient is low in all drink
func LowIngredientsBasedOnDrink(inputData CoffeeMachine) {

	hotWater := inputData.Machine.TotalItemsQuantity.HotWater
	hotMilk := inputData.Machine.TotalItemsQuantity.HotMilk
	gingerSyrup := inputData.Machine.TotalItemsQuantity.GingerSyrup
	sugarSyrup := inputData.Machine.TotalItemsQuantity.SugarSyrup
	teaLeaveSyrup := inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup
	elaichiSyrup := inputData.Machine.TotalItemsQuantity.ElaichiSyrup
	cofferrSyrup := inputData.Machine.TotalItemsQuantity.CoffeeSyrup

	lowIngredientsForginger(hotWater, hotMilk, gingerSyrup, sugarSyrup, teaLeaveSyrup, inputData)
	lowIngredientsForelaichi(hotWater, hotMilk, elaichiSyrup, sugarSyrup, teaLeaveSyrup, inputData)
	lowIngredientsForCoffee(hotWater, hotMilk, cofferrSyrup, sugarSyrup, inputData)
	lowIngredientsForHotWater(hotWater, inputData)
	lowIngredientsForHotMilk(hotMilk, inputData)
	return

}

//low gredient to make ginger
func lowIngredientsForginger(hotWater, hotMilk, gingerSyrup, sugarSyrup, teaLeaveSyrup int, inputData CoffeeMachine) {
	reqHotWater := inputData.Machine.Beverages.GingerTea.HotWater
	reqHotMilk := inputData.Machine.Beverages.GingerTea.HotMilk
	reqGingerSyrup := inputData.Machine.Beverages.GingerTea.GingerSyrup
	reqsugarSyrup := inputData.Machine.Beverages.GingerTea.SugarSyrup
	reqTeaLeaveSyrup := inputData.Machine.Beverages.GingerTea.TeaLeavesSyrup
	output := make([]string, 0)
	if hotWater < reqHotWater {
		output = append(output, "Hot Water")
	}

	if hotMilk < reqHotMilk {
		output = append(output, "Hot Milk")
	}

	if gingerSyrup < reqGingerSyrup {
		output = append(output, "Ginger Syrup ")
	}

	if sugarSyrup < reqsugarSyrup {
		output = append(output, "Sugar syrup")
	}

	if teaLeaveSyrup < reqTeaLeaveSyrup {
		output = append(output, "Tea leaves")
	}

	if len(output) > 0 {
		fmt.Println("For Ginger Tea below ingredients are low :")
		for _, val := range output {
			fmt.Println(val)
		}
		fmt.Println()
	}

}

//low gredient to make elachi tea
func lowIngredientsForelaichi(hotWater, hotMilk, elaichiSyrup, sugarSyrup, teaLeaveSyrup int, inputData CoffeeMachine) {
	reqHotWater := inputData.Machine.Beverages.ElaichiTea.HotWater
	reqHotMilk := inputData.Machine.Beverages.ElaichiTea.HotMilk
	reqsugarSyrup := inputData.Machine.Beverages.ElaichiTea.SugarSyrup
	reqTeaLeaveSyrup := inputData.Machine.Beverages.ElaichiTea.TeaLeavesSyrup
	reqElaichiSyrup := inputData.Machine.Beverages.ElaichiTea.ElaichiSyrup
	output := make([]string, 0)
	if hotWater < reqHotWater {
		output = append(output, "Hot Water")
	}
	if hotMilk < reqHotMilk {
		output = append(output, "Hot Milk ")
	}
	if sugarSyrup < reqsugarSyrup {
		output = append(output, "Sugar syrup ")
	}
	if teaLeaveSyrup < reqTeaLeaveSyrup {
		output = append(output, "Tea leaves ")
	}
	if elaichiSyrup < reqElaichiSyrup {
		output = append(output, "Elaichi syrup ")
	}

	if len(output) > 0 {
		fmt.Println("For Elaichi Tea below ingredients are low :")
		for _, val := range output {
			fmt.Println(val)
		}
		fmt.Println()
	}

}

//low gredient to make coffee
func lowIngredientsForCoffee(hotWater, hotMilk, cofferrSyrup, sugarSyrup int, inputData CoffeeMachine) {
	reqHotWater := inputData.Machine.Beverages.Coffee.HotWater
	reqHotMilk := inputData.Machine.Beverages.Coffee.HotMilk
	reqsugarSyrup := inputData.Machine.Beverages.Coffee.SugarSyrup
	reqCoffeeSyrup := inputData.Machine.Beverages.Coffee.CoffeeSyrup
	output := make([]string, 0)
	if hotWater < reqHotWater {
		output = append(output, "Hot Water ")
	}

	if hotMilk < reqHotMilk {
		output = append(output, "Hot Milk ")
	}

	if sugarSyrup < reqsugarSyrup {
		output = append(output, "Sugar syrup ")
	}

	if cofferrSyrup < reqCoffeeSyrup {
		output = append(output, "Coffee syrup ")
	}
	if len(output) > 0 {
		fmt.Println("For Coffee below ingredients are low :")
		for _, val := range output {
			fmt.Println(val)
		}
		fmt.Println()
	}
}

//low gredient to make hot water
func lowIngredientsForHotWater(hotWater int, inputData CoffeeMachine) {
	reqHotWater := inputData.Machine.Beverages.Water.HotWater
	output := make([]string, 0)
	if hotWater < reqHotWater {
		output = append(output, "Hot Water ")
	}
	if len(output) > 0 {
		fmt.Println("For Hot water below ingredients are low :")
		for _, val := range output {
			fmt.Println(val)
		}
		fmt.Println()
	}
}

//low gredient to make hot milk
func lowIngredientsForHotMilk(hotMilk int, inputData CoffeeMachine) {
	reqHotMilk := inputData.Machine.Beverages.Milk.HotMilk
	output := make([]string, 0)

	if hotMilk < reqHotMilk {
		output = append(output, "Hot Milk ")
	}
	if len(output) > 0 {
		fmt.Println("For Hot Milk below ingredients are low :")
		for _, val := range output {
			fmt.Println(val)
		}
		fmt.Println()
	}
}

//Refill the quantity of specfic ingredients
func RefillIngredients(ingredients string, quantity int, inputData CoffeeMachine) {
	fmt.Println("Refill ingredients " + ingredients + " of quantity " + strconv.Itoa(quantity))
	data := StructToMap(inputData.Machine.TotalItemsQuantity)
	str := fmt.Sprintf("%v", data[ingredients])
	old_quantity, _ := strconv.Atoi(str)
	//ingredientsquantity := data[ingredients]
	fmt.Println("Earlier total quantity was ", old_quantity)
	fmt.Println("Now total quantity is ", old_quantity+quantity)
	return
}
