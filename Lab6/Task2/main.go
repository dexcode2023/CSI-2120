//Name: Dexter Day
//Student number: 300192263

package main

import "fmt"

// Define a suitable structure

type Course struct {
	NStudents int
	Avg       float64
	Professor string
}

func main() {
	// Create a dynamic map m
	// Add the courses CSI2120 and CSI2101 to the map
	m := make(map[string]Course)
	m["CSI2120"] = Course{300, 50, "Aziz"}
	m["CSI3131"] = Course{30, 90, "Rami"}

	for k, v := range m {
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.NStudents)
		fmt.Printf("Professor: %s\n", v.Professor)
		fmt.Printf(": %f\n\n", v.Avg)
	}
}
