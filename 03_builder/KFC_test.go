package main

import "testing"

func TestMealBuilder(t *testing.T) {
	mealBuilder := MealBuilder{}

	vegMeal := mealBuilder.PrepareVegMeal()
	vegMeal.ShowMeals()

	nonVegMeal := mealBuilder.prepareNonVegMeal()
	nonVegMeal.ShowMeals()
}
