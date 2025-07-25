package main

import (
	"fmt"
)

func average(grades map[string]float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var userName string
	var numSubjects int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	fmt.Print("Enter number of subjects: ")
	_, err := fmt.Scanln(&numSubjects)
	if err != nil || numSubjects <= 0 {
		fmt.Println("Invalid number of subjects!")
		return
	}

	grades := make(map[string]float64)

	for i := 0; i < numSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter Subject title %d: ", i+1)
		fmt.Scanln(&subject)

		fmt.Print("Enter grade (0-100): ")
		_, err = fmt.Scanln(&grade)
		if err != nil {
			fmt.Println("Invalid input")
			i--
			continue
		}
		if grade < 0 || grade > 100 {
			fmt.Println("Grade must be between 0 and 100")
			i--
			continue
		}

		grades[subject] = grade
	}

	average := average(grades)

	fmt.Printf("\nGrade report for %s\n", userName)
	fmt.Println("---------------------------------")

	for subject, grade := range grades {
		fmt.Printf("Subject: %s    Grade: %.2f\n", subject, grade)
	}

	fmt.Println("---------------------------------")
	fmt.Printf("Average grade: %.2f\n", average)
}
