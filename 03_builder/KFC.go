package main

import "fmt"

type IPack interface {
	pack()
}

type IBurger interface {
	GetName() string
	Packing() string
	GetPrice() float64
}

type IDrink interface {
	GetName() string
	Packing() string
	GetPrice() float64
}

type Wrapper struct {
}

func (w Wrapper) pack() string {
	return "袋装食品打包中......"
}

type Bottle struct {
}

func (b Bottle) pack() string {
	return "瓶装食品打包中......"
}

type VegBurger struct {
}

func (v VegBurger) GetName() string {
	return "Veg Burger"
}

func (v VegBurger) GetPrice() float64 {
	return 25.0
}

func (v VegBurger) Packing() string {
	return Wrapper{}.pack()
}

type ChickenBurger struct {
}

func (c ChickenBurger) GetName() string {
	return "Chicken Burger"
}

func (c ChickenBurger) GetPrice() float64 {
	return 50.0
}

func (c ChickenBurger) Packing() string {
	return Wrapper{}.pack()
}

type Coke struct {
}

func (c Coke) GetName() string {
	return "Coke"
}

func (c Coke) GetPrice() float64 {
	return 30.0
}

func (c Coke) Packing() string {
	return Bottle{}.pack()
}

type Pepsi struct {
}

func (p Pepsi) GetName() string {
	return "Pepsi"
}

func (p Pepsi) GetPrice() float64 {
	return 40.0
}

func (p Pepsi) Packing() string {
	return Bottle{}.pack()
}

type Meal struct {
	BurgerList []IBurger
	DrinkList  []IDrink
}

func (m *Meal) AddBurger(burger IBurger) {
	if m.BurgerList == nil {
		m.BurgerList = []IBurger{burger}
	} else {
		m.BurgerList = append(m.BurgerList, burger)
	}
}

func (m *Meal) AddDrink(drink IDrink) {
	if m.DrinkList == nil {
		m.DrinkList = []IDrink{drink}
	} else {
		m.DrinkList = append(m.DrinkList, drink)
	}
}

func (m *Meal) GetMealCost() float64 {
	mealCost := 0.0
	for _, drink := range m.DrinkList {
		mealCost += drink.GetPrice()
	}
	for _, burger := range m.BurgerList {
		mealCost += burger.GetPrice()
	}
	return mealCost
}

func (m *Meal) ShowMeals() {
	fmt.Println("套餐准备中----------------------")
	for _, drink := range m.DrinkList {
		fmt.Printf("Item : %v , Packing : %s , Price : %5.2f\n", drink.GetName(), drink.Packing(), drink.GetPrice())
	}
	for _, burger := range m.BurgerList {
		fmt.Printf("Item : %v , Packing : %s , Price : %5.2f\n", burger.GetName(), burger.Packing(), burger.GetPrice())
	}
	fmt.Println("套餐准备完毕----------------------")
}

type MealBuilder struct {
}

func (m MealBuilder) PrepareVegMeal() Meal {
	meal := Meal{}
	meal.AddBurger(VegBurger{})
	meal.AddDrink(Coke{})
	return meal
}

func (m MealBuilder) prepareNonVegMeal() Meal {
	meal := Meal{}
	meal.AddBurger(ChickenBurger{})
	meal.AddDrink(Pepsi{})
	return meal
}

func main() {
	mealBuilder := MealBuilder{}

	vegMeal := mealBuilder.PrepareVegMeal()
	vegMeal.ShowMeals()

	nonVegMeal := mealBuilder.prepareNonVegMeal()
	nonVegMeal.ShowMeals()
}
