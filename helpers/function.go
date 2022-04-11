package helpers

import "github.com/brianvoe/gofakeit/v6"

// Function returns random student struct using Faker
func InitStudent() Student {
	student := Student{}
	gofakeit.Struct(&student)

	return student
}
