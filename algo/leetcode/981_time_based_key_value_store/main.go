package main

import "fmt"

type TimeMap map[string][]data

type data struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return make(map[string][]data, 1024)
}

func (this TimeMap) Set(key string, value string, timestamp int) {
	if val, ok := this[key]; ok {
		val = append(val, data{value, timestamp})
		this[key] = val
		return
	}
	this[key] = []data{{value, timestamp}}

}

func (this TimeMap) Get(key string, timestamp int) string {
	if val, ok := this[key]; ok {
		if len(val) == 0 {
			return ""
		}
		// binary search
		left, right, mid := 0, len(val)-1, 0
		for {
			if left <= right {
				mid = left + right/2
				if timestamp == val[mid].timestamp {
					return val[mid].value
				}
				if timestamp < val[mid].timestamp {
					right = mid - 1
				} else {
					left = mid + 1
				}

			} else {
				if right >= 0 {
					return val[right].value
				}
				break
			}

		}
	}
	return ""
}

func main() {

	//["TimeMap","set","set","get","set","get","set","set","get","set","get","set","get","set","get","set","get","get","get","get","get","get","set","set","set","get","get","set","set","get","set"]
	//	[[],["rtzoj","kuexwze",1],["xcywxndnz","herqmazp",2],["xcywxndnz",3],["rtzoj","dgpguflin",4],["xcywxndnz",5],["dgpguflin","lvrexco",6],["xcywxndnz","dgpguflin",7],["xcywxndnz",8],["rtzoj","wxqixmxs",9],["xcywxndnz",10],["kuexwze","lvrexco",11],["dgpguflin",12],["lvrexco","wxqixmxs",13],["xcywxndnz",14],["herqmazp","vjfhio",15],["dgpguflin",16],["herqmazp",17],["herqmazp",18],["rtzoj",19],["herqmazp",20],["herqmazp",21],["kuexwze","vjfhio",22],["dgpguflin","qrkihrb",23],["kuexwze","dgpguflin",24],["rtzoj",25],["dgpguflin",26],["herqmazp","rtzoj",27],["lvrexco","iztpo",28],["lvrexco",29],["kuexwze","lvrexco",30]]
	
	//["TimeMap","set","set","get","get","get","get","get"]
	//[[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]

	//output: [null, null, null, "low", "low","low", "low", "low"]
	//expected: [null, null, null, "", "high", "high", "low", "low"]

	m := Constructor()
	m.Set("love", "high", 10)
	m.Set("love", "low", 20)
	fmt.Println(m.Get("love", 5))
	fmt.Println(m.Get("love", 10))
	fmt.Println(m.Get("love", 15))
	fmt.Println(m.Get("love", 20))
	fmt.Println(m.Get("love", 25))
	//["TimeMap", "set", "get", "get", "set", "get", "get"]
	//	[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]

	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)       // store the key "foo" and value "bar" along with timestamp = 1.
	fmt.Println(timeMap.Get("foo", 1)) // return "bar"
	fmt.Println(timeMap.Get("foo", 3)) // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	timeMap.Set("foo", "bar2", 4)      // store the key "foo" and value "bar2" along with timestamp = 4.
	fmt.Println(timeMap.Get("foo", 4)) // return "bar2"
	fmt.Println(timeMap.Get("foo", 5)) // return "bar2"
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
