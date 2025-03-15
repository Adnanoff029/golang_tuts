// Map keys may be of any type that is comparable.
// boolean, numeric, string, pointer, channel, and interface types, and structs or arrays
// Map also hold references to an underlying data structure.
// If some key is not present the the map returns with 0. Hence we use another variable to indicate whether the key is present or not.

package main

import "fmt"

type entry struct {
	v1 string
	v2 string
}

func main() {
	mp := make(map[int]int)
	mp[0] = 100
	mp[1] = 101
	mp[7] = 1002
	mp[0] = 10
	for key, value := range mp {
		fmt.Println(key, value)
	}
	// Deleting an element from mp delete (map_name, key_value)
	delete(mp, 0)

	fmt.Println(mp)

	val, ok := mp[0]
	if !ok {
		fmt.Println("No value found!")
	} else {
		fmt.Println(val)
	}

	mp2 := map[string]int{
		"Adnan": 10,
		"Khan":  8,
	}
	for key, value := range mp2 {
		fmt.Println(key, value)
	}

	mp3 := make(map[entry]int)
	mp3[entry{v1: "Adnan", v2: "Khan"}] = 9
	mp3[entry{v1: "Saquib", v2: "Khan"}] = 10
	searchKey := entry{v1: "Saquib", v2: "Khan"}
	if value, ok := mp3[searchKey]; ok {
		fmt.Printf("%v : %v \n", searchKey, value)
	}

	// Nested maps
	mp4 := make(map[string]map[string]int)
	mp4["Adnan"] = make(map[string]int)
	mp5 := make(map[string]int)
	mp5["Khan"] = 9
	mp5["Hussain"] = 12
	mp4["Adnan"] = mp5
	for key, value := range mp4 {
		fmt.Printf("Values for %v :: ", key)
		for key2, value2 := range value {
			fmt.Printf("(%v, %v) : ", key2, value2)
		}
		fmt.Printf("\n")
	}
}
