package main

import (
	"fmt"
)

type Product struct {
	title string
	id    int64
	price float64
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.

	hobbies := [3]string{"learning", "home brewing", "golfing"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list

	fmt.Println(hobbies[0])
	lastTwoHobbies := hobbies[1:]
	fmt.Println(lastTwoHobbies)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	sliceTwoElems := hobbies[0:2]
	fmt.Println(sliceTwoElems)
	sliceTwoElems = hobbies[:2]
	fmt.Println(sliceTwoElems)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	sliceTwoElems = hobbies[1:]
	fmt.Println(sliceTwoElems)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"understand GO better", "Finish the course"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "New goal"
	goals = append(goals, "This is my 3rd")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	//      we can omit the struct name
	poducts := []Product{{title: "titl1", id: 1, price: 2.3}, {title: "titl2", id: 2, price: 3.4}}
	poducts = append(poducts, Product{title: "titl3", id: 3, price: 4.5})
	fmt.Println(poducts)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.

// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list

// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)

// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.

// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
