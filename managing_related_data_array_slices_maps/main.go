package main

import "fmt"

type floatMap

//--------Make - array----

func main() {

	//userNames := make([]string, 2) //create an array of length 2
	userNames := make([]string, 2, 5) //create an array of length 2 with a capacity of 5

	userNames = append(userNames, "Jeremy")
	userNames = append(userNames, "Nathan")
	fmt.Println(userNames)

	//---------Making Maps------
	//courseRatings := map[string]float64{} // this is the Map literal
	//every time we add to the array Go has to reallocate some memory

	courseRatings := make(map[string]float64, 3) // this is the Map with make
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["node"] = 4.7 //since this one is beyond our initial capacity go will have to allocate some memory for this one. 

	//----------Working with Aliases-------

}
