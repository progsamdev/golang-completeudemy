package main

import "fmt"

func main() {

	// 1) Create a new array with three hobbies
	myHobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println("My hobbies:", myHobbies)

	// 2) Output specific elements
	fmt.Println("First hobby:", myHobbies[0])
	fmt.Println("Second and third hobbies:", myHobbies[1:])

	// 3) Create slices in two different ways
	firstSlice := myHobbies[0:2] // Using slice syntax
	secondSlice := myHobbies[:2] // Using shorthand syntax
	fmt.Println("First slice:", firstSlice)
	fmt.Println("Second slice:", secondSlice)

	// 4) Re-slice to get second and last elements
	reslicedHobbies := firstSlice[1:3]
	fmt.Println("Resliced hobbies:", reslicedHobbies)

	// 5) Create dynamic array (slice) for course goals
	courseGoals := []string{"Learn Go Programming", "Build Real Projects"}
	fmt.Println("Initial course goals:", courseGoals)

	// 6) Modify second goal and add third goal
	courseGoals[1] = "Master Go Fundamentals"
	courseGoals = append(courseGoals, "Create Web Applications")
	fmt.Println("Updated course goals:", courseGoals)

	// 7) Create Product struct and dynamic list of products
	type Product struct {
		title string
		id    int
		price float64
	}

	products := []Product{
		{title: "Book", id: 1, price: 29.99},
		{title: "Course", id: 2, price: 49.99},
	}

	products = append(products, Product{title: "Workshop", id: 3, price: 99.99})
	fmt.Println("Products:", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
