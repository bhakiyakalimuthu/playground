package main

import "fmt"

func main() {
	//TODO(Note): can't create constant map  const x = map[interface{}]interface{}
	// declaration
	var x map[interface{}]interface{}
	fmt.Printf("x interface: %v\n", x)
	var map1 map[string]struct{}
	map2 := make(map[string]struct{})
	map3 := map[string]struct{}{"testplay": {}}
	for k, v := range map1 {
		fmt.Printf("map1 - k: %v\n v: %v\n", k, v)
	}
	for k, v := range map2 {
		fmt.Printf("map2 - k: %v\n v: %v\n", k, v)
	}
	for k, v := range map3 {
		fmt.Printf("map3 - k: %v\n v: %v\n", k, v)
	}

	if val, ok := map2["testplay"]; ok {
		fmt.Printf("val: %v\n", val)
	}
	//fmt.Printf("map1: %v\n map2: %v\n map2: %v\n", map1, map2, map3)
	nameMap := make(map[string]struct{}, 26*26*10*10*10)
	fmt.Printf("len of nameMap: %d\n", len(nameMap))

	countryCodes := map[string]string{
		"93":  "Afghanistan",
		"374": "Armenia",
		"61":  "Australia",
	}

	countryCodes["6"] = "sweden"
	if country, ok := countryCodes["6"]; ok {
		fmt.Println(country, "is a beautiful place")
	} else {
		fmt.Println("I don't know that country")
	}
	fmt.Println(countryCodes["6"])
}
