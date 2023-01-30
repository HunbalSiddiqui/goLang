package main

import "fmt"

func main() {
	welcome := "String"
	fmt.Println(welcome)
	fmt.Printf("Var is of type: %T \n", welcome)

	var courses = []string{"reactjs", "java", "go", "c#"}
	fmt.Println(courses)
	idx := 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)
}
