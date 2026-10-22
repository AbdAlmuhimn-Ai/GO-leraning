package main

import (
	"example.com/exam04/rental"
	"example.com/exam04/report"
	"fmt"
)

func main() {
	items := []rental.Item{
		{
			Code:      "CAM",
			Name:      "Camera",
			DailyRate: 30,
			Available: 4,
			Rented:    0,
		},
		{
			Code:      "MIC",
			Name:      "Microphone",
			DailyRate: 10,
			Available: 2,
			Rented:    0,
		},
	}

	if err := rental.RentByCode(items, "CAM", 2); err == nil {
		fmt.Println("Rented CAM")
	}

	if err := rental.RentByCode(items, "MIC", 3); err != nil {
		fmt.Println(err)
	}

	if err := rental.RentByCode(items, "BAD", 1); err != nil {
		fmt.Println(err)
	}

	if err := rental.RentByCode(items, "CAM", 0); err != nil {
		fmt.Println(err)
	}

	if err := rental.ReturnByCode(items, "CAM", 1); err == nil {
		fmt.Println("Returned CAM")
	}

	if err := rental.ReturnByCode(items, "MIC", 1); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(items); i++ {
		fmt.Println(items[i].Code, items[i].Available, items[i].Rented)
	}

	rentedByCode := make(map[string]int)
	for i := 0; i < len(items); i++ {
		rentedByCode[items[i].Code] = items[i].Rented
	}
	fmt.Printf("Rental records: CAM=%d MIC=%d\n", rentedByCode["CAM"], rentedByCode["MIC"])

	setupService := rental.Service{
		Name:     "Setup",
		Fee:      15,
		Sessions: 2,
	}

	valuables := []report.Valuable{
		items[0],
		items[1],
		setupService,
	}

	for i := 0; i < len(valuables); i++ {
		fmt.Println(valuables[i].Value())
	}

	fmt.Println(report.Sum(valuables))
	fmt.Println(report.CountAbove(valuables, 25))
}
