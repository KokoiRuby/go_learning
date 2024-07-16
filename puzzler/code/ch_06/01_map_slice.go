package main

import "fmt"

func main() {
	// make map slice
	msl := make([]map[string]string, 1)
	if msl[0] == nil {
		// make map (in the slice)
		msl[0] = make(map[string]string, 3)
		msl[0]["k1"] = "v1"
		msl[0]["k2"] = "v2"
		msl[0]["k3"] = "v3"
	}
	fmt.Println(msl)

	// append if cap is reached
	newMap := map[string]string{
		"k4": "v4",
		"k5": "v5",
		"k6": "v6",
	}
	msl = append(msl, newMap)
	println(msl)
}
