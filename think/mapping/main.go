package main

import "fmt"

func main() {
	
	// maps pydict
	presAge := make(map[string] int)
	
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(len(presAge))
	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge))
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))
	fmt.Println(presAge)
	
	
	// We can store multiple items in a map as well
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname":"Clark Kent",
			"city":"Metropolis",
		},
		"Batman": map[string]string{
			"realname":"Bruce Wayne",
			"city":"Gotham City",
		},
	}
	// We can output data where the key matches Superman
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}
}
