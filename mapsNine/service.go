package mapsnine

import "fmt"

func Sample() {
	fmt.Println("Hello world from Maps")

	//! making map of students name and grade
	studentGrades := make(map[string]int)
	studentGrades["Prince"] = 34
	studentGrades["Boab"] = 89
	studentGrades["Singh"] = 42
	studentGrades["Aman"] = 100

	fmt.Println(studentGrades)
	fmt.Println("Marks of Boab:", studentGrades["Boabs"])

	//! update
	studentGrades["Boab"] = 121

	//! delete
	delete(studentGrades, "Boab")
	fmt.Println(studentGrades)

	//! is exist
	grades, exists := studentGrades["David"]
	fmt.Println(exists, grades)

}
