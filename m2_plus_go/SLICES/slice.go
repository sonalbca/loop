package main

import (
	"fmt"
	"reflect"
)

/*
Golang Slices
Introduction of Slices, managing collections of data with slices and adding and removing elements from a slice.

Introduction
A slice is a flexible and extensible data structure to implement and manage collections of data.
Slices are made up of multiple elements, all of the same type.
A slice is a segment of dynamic arrays that can grow and shrink as you see fit. Like arrays, slices are index-able and have a length. Slices have a capacity and length property.
*/

func main() {
	//Slice1() //Create Empty Slice
	//Slice2() //Declare Slice using Make
	//Slice3()          //Initialize Slice with values using a Slice Literal
	//Slice4()			//Declare Slice using new Keyword
	//Slice5()			//Add Items
	//Slice6()          //Access Items
	//Slice7()			//Change Item Value
	Slice8() //Remove Item from Slice
	//Slice9() //Copy a Slice
	//Slice10() //Tricks of Slicing
	//Slice11()			//Loop Through a Slice
	//Slice12()			//Append a slice to an existing slice
	//Slice13()			//Check if Item Exists

}

//-------------------------------------------- 1 -----------------------------------------//
//Create Empty Slice
//To declare the type for a variable that holds a slice, use an empty pair of square brackets,
//followed by the type of elements the slice will hold.

func Slice1() {
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

/*
Output
slice
slice
*/

//-------------------------------------------- 2 -----------------------------------------//
//Declare Slice using Make
//Slice can be created using the built-in function make. When you use make, one option you have is to specify the length of the slice.
//When you just specify the length, the capacity of the slice is the same.

func Slice2() {
	var intSlice = make([]int, 10)        // when length and capacity is same
	var strSlice = make([]string, 10, 20) // when length and capacity is different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

/*
Output
intSlice        Len: 10         Cap: 10
slice
strSlice        Len: 10         Cap: 20
slice
*/

//-------------------------------------------- 3 -----------------------------------------//
//Initialize Slice with values using a Slice Literal

//A slice literal contain empty brackets followed by the type of elements the slice will hold,
//and a list of the initial values each element will have in curly braces.

func Slice3() {
	var intSlice = []int{10, 20, 30, 40}
	var strSlice = []string{"India", "Canada", "Japan"}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}

//-------------------------------------------- 4 -----------------------------------------//
//Declare Slice using new Keyword

//A slice can be declare using new keyword followed by capacity in square brackets then type of elements the slice will hold.

func Slice4() {
	var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}

/*
Output
slice
intSlice        Len: 10         Cap: 50
[0 0 0 0 0 0 0 0 0 0]
*/

//-------------------------------------------- 5 -----------------------------------------//
//Add Items
//To add an item to the end of the slice, use the append() method.

func Slice5() {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
}

/*
Output
Slice A: [10 20]
Length is 2 Capacity is 5
Slice A after appending data: [10 20 30 40 50 60 70 80 90]
Length is 9 Capacity is 12
*/
//If there's sufficient capacity in the underlying slice, the element is placed after the last element and the length get incremented. However,
//if there is not sufficient capacity, a new slice is created, all of the existing elements are copied over, the new element is added onto the end,
//and the new slice is returned.

//-------------------------------------------- 6 -----------------------------------------//
//Access Items
//You access the slice items by referring to the index number.

func Slice6() {
	var intSlice = []int{10, 20, 30, 40}

	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])
}

//-------------------------------------------- 7 -----------------------------------------//
//Change Item Value
//To change the value of a specific item, refer to the index number.

func Slice7() {
	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Println(strSlice)

	strSlice[2] = "Germany"
	fmt.Println(strSlice)
}

/*
Output
[India Canada Japan]
[India Canada Germany]
*/

//-------------------------------------------- 8 -----------------------------------------//
//Remove Item from Slice
//RemoveIndex function created to remove specific item from String slice.

func Slice8() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

/*
Output
[India Canada Japan Germany Italy]
[India Canada Japan Italy]
*/

//-------------------------------------------- 9 -----------------------------------------//
//Copy a Slice
//The built-in copy function is used to copy data from one slice to another.

func Slice9() {
	a := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, a)              // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}

/*
Output
[Slice:A] Length is 3 Capacity is 3
[Slice:B] Length is 5 Capacity is 10
Slice B after copying: [5 6 7 0 0]
Slice B after adding elements: [5 6 7 8 9]
*/

//-------------------------------------------- 10 -----------------------------------------//
//Tricks of Slicing
//Slicing is a computationally fast way to methodically access parts of your data.

func Slice10() {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])
}

/*
Output
Countries: [india japan canada australia russia]
:2 [india japan]
1:3 [japan canada]
2: [canada australia russia]
2:5 [canada australia russia]
0:3 [india japan canada]
Last element: russia
Last element: russia
Last element: [russia]
All elements: [india japan canada australia russia]
Last two elements: [australia russia]
Last two elements: [australia russia]
[india japan canada australia russia]
[india japan canada australia russia]
[india japan canada australia russia]
*/

//-------------------------------------------- 11 -----------------------------------------//
//Loop Through a Slice
//You can loop through the list items by using a for loop.

func Slice11() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Println("\n---------------Example 1 --------------------\n")
	for index, element := range strSlice {
		fmt.Println(index, "--", element)
	}

	fmt.Println("\n---------------Example 2 --------------------\n")
	for _, value := range strSlice {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 3 --------------------\n")
	for range strSlice {
		fmt.Println(strSlice[j])
		j++
	}
}

//-------------------------------------------- 12 -----------------------------------------//
//Append a slice to an existing slice
//The usage of triple-dot ... ellipsis used to append a slice.

func Slice12() {
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice2 = append(slice2, slice1...)
}

//-------------------------------------------- 13 -----------------------------------------//
//Check if Item Exists
//To determine if a specified item is present in a slice iterate slice item and check using if condition.

func Slice13() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
