package main

import "fmt"

//------------Maps--------

func main() {
	//declaring an empty map
	//websites := map[string]string{}

	websites := map[string]string{
		"Goggle": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)

	//exact an key-value pair
	fmt.Println(websites["AWS"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

}
