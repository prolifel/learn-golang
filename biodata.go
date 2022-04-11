package main

import (
	"biodata/helpers"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Create a slice of student
var students []helpers.Student

// Append every instance of student to the slice until 10 students
func init() {
	log.Println("Init biodata")

	for i := 1; i <= 10; i++ {
		students = append(students, helpers.InitStudent())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor urut absensi>")
		os.Exit(1)
	}

	// Get arguments
	argument, _ := strconv.Atoi(os.Args[1])

	// Check if argument is more than 10
	if argument > 10 {
		fmt.Println("Nomor urut absensi tidak ada")
		os.Exit(1)
	}

	fmt.Printf("Nomor urut absensi: %v\nNama: %v\nAlamat: %v\nPekerjaan: %v\nAlasan memilih kelas Golang: %v", argument, students[argument-1].Name, students[argument-1].Address, students[argument-1].Job, students[argument-1].Reason)
}
