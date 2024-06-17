package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []int
}

type Class struct {
	Class    int
	Students []Student
}

func find(classes []int, grade int) bool {
	for _, value := range classes {
		if value == grade {
			return true
		}
	}
	return false
}

func printClasses(classMap map[int]*Class) {
	for grade, class := range classMap {
		fmt.Printf("Grade %v\n", grade)
		for _, students := range class.Students {
			fmt.Println(students.Name, students.Grades)
		}
	}
}

func main() {
	var name string
	var grade int
	classes := []int{1, 2, 3, 4, 5}
	classMap := make(map[int]*Class)
	var response string

	for {
		fmt.Println("Enter the student name:")
		fmt.Scan(&name)
		fmt.Println("Enter the grade in which you want to admit the student:")
		fmt.Scan(&grade)

		if find(classes, grade) {
			fmt.Printf("Student can be admitted to grade %d\n", grade)

			if class, exists := classMap[grade]; exists {
				class.Students = append(class.Students, Student{
					Name:   name,
					Grades: []int{},
				})
			} else {
				classMap[grade] = &Class{
					Class: grade,
					Students: []Student{
						{
							Name:   name,
							Grades: []int{},
						},
					},
				}
			}
		} else {
			fmt.Println("No grade is available for this")
		}

		fmt.Println("Do you want to do entry of new student (y/n)?")
		fmt.Scan(&response)
		if response == "n" {
			break
		}
	}

	fmt.Println("Do you want to view the classes (y/n)?")
	fmt.Scan(&response)
	if response == "y" {
		printClasses(classMap)
	}
}
