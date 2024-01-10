package main

import "fmt"

//----------Working with Aliases-------

//-------Type Aliases
type floatMap map[string]float64

//add a function to it - (m floatMap) is the receiver
func (m floatMap) output() {
	fmt.Println(m)
}

//--------Make - array----

func main() {

	//userNames := make([]string, 2) //create an array of length 2
	userNames := make([]string, 2, 5) //create an array of length 2 with a capacity of 5

	userNames = append(userNames, "Jeremy")
	userNames = append(userNames, "Nathan")
	fmt.Println(userNames) //the first 2 are blank because when we setup the array we gave it an initial value of 2 and then append in 2 more users [ ,,"jeremy","nathan"]

	//---------Making Maps------
	//courseRatings := map[string]float64{} // this is the Map literal
	//every time we add to the array Go has to reallocate some memory

	courseRatings := make(floatMap, 3) // this is the Map with make
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["node"] = 4.7 //since this one is beyond our initial capacity go will have to allocate some memory for this one.

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
