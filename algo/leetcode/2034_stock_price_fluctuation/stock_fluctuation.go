package main

import "fmt"

type StockPrice struct {
	recordMap map[int]int
	records   []int
	current   record
}

type record struct {
	timestamp, price int
}

func Constructor() StockPrice {
	return StockPrice{make(map[int]int), make([]int, 0), record{}}
}

func (this *StockPrice) Update(timestamp int, price int) {

	r := record{timestamp, price}

	if len(this.records) == 0 {
		this.current = r
		this.recordMap[timestamp] = price
		this.records = append(this.records, price)
		return
	}
	// If latest timestamp
	if timestamp >= this.current.timestamp {
		this.current = r
	}
	if p, ok := this.recordMap[timestamp]; ok {
		this.recordMap[timestamp] = price
		this.priceSearchAndInsert(p, price)
		this.quickSortOnPrice(0, len(this.records)-1)
		return
	}
	this.recordMap[timestamp] = price
	this.priceSearchAndInsert(price, 0)
}

func (this *StockPrice) quickSortOnPrice(left, right int) {
	if left < right {
		p := this.partition(left, right)
		this.quickSortOnPrice(left, p-1)
		this.quickSortOnPrice(p+1, right)
	}
}

func (this *StockPrice) partition(left, right int) int {
	pivot := this.records[right]
	pIndex := left
	for i := left; i < right; i++ {
		if this.records[i] < pivot {
			this.swap(i, pIndex)
			pIndex++
		}
	}
	this.swap(pIndex, right)
	return pIndex
}

func (this *StockPrice) swap(i, j int) {
	temp := this.records[i]
	this.records[i] = this.records[j]
	this.records[j] = temp
}

func (this *StockPrice) priceSearchAndInsert(record, update int) {
	left, right := 0, len(this.records)-1
	for i := 0; i < len(this.records); i++ {
		mid := (left + right) / 2
		if record == this.records[mid] {
			if update != 0 {
				this.records[mid] = update
			}
			return
		}
		if record < this.records[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	this.insertAt(left, record)
	return
}

func (this *StockPrice) insertAt(index, record int) {
	if index == len(this.records) {
		this.records = append(this.records, record)
		return
	}
	this.records = append(this.records[:index+1], this.records[index:]...)
	if this.records[index] < record {
		this.records[index+1] = record
	} else {
		this.records[index] = record
	}
	return
}

func (this *StockPrice) Current() int {
	return this.current.price
}

func (this *StockPrice) Maximum() int {
	max := 0
	if len(this.records) != 0 {
		max = this.records[len(this.records)-1]
	}
	return max
}

func (this *StockPrice) Minimum() int {
	min := 0
	if len(this.records) != 0 {
		min = this.records[0]
	}
	return min
}

func main() {

	//["StockPrice", "update", "maximum", "current", "minimum","maximum", "maximum", "maximum", "minimum", "minimum","maximum",
	//"update", "maximum", "minimum",
	//"update","maximum", "minimum", "current", "maximum",
	//"update","minimum", "maximum",
	//"update", "maximum", "maximum","current",
	//"update", "current", "minimum",
	//"update",
	//"update", "minimum", "minimum",

	//"update", "current",
	//"update", "maximum",
	//"update", "minimum"]
	//[[], [38, 2308], [], [], [], [], [], [], [],[], [], [47, 7876], [], [], [58, 1866], [], [],[], [], [43, 121], [], [], [40, 5339], [], [],[], [32, 5339], [], [],
	//[43, 6414], [49, 9369], [],[], [36, 3192], [], [48, 1006], [], [53, 8013], []]

	//[null,null,2308,2308,2308,2308,2308,2308,2308,2308,2308,null,7876,2308,
	// null,7876,1866,1866,7876,null,121,7876,null,7876,7876,
	//1866,null,1866,121,null,null,1866,1866,null,1866,null,9369,null,1006]

	c := Constructor()
	c.Update(38, 2308)
	fmt.Printf("Expected -> 2308 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 2308 Current -> %d \n", c.Current())
	fmt.Printf("Expected -> 2308 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 2308 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 2308 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 2308 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 2308 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 2308 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 2308 Maximum -> %d \n", c.Maximum())
	c.Update(47, 7876)
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 2308 Minimum -> %d \n", c.Minimum())
	c.Update(58, 1866)
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 1866 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 1866 Current -> %d \n", c.Current())
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())

	c.Update(43, 121)
	fmt.Printf("Expected -> 121 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())

	c.Update(40, 5339)
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 7876 Maximum -> %d \n", c.Maximum())
	fmt.Printf("Expected -> 1866 Current -> %d \n", c.Current())

	c.Update(32, 5339)
	fmt.Printf("Expected -> 1866 Current -> %d \n", c.Current())
	fmt.Printf("Expected -> 121 Minimum -> %d \n", c.Minimum())

	c.Update(43, 6414)
	c.Update(49, 9369)
	fmt.Printf("Expected -> 1866 Minimum -> %d \n", c.Minimum())
	fmt.Printf("Expected -> 1866 Minimum -> %d \n", c.Minimum())

	c.Update(36, 3192)
	fmt.Printf("Expected -> 1866 Current -> %d \n", c.Current())

	c.Update(48, 1006)
	fmt.Printf("Expected -> 9369 Maximum -> %d \n", c.Maximum())
	c.Update(53, 8013)
	fmt.Printf("Expected -> 1006 Minimum -> %d \n", c.Minimum())

	//["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
	//	[[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
	s := Constructor()
	s.Update(1, 10)
	s.Update(2, 5)
	fmt.Printf("Current -> %d \n", s.Current())
	fmt.Printf("Maximum -> %d \n", s.Maximum())
	s.Update(1, 3)
	fmt.Printf("Current -> %d \n", s.Current())
	fmt.Printf("Maximum -> %d \n", s.Maximum())
	s.Update(4, 2)
	fmt.Printf("Current -> %d \n", s.Current())
	fmt.Printf("Minimum -> %d \n", s.Minimum())
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
