/*
Lab Assignment: Building an Exam Preparation Game
Given on Friday 020224
By Yaakov Miller
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// different stages of exam preparation
	stages := []string{
		"Gather your notes and study material",
		"Divide the material into equal sections",
		"Review the material",
	}

	// preparation activities
	activities := [][]string{
		{
			"Open your notes (hand written, books, articles and online links)",
			"Buy or borrow books from the library or eformat",
			"Use a new broswer window for or the links",
		},
		{
			"Make an inventory of the study material",
			"Try to divided by time periods acording to the exam schedule",
			"Consider adding time to reviews",
		},
		{
			"Do a brids view of the material",
			"Start the review as soon as posible",
			"Do multiple reviews if possible",
		},
	}
	// exam preparation tips
	tips := [][]string{
		{
			"Find a quite place, music if optional!",
			"Don't cram all the material in one session",
			"Take breaks",
		},
		{
			"Go by topics when possible",
			"Organize by diffulcty",
			"Give yourself enough time",
		},
		{
			"Start by the hardest part",
			"Give favorite part priority",
			"Avoid distractions",
		},
	}
	// Welcome msg
	fmt.Println("Start your engine to get ready for that Exam!!")
	// displaying stages
	for i, stage := range stages {
		fmt.Printf("%d - %s\n", i, stage)
	}
	for {
		// user option and validation
		var option int
		fmt.Print("\nMake a selection (enter number, 0 to exit): ")
		fmt.Scan(&option)

		if option >= 1 && option < 4 {
			fmt.Println("\nEnter your number.")
			fmt.Scan(&option)

			// randomly select a stage using the local random generator

			activity := selector(option, activities)
			tip := selector(option, tips)

			// outout the selected stage
			fmt.Printf("Your exam preparation activity and tip are: \n%s\n%s \n", activity, tip)
		} else if option == 0 {
			break
		} else {
			fmt.Println("Please only numbers 1 to 3 or x to exit")
			continue
		}

	}

}

func selector(option int, section [][]string) string {
	/* randomly chioces an element
	from the activities and tip sections*/
	// generating a ramdom choice from the user

	src := rand.NewSource(time.Now().UnixMicro() + int64(option))
	r := rand.New(src)
	index := r.Intn(len(section))
	// randomly select a topic using the local random generator
	selection := section[option][index]
	return selection

}
