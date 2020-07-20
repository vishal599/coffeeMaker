package logic

import "fmt"

const (
	ConstGingerTea  = "giner_tea"
	ConstElaichiTea = "elaichi_tea"
	ConstCoffee     = "coffee"
	ConstWater      = "water"
	ConstMilk       = "milk"
)

type CoffeeMachine struct {
	Machine struct {
		Outlets struct {
			CountN int `json:"count_n"`
		} `json:"outlets"`
		TotalItemsQuantity struct {
			HotWater       int `json:"hot_water"`
			HotMilk        int `json:"hot_milk"`
			GingerSyrup    int `json:"ginger_syrup"`
			SugarSyrup     int `json:"sugar_syrup"`
			TeaLeavesSyrup int `json:"tea_leaves_syrup"`
			ElaichiSyrup   int `json:"elaichi_syrup"`
			CoffeeSyrup    int `json:"coffee_syrup"`
		} `json:"total_items_quantity"`
		Beverages struct {
			GingerTea struct {
				HotWater       int `json:"hot_water"`
				HotMilk        int `json:"hot_milk"`
				GingerSyrup    int `json:"ginger_syrup"`
				SugarSyrup     int `json:"sugar_syrup"`
				TeaLeavesSyrup int `json:"tea_leaves_syrup"`
			} `json:"ginger_tea"`
			ElaichiTea struct {
				HotWater       int `json:"hot_water"`
				ElaichiSyrup   int `json:"elaichi_syrup"`
				HotMilk        int `json:"hot_milk"`
				SugarSyrup     int `json:"sugar_syrup"`
				TeaLeavesSyrup int `json:"tea_leaves_syrup"`
			} `json:"elaichi_tea"`
			Coffee struct {
				HotWater    int `json:"hot_water"`
				HotMilk     int `json:"hot_milk"`
				SugarSyrup  int `json:"sugar_syrup"`
				CoffeeSyrup int `json:"coffee_syrup"`
			} `json:"coffee"`
			Water struct {
				HotWater int `json:"hot_water"`
			} `json:"water"`
			Milk struct {
				HotMilk int `json:"hot_milk"`
			} `json:"milk"`
		} `json:"beverages"`
	} `json:"machine"`
}

//Print data for different beverages
var OutputFormat = map[string]string{
	"HWNA": "Hot Water is not Available",
	"HMNA": "Hot Milk is not Available",
	"GSNA": "Ginger Syrup is not Available",
	"SSNA": "Sugar Syrup is not Available",
	"TLNA": "Tea leaves Syrup is not Available",
	"ESNA": "Elaichi Syrup is not Available",
	"CSNA": "Coffee Syrup is not Available",
	"HMNS": "Hot Milk is not Sufficient",
	"HWNS": "Hot Water is not Sufficient",
	"GSNS": "Ginger Syrup is not Sufficient",
	"SSNS": "Sugar Syrup is not Sufficient",
	"TLNS": "Tea leaves Syrup is not Sufficient",
	"ESNS": "Elaichi Syrup is not Sufficient",
	"CSNS": "Coffee Syrup is not Sufficient",
}

//is Beverage available
var BeverageAvailability = map[string]bool{
	ConstGingerTea:  true,
	ConstElaichiTea: true,
	ConstCoffee:     true,
	ConstWater:      true,
	ConstMilk:       true,
}

//random drink order from N outlet
var randomBeverageMapping = map[string]string{
	"0": ConstGingerTea,
	"1": ConstElaichiTea,
	"2": ConstCoffee,
	"3": ConstWater,
	"4": ConstMilk,
}

var charset = "01234"

//based on random number generator,we are trying to prepare that drink and Continue until machine can serve any kind of beverage
//if that drink is possible to make then decreases value from json
func MakeBeverages(inputData CoffeeMachine) {
	//if outlet count is zero in that case can't serve any beverages
	if inputData.Machine.Outlets.CountN == 0 {
		fmt.Println("Outlet count is zero so can't serve any beverages")
		return
	}
	//assign default value that outlet can serve all drink
	isBeverageAvailable := BeverageAvailability
	for {

		//generate random number
		randomNum := string(charset[SeededRand.Intn(len(charset))])
		//beverage order by customer
		//if isBeverageAvailable[beverageOrderBycust] = true then check it can be still make or not
		//if that is false means we can't prepare that beverage then print deficient ingredient
		//if all availbale beverage are false then break
		if CanWeNotMakeAnyBeverage(isBeverageAvailable) {
			resetBeverages(isBeverageAvailable)
			break
		}
		if !isBeverageAvailable[randomBeverageMapping[randomNum]] {
			continue
		}

		//Actual logic should be to acquire distrubuted lock in mutes here, but for that I have to either use mutex/dynalo db,
		//but that code I can't know now because I need credential for that,that's why writting without that
		if randomBeverageMapping[randomNum] == ConstGingerTea {
			inputData, isBeverageAvailable = CanBePreparedGingerTea(inputData, isBeverageAvailable)
		}
		if randomBeverageMapping[randomNum] == ConstElaichiTea {
			inputData, isBeverageAvailable = CanBePreparedElaichiTea(inputData, isBeverageAvailable)
		}
		if randomBeverageMapping[randomNum] == ConstCoffee {
			inputData, isBeverageAvailable = CanBePreparedCoffee(inputData, isBeverageAvailable)
		}
		if randomBeverageMapping[randomNum] == ConstWater {
			inputData, isBeverageAvailable = CanBePreparedWater(inputData, isBeverageAvailable)
		}
		if randomBeverageMapping[randomNum] == ConstMilk {
			inputData, isBeverageAvailable = CanBePreparedMilk(inputData, isBeverageAvailable)
		}

	}

	return
}

//reset beverages to true
func resetBeverages(isBeverageAvailable map[string]bool) {
	isBeverageAvailable[ConstGingerTea] = true
	isBeverageAvailable[ConstElaichiTea] = true
	isBeverageAvailable[ConstCoffee] = true
	isBeverageAvailable[ConstWater] = true
	isBeverageAvailable[ConstMilk] = true

}

//check can we make any gredients or not from available quantity
func CanWeNotMakeAnyBeverage(isBeverageAvailable map[string]bool) bool {
	if !isBeverageAvailable[ConstGingerTea] && !isBeverageAvailable[ConstElaichiTea] && !isBeverageAvailable[ConstCoffee] && !isBeverageAvailable[ConstWater] && !isBeverageAvailable[ConstMilk] {
		return true
	}
	return false
}

//if quantity is 0 means not availble, if >0 but less than needed then not sufficient
//if all quantity is availble then that items can be prepared and decrease ingrdients value from input
func CanBePreparedGingerTea(inputData CoffeeMachine, isBeverageAvailable map[string]bool) (CoffeeMachine, map[string]bool) {
	output := make([]string, 0)
	//we are doing like this becuase reading data from json,if read data from db we can use transaction
	//if any quantity is less do o.Rollback(), else o.Commit(), in that case we don't need to do so much if else
	if isSufficientHotWaterPresentForGinger(inputData) && isSufficientHotMilkPresentForGinger(inputData) && isSufficientGingerSyrupPresentForGinger(inputData) && isSufficientSugarSyrupPresentForGinger(inputData) && isSufficientLeavesSyrupPresentForGinger(inputData) {
		inputData.Machine.TotalItemsQuantity.HotWater = inputData.Machine.TotalItemsQuantity.HotWater - inputData.Machine.Beverages.GingerTea.HotWater
		inputData.Machine.TotalItemsQuantity.HotMilk = inputData.Machine.TotalItemsQuantity.HotMilk - inputData.Machine.Beverages.GingerTea.HotMilk
		inputData.Machine.TotalItemsQuantity.GingerSyrup = inputData.Machine.TotalItemsQuantity.GingerSyrup - inputData.Machine.Beverages.GingerTea.GingerSyrup
		inputData.Machine.TotalItemsQuantity.SugarSyrup = inputData.Machine.TotalItemsQuantity.SugarSyrup - inputData.Machine.Beverages.GingerTea.SugarSyrup
		inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup = inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup - inputData.Machine.Beverages.GingerTea.TeaLeavesSyrup
		fmt.Println("Ginger Tea is Prepared")
		return inputData, isBeverageAvailable
	}
	//mark flag to can't make this drink
	isBeverageAvailable[ConstGingerTea] = false
	if inputData.Machine.TotalItemsQuantity.HotWater == 0 {
		output = append(output, OutputFormat["HWNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk == 0 {
		output = append(output, OutputFormat["HMNA"])
	}

	if inputData.Machine.TotalItemsQuantity.GingerSyrup == 0 {
		output = append(output, OutputFormat["GSNA"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup == 0 {
		output = append(output, OutputFormat["SSNA"])
	}
	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup == 0 {
		output = append(output, OutputFormat["TLNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk < inputData.Machine.Beverages.GingerTea.HotMilk && inputData.Machine.TotalItemsQuantity.HotMilk > 0 {
		output = append(output, OutputFormat["HMNS"])
	}

	if inputData.Machine.TotalItemsQuantity.HotWater < inputData.Machine.Beverages.GingerTea.HotWater && inputData.Machine.TotalItemsQuantity.HotWater > 0 {
		output = append(output, OutputFormat["HWNS"])
	}

	if inputData.Machine.TotalItemsQuantity.GingerSyrup < inputData.Machine.Beverages.GingerTea.GingerSyrup && inputData.Machine.TotalItemsQuantity.GingerSyrup > 0 {
		output = append(output, OutputFormat["GSNS"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup < inputData.Machine.Beverages.GingerTea.SugarSyrup && inputData.Machine.TotalItemsQuantity.SugarSyrup > 0 {
		output = append(output, OutputFormat["SSNS"])
	}

	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup < inputData.Machine.Beverages.GingerTea.TeaLeavesSyrup && inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup > 0 {
		output = append(output, OutputFormat["TLNS"])
	}

	fmt.Print("Ginger Tea is Not Prepared because ")
	for _, val := range output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println("")
	return inputData, isBeverageAvailable
}

//check can we prepare elaichi from available gredient or not
func CanBePreparedElaichiTea(inputData CoffeeMachine, isBeverageAvailable map[string]bool) (CoffeeMachine, map[string]bool) {
	output := make([]string, 0)
	//we are doing like this becuase reading data from json,if read data from db we can use transaction
	//if any quantity is less do o.Rollback(), else o.Commit(), in that case we don't need to do so much if else
	if isSufficientHotWaterPresentForElaichi(inputData) && isSufficientHotMilkPresentForElaichi(inputData) && isSufficientElaichiSyrupPresentForElaichi(inputData) && isSufficientSugarSyrupPresentForElaichi(inputData) && isSufficientLeavesSyrupPresentForElaichi(inputData) {
		inputData.Machine.TotalItemsQuantity.HotWater = inputData.Machine.TotalItemsQuantity.HotWater - inputData.Machine.Beverages.ElaichiTea.HotWater
		inputData.Machine.TotalItemsQuantity.HotMilk = inputData.Machine.TotalItemsQuantity.HotMilk - inputData.Machine.Beverages.ElaichiTea.HotMilk
		inputData.Machine.TotalItemsQuantity.ElaichiSyrup = inputData.Machine.TotalItemsQuantity.ElaichiSyrup - inputData.Machine.Beverages.ElaichiTea.ElaichiSyrup
		inputData.Machine.TotalItemsQuantity.SugarSyrup = inputData.Machine.TotalItemsQuantity.SugarSyrup - inputData.Machine.Beverages.ElaichiTea.SugarSyrup
		inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup = inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup - inputData.Machine.Beverages.ElaichiTea.TeaLeavesSyrup
		fmt.Println("Elaichi Tea is Prepared")
		return inputData, isBeverageAvailable
	}
	isBeverageAvailable[ConstElaichiTea] = false
	if inputData.Machine.TotalItemsQuantity.HotWater == 0 {
		output = append(output, OutputFormat["HWNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk == 0 {
		output = append(output, OutputFormat["HMNA"])
	}

	if inputData.Machine.TotalItemsQuantity.ElaichiSyrup == 0 {
		output = append(output, OutputFormat["ESNA"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup == 0 {
		output = append(output, OutputFormat["SSNA"])
	}
	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup == 0 {
		output = append(output, OutputFormat["TLNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk < inputData.Machine.Beverages.ElaichiTea.HotMilk && inputData.Machine.TotalItemsQuantity.HotMilk > 0 {
		output = append(output, OutputFormat["HMNS"])
	}

	if inputData.Machine.TotalItemsQuantity.HotWater < inputData.Machine.Beverages.ElaichiTea.HotWater && inputData.Machine.TotalItemsQuantity.HotWater > 0 {
		output = append(output, OutputFormat["HWNS"])
	}

	if inputData.Machine.TotalItemsQuantity.ElaichiSyrup < inputData.Machine.Beverages.ElaichiTea.ElaichiSyrup && inputData.Machine.TotalItemsQuantity.ElaichiSyrup > 0 {
		output = append(output, OutputFormat["ESNS"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup < inputData.Machine.Beverages.ElaichiTea.SugarSyrup && inputData.Machine.TotalItemsQuantity.SugarSyrup > 0 {
		output = append(output, OutputFormat["SSNS"])
	}

	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup < inputData.Machine.Beverages.ElaichiTea.TeaLeavesSyrup && inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup > 0 {
		output = append(output, OutputFormat["TLNS"])
	}

	fmt.Print("Elaichi Tea is Not Prepared because ")
	for _, val := range output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println("")
	return inputData, isBeverageAvailable
}

//check can we prepare coffee from available gredient or not
func CanBePreparedCoffee(inputData CoffeeMachine, isBeverageAvailable map[string]bool) (CoffeeMachine, map[string]bool) {
	output := make([]string, 0)
	//we are doing like this becuase reading data from json,if read data from db we can use transaction
	//if any quantity is less do o.Rollback(), else o.Commit(), in that case we don't need to do so much if else
	if isSufficientHotMilkPresentForCofee(inputData) && isSufficientCofeeSyrupPresentForCofee(inputData) && isSufficientSugarSyrupPresentForCofee(inputData) && isSufficientHotWaterPresentForCofee(inputData) {
		inputData.Machine.TotalItemsQuantity.HotWater = inputData.Machine.TotalItemsQuantity.HotWater - inputData.Machine.Beverages.Coffee.HotWater
		inputData.Machine.TotalItemsQuantity.HotMilk = inputData.Machine.TotalItemsQuantity.HotMilk - inputData.Machine.Beverages.Coffee.HotMilk
		inputData.Machine.TotalItemsQuantity.CoffeeSyrup = inputData.Machine.TotalItemsQuantity.CoffeeSyrup - inputData.Machine.Beverages.Coffee.CoffeeSyrup
		inputData.Machine.TotalItemsQuantity.SugarSyrup = inputData.Machine.TotalItemsQuantity.SugarSyrup - inputData.Machine.Beverages.ElaichiTea.SugarSyrup
		fmt.Println("Coffee is Prepared")
		return inputData, isBeverageAvailable
	}
	isBeverageAvailable[ConstCoffee] = false
	if inputData.Machine.TotalItemsQuantity.HotWater == 0 {
		output = append(output, OutputFormat["HWNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk == 0 {
		output = append(output, OutputFormat["HMNA"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup == 0 {
		output = append(output, OutputFormat["SSNA"])
	}
	if inputData.Machine.TotalItemsQuantity.CoffeeSyrup == 0 {
		output = append(output, OutputFormat["CSNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotMilk < inputData.Machine.Beverages.Coffee.HotMilk && inputData.Machine.TotalItemsQuantity.HotMilk > 0 {
		output = append(output, OutputFormat["HMNS"])
	}

	if inputData.Machine.TotalItemsQuantity.HotWater < inputData.Machine.Beverages.Coffee.HotWater && inputData.Machine.TotalItemsQuantity.HotWater > 0 {
		output = append(output, OutputFormat["HWNS"])
	}

	if inputData.Machine.TotalItemsQuantity.CoffeeSyrup < inputData.Machine.Beverages.Coffee.CoffeeSyrup && inputData.Machine.TotalItemsQuantity.CoffeeSyrup > 0 {
		output = append(output, OutputFormat["ESNS"])
	}

	if inputData.Machine.TotalItemsQuantity.SugarSyrup < inputData.Machine.Beverages.Coffee.SugarSyrup && inputData.Machine.TotalItemsQuantity.SugarSyrup > 0 {
		output = append(output, OutputFormat["SSNS"])
	}

	fmt.Print("Coffee is Not Prepared because ")
	for _, val := range output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println("")
	return inputData, isBeverageAvailable

}

//check can we prepare Hotwater from available gredient or not
func CanBePreparedWater(inputData CoffeeMachine, isBeverageAvailable map[string]bool) (CoffeeMachine, map[string]bool) {
	output := make([]string, 0)
	//we are doing like this becuase reading data from json,if read data from db we can use transaction
	//if any quantity is less do o.Rollback(), else o.Commit(), in that case we don't need to do so much if else
	if isSufficientWaterPresentForHotWater(inputData) {
		inputData.Machine.TotalItemsQuantity.HotWater = inputData.Machine.TotalItemsQuantity.HotWater - inputData.Machine.Beverages.Water.HotWater

		fmt.Println("Hot water is Prepared")
		return inputData, isBeverageAvailable
	}
	isBeverageAvailable[ConstWater] = false
	if inputData.Machine.TotalItemsQuantity.HotWater == 0 {
		output = append(output, OutputFormat["HWNA"])
	}

	if inputData.Machine.TotalItemsQuantity.HotWater < inputData.Machine.Beverages.Water.HotWater && inputData.Machine.TotalItemsQuantity.HotWater > 0 {
		output = append(output, OutputFormat["HWNS"])
	}

	fmt.Print("Hot water is Not Prepared because ")
	for _, val := range output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println("")
	return inputData, isBeverageAvailable
}

//check can we prepare Milk from available gredient or not
func CanBePreparedMilk(inputData CoffeeMachine, isBeverageAvailable map[string]bool) (CoffeeMachine, map[string]bool) {
	output := make([]string, 0)
	//we are doing like this becuase reading data from json,if read data from db we can use transaction
	//if any quantity is less do o.Rollback(), else o.Commit(), in that case we don't need to do so much if else
	if isSufficientMilkPresentForHotMilk(inputData) {
		inputData.Machine.TotalItemsQuantity.HotMilk = inputData.Machine.TotalItemsQuantity.HotMilk - inputData.Machine.Beverages.Milk.HotMilk

		fmt.Println("Hot Milk is Prepared")
		return inputData, isBeverageAvailable
	}
	isBeverageAvailable[ConstMilk] = false
	if inputData.Machine.TotalItemsQuantity.HotMilk == 0 {
		output = append(output, OutputFormat["HMNA"])

	}
	if inputData.Machine.TotalItemsQuantity.HotMilk < inputData.Machine.Beverages.Milk.HotMilk && inputData.Machine.TotalItemsQuantity.HotMilk > 0 {
		output = append(output, OutputFormat["HMNS"])
	}

	fmt.Print("Hot Milk is Not Prepared because ")
	for _, val := range output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println("")
	return inputData, isBeverageAvailable
}

//check Suffcinet Hot water availble to make ginger tae or not
func isSufficientHotWaterPresentForGinger(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotWater >= inputData.Machine.Beverages.GingerTea.HotWater {
		isSuff = true
	}
	return
}

//check Suffcinet Hot MIlk availble to make ginger tae or not
func isSufficientHotMilkPresentForGinger(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotMilk >= inputData.Machine.Beverages.GingerTea.HotMilk {
		isSuff = true
	}
	return
}

//check Suffcinet Ginger syrup availble to make ginger tea or not
func isSufficientGingerSyrupPresentForGinger(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.GingerSyrup >= inputData.Machine.Beverages.GingerTea.GingerSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet Sugar syrup availble to make ginger tea or not
func isSufficientSugarSyrupPresentForGinger(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.SugarSyrup >= inputData.Machine.Beverages.GingerTea.SugarSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet Leaves Syrup availble to make ginger tea or not
func isSufficientLeavesSyrupPresentForGinger(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup >= inputData.Machine.Beverages.GingerTea.TeaLeavesSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet Hot water availble to make Elacihi tea or not
func isSufficientHotWaterPresentForElaichi(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotWater >= inputData.Machine.Beverages.ElaichiTea.HotWater {
		isSuff = true
	}
	return
}

//check Suffcinet Hot Milk availble to make elachi tea or not
func isSufficientHotMilkPresentForElaichi(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotMilk >= inputData.Machine.Beverages.ElaichiTea.HotMilk {
		isSuff = true
	}
	return
}

//check Suffcinet elachi syrup availble to make elaichi tea or not
func isSufficientElaichiSyrupPresentForElaichi(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.ElaichiSyrup >= inputData.Machine.Beverages.ElaichiTea.ElaichiSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet sugar syrup availble to make elaichi tae or not
func isSufficientSugarSyrupPresentForElaichi(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.SugarSyrup >= inputData.Machine.Beverages.ElaichiTea.SugarSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet leave syrup availble to make elaichi tae or not
func isSufficientLeavesSyrupPresentForElaichi(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.TeaLeavesSyrup >= inputData.Machine.Beverages.ElaichiTea.TeaLeavesSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet Hot milk availble to make  coffee or not
func isSufficientHotMilkPresentForCofee(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotMilk >= inputData.Machine.Beverages.Coffee.HotMilk {
		isSuff = true
	}
	return
}

//check Suffcinet coffeesyrup availble to make  coffee or not
func isSufficientCofeeSyrupPresentForCofee(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.CoffeeSyrup >= inputData.Machine.Beverages.Coffee.CoffeeSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet sugar syrup availble to make  coffee or not
func isSufficientSugarSyrupPresentForCofee(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.SugarSyrup >= inputData.Machine.Beverages.Coffee.SugarSyrup {
		isSuff = true
	}
	return
}

//check Suffcinet Hot water availble to make  coffee or not
func isSufficientHotWaterPresentForCofee(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotWater >= inputData.Machine.Beverages.Coffee.HotWater {
		isSuff = true
	}
	return
}

//check Suffcinet water availble to make  hot water
func isSufficientWaterPresentForHotWater(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotWater >= inputData.Machine.Beverages.Water.HotWater {
		isSuff = true
	}
	return
}

//check Suffcinet milk availble to make  Hot milk
func isSufficientMilkPresentForHotMilk(inputData CoffeeMachine) (isSuff bool) {
	if inputData.Machine.TotalItemsQuantity.HotMilk >= inputData.Machine.Beverages.Milk.HotMilk {
		isSuff = true
	}
	return
}
