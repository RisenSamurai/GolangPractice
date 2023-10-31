package models

import (
	"fmt"
	"net/http"
)

func Maps(w http.ResponseWriter, r *http.Request) {

	ShowGrades()

}

func ShowGrades() {

	type Students struct {
		Name    string
		Subject map[string]int
	}

	s := Students{}
	s.Subject = make(map[string]int)
	s.Name = "Sarah"
	s.Subject["Math"] = 80
	s.Subject["History"] = 90

	var subjects int
	var i int

	for key, value := range s.Subject {
		i++
		subjects += value
		fmt.Println(key+" ", value)
	}

	total := subjects / i

	fmt.Println("Student: "+s.Name+" Subject: ", s.Subject, "average: ", total)
}

func Book() {
	type Book struct {
		Title  string
		Author string
		Price  float32
	}

	books := []Book{

		{Title: "1984", Author: "George Orwell", Price: 10.0},
		{Title: "Brave New World", Author: "Aldous Huxley", Price: 12.0},
	}

	var booksPrice float32

	for i := 0; i < len(books); i++ {
		booksPrice += books[i].Price
		fmt.Println(books[i].Price)
	}

	fmt.Println("Total books price: ", booksPrice)

}

func mergeSlices() {

	var thirdSlice []int
	firstSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}
	thirdSlice = append(thirdSlice, firstSlice...)
	thirdSlice = append(thirdSlice, secondSlice...)

	for i := 0; i < len(thirdSlice); i++ {
		fmt.Println(thirdSlice[i])
	}
}

func sumOfSlice() {

	mySlice := []int{1, 2, 3}

	var amount int

	for i := 0; i < len(mySlice); i++ {
		amount += mySlice[i]
	}

	fmt.Println("Slice amount: ", amount)

}

func reverseFruits(a string, b string) {

	fruits := make(map[int64]string)

	fruits[1] = b
	fruits[2] = b

	for key, value := range fruits {
		fmt.Println(key, ": ", value)
	}

}

func returnFruits(a int64, b int64) {
	fruits := make(map[string]int64)

	fruits["apples"] = a
	fruits["bananas"] = b

	for key, value := range fruits {

		fmt.Println(key, ": ", value)
	}
}
