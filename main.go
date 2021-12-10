//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
//
// Author Bhakiyaraj Kalimuthu
//

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	m := map[string]int{}
	go func() {
		m["hello"] = 1
		m["world"] = 2
	}()
	time.Sleep(5 * time.Second)
	fmt.Println(m)
}

func locFunc() string {
	return func() string {
		return fmt.Sprint("im a local func ")
	}()
}
func main1() {

	zoneInfo := map[string]map[string]int64{}
	for _, s := range getZones() {
		zoneInfo[s] = make(map[string]int64)
	}

	for _, s := range getServers() { // slice
		for k, v := range s { // go_map 1
			numOfServers := zoneInfo[k] // key is zone value a, b, c
			for _, v1 := range v {      // go_map 2 - key is server name
				if numOfServers[v1] <= 1 {
					numOfServers[v1] += 1
				}
				if numOfServers[v1] > 1 {
					log.Fatalf("services are not evenly distributed")
				}
			}

		}
	}

	fmt.Println(zoneInfo)
}

func getZones() []string {

	return []string{"a", "b", "c"}
}

func getServers() []map[string]map[string]string {
	s := make([]map[string]map[string]string, 3)

	m1 := make(map[string]map[string]string)
	m11 := make(map[string]string)
	m11["agent"] = "agent"
	m11["coordinator"] = "coordinator"
	m11["db"] = "db"

	m11["coordinator1"] = "coordinator"
	m11["db1"] = "db"

	m1["a"] = m11
	m2 := make(map[string]map[string]string)
	m22 := make(map[string]string)
	m22["agent"] = "agent"
	m22["coordinator"] = "coordinator"
	m22["db"] = "db"
	m2["b"] = m22

	m3 := make(map[string]map[string]string)
	m33 := make(map[string]string)
	m33["agent"] = "agent"
	m33["coordinator"] = "coordinator"
	m33["db"] = "db"
	m2["c"] = m33

	s[0] = m1
	s[1] = m2
	s[2] = m3

	return s
}
func getServers1() []map[string]map[string]int64 {
	s := make([]map[string]map[string]int64, 3)

	m1 := make(map[string]map[string]int64)
	m11 := make(map[string]int64)
	m11["agent"] = 1
	m11["coordinator"] = 1
	m11["db"] = 2
	m1["a"] = m11

	m2 := make(map[string]map[string]int64)
	m22 := make(map[string]int64)
	m22["agent"] = 1
	m22["coordinator"] = 1
	m22["db"] = 0
	m2["b"] = m22

	m3 := make(map[string]map[string]int64)
	m33 := make(map[string]int64)
	m33["agent"] = 1
	m33["coordinator"] = 1
	m33["db"] = 1
	m2["c"] = m33

	s[0] = m1
	s[1] = m2
	s[2] = m3

	return s
}
