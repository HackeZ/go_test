package main

import (
	"errors"
	"fmt"
)

func main() {
	main1()
}

// === Adder Closure ===
func adder() func(int) int {
	sum := 0

	innerFunc := func(x int) int {
		sum += x
		return sum
	}
	return innerFunc
}

func main1() {
	pos, neg := adder(), adder()
	for i := 0; i <= 10; i++ {
		fmt.Println(pos(i), " ", neg(-2*i))
	}
}

// === Adder Closure ===

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

// ===- Slice Closure -===
type Traveser func(ele interface{})

func Process(array interface{}, traveser Traveser) error {
	if array == nil {
		return errors.New("nil point")
	}

	var length int
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))
	default:
		return errors.New("unsupport type")
	}

	if length == 0 {
		return errors.New("array length is zero")
	}
	traveser(array)
	return nil
}

// QuickSort ...
// array []int
// left 0
// right len(array)-1
// smallTobig bool
func QuickSort(array []int, left, right int, smallTobig bool) {
	if left > right {
		return
	}

	s := right
	i, j := left-1, left

	for j <= s {
		if smallTobig {
			if array[j] >= array[s] {
				array[i+1], array[j] = array[j], array[i+1]
				i++
			}
			j++
		} else {
			if array[j] <= array[s] {
				array[i+1], array[j] = array[j], array[i+1]
				i++
			}
			j++
		}
	}

	QuickSort(array, left, i-1, smallTobig)
	QuickSort(array, i+1, right, smallTobig)
}

// SortByAscending 升序
func SortByAscending(ele interface{}) {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)

	QuickSort(intSlice, 0, length-1, true)
}

// SortByDescending 降序
func SortByDescending(ele interface{}) {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)

	QuickSort(intSlice, 0, length-1, false)
}

func main2() {

	// intSlice
	intSlice := []int{3, 1, 2, 4}

	fmt.Printf("Now is slice of int\n")

	Process(intSlice, SortByAscending)
	fmt.Println(intSlice)

	Process(intSlice, SortByDescending)
	fmt.Println(intSlice)

	// stringSlice
	stringSlice := []string{"hello", "world", "golang"}

	Process(stringSlice, func(ele interface{}) {
		if slice, ok := ele.([]string); ok {
			fmt.Printf("\nNow is slice of string\n")
			for k, v := range slice {
				fmt.Println("index: ", k, " => value: ", v)
			}
		}
	})

	// floatSlice
	floatSlice := []float32{2.4, 21.5, 1.2}

	Process(floatSlice, func(ele interface{}) {
		if slice, ok := ele.([]float32); ok {
			for k := range slice {
				slice[k] *= 2
			}
		}
	})

	fmt.Printf("\nNow is slice of float32\n")
	fmt.Println(floatSlice)
}

// ===- Slice Closure -===

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

//  ===- Filter Closure -===

type FilterFunc func(ele interface{}) interface{}

func Data(array interface{}, filterFunc FilterFunc) interface{} {
	slice := make([]int, 0)
	arr, _ := array.([]int)

	for _, v := range arr {
		if result, ok := filterFunc(v).(int); ok {
			slice = append(slice, result)
		}
	}

	return slice
}

func EvenFilter(ele interface{}) interface{} {
	if integer, ok := ele.(int); ok {
		if integer%2 == 1 {
			integer++
		}
		return integer
	}
	return ele
}

func OddFilter(ele interface{}) interface{} {
	if integer, ok := ele.(int); ok {
		if integer%2 == 0 {
			integer++
		}
		return integer
	}
	return ele
}

func main3() {
	sliceEven := []int{1, 2, 3, 4, 5}
	sliceEven = Data(sliceEven, EvenFilter).([]int)
	fmt.Println("EvenFilter =>", sliceEven)

	sliceOdd := []int{1, 2, 3, 4, 5}
	sliceOdd = Data(sliceOdd, OddFilter).([]int)
	fmt.Println("OddFilter =>", sliceOdd)
}

//  ===- Filter Closure -===
