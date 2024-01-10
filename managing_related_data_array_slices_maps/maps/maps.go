package maps

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

//map vs struct

// Map    -  any vlaue can be a key , string, int, struct....const
// Struct - predefine data structure
// https://articles.wesionary.team/map-vs-struct-in-golang-when-to-use-b0b66446627a

// Maps vs. Structs
// For map:
// - All key and value are of same type.
// - When keys are indexed and we can iterate over them.
// - Closely related and significant value type.
// - Don’t need to know all the keys at compile time.
// - Key are indexed- we can iterate over them.
// - Reference type

// For struct:
// - All values can be of different type.
// - Need to know all the different fields at compile time.
// - Keys don’t support indexing
// - Value type.

// When to use?
// When to use structs? If we have close set of keys means the fixed data size with keys we
// will be using structs. Using structs are safe way and easy way while working with JSON data also.

// When to use maps? If we are creating some kind of relationship between keys and values and we don’t
// really know what that collection of values going to be at compile time or as we are writing our code
// then we got the great use-case of using a map.

// Most of cases, vast majority of time we mostly, use structs than maps. But it all depends on nature
// and type of the application and requirement of the project.
