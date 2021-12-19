package main

import "fmt"

type StockPrice struct {
	records []record
	current int
}

type record struct {
	timestamp, price int
}

func Constructor() StockPrice {
	return StockPrice{make([]record, 0), 0}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.current = price
	r := record{timestamp, price}
	if len(this.records) == 0 {
		this.records = append(this.records, r)
		return
	}
	if ok, index := this.timestampSearch(r); ok {
		this.records[index].price = price
		this.quickSortOnPrice(0, len(this.records)-1)
		return
	}
	this.priceSearchAndInsert(r)
}

func (this *StockPrice) quickSortOnPrice(left, right int) {
	if left < right {
		p := this.partition(left, right)
		this.quickSortOnPrice(left, p-1)
		this.quickSortOnPrice(p+1, right)
	}
}

func (this *StockPrice) partition(left, right int) int {
	pivot := this.records[right].price
	pIndex := left
	for i := left; i < right; i++ {
		if this.records[i].price < pivot {
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

func (this *StockPrice) timestampSearch(record record) (bool, int) {
	for i := 0; i < len(this.records); i++ {
		if this.records[i].timestamp == record.timestamp {
			return true, i
		}
	}
	return false, 0
}

func (this *StockPrice) priceSearchAndInsert(record record) {
	left, right := 0, len(this.records)-1
	for i := 0; i < len(this.records); i++ {
		mid := (left + right) / 2
		if record.price == this.records[mid].price {
			this.records[mid].price = record.price
			return
		}
		if record.price < this.records[mid].price {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	this.insertAt(left, record)
	return
}

func (this *StockPrice) insertAt(index int, record record) {
	if index == len(this.records) {
		this.records = append(this.records, record)
		return
	}
	this.records = append(this.records[:index+1], this.records[index:]...)
	if this.records[index].price < record.price {
		this.records[index+1] = record
	} else {
		this.records[index] = record
	}
	return
}

func (this *StockPrice) Current() int {
	return this.current
}

func (this *StockPrice) Maximum() int {
	max := 0
	if len(this.records) != 0 {
		max = this.records[len(this.records)-1].price
	}
	return max
}

func (this *StockPrice) Minimum() int {
	min := 0
	if len(this.records) != 0 {
		min = this.records[0].price
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
	fmt.Printf("Expected -> 5339 Current -> %d \n", c.Current())

	c.Update(32, 5339)
	fmt.Printf("Current -> %d \n", c.Current())
	fmt.Printf("Minimum -> %d \n", c.Minimum())

	c.Update(43, 6414)
	c.Update(49, 9369)
	fmt.Printf("Minimum -> %d \n", c.Minimum())
	fmt.Printf("Minimum -> %d \n", c.Minimum())

	c.Update(36, 3192)
	fmt.Printf("Current -> %d \n", c.Current())

	c.Update(48, 1006)
	fmt.Printf("Maximum -> %d \n", c.Maximum())
	c.Update(53, 8013)
	fmt.Printf("Minimum -> %d \n", c.Minimum())

	//["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
	//	[[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
	//s := Constructor()
	//	//s.Update(1, 10)
	//	//s.Update(2, 5)
	//	//fmt.Printf("Current -> %d \n", s.Current())
	//	//fmt.Printf("Maximum -> %d \n", s.Maximum())
	//	//s.Update(1, 3)
	//	//fmt.Printf("Current -> %d \n", s.Current())
	//	//fmt.Printf("Maximum -> %d \n", s.Maximum())
	//	//s.Update(4, 2)
	//	//fmt.Printf("Current -> %d \n", s.Current())
	//	//fmt.Printf("Minimum -> %d \n", s.Minimum())
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
