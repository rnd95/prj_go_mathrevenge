package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsPerQuestion = 5
)

func main() {
	fmt.Println("Вітаємо у грі MATHREVENGE!")
	time.Sleep(1 * time.Second)

	for i := 5; i > 0; i-- {
		fmt.Println(i, " ")
		time.Sleep(1 * time.Second)
	}
	myPoints := 0

	start := time.Now()

	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)
		ans := " "
		fmt.Scan(&ans)
		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Try Again")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				points := totalPoints - myPoints
				fmt.Println("Чудово, ти набрав", myPoints, "очок!")
				fmt.Printf("Залишилося набрати %v!\n", points)
			} else {
				fmt.Println("Try Again!")

			}
		}
	}
	end := time.Now()
	timeSpent := end.Sub(start)

	fmt.Printf("Молодець, впорався всього за %v", timeSpent)
	time.Sleep(5 * time.Second)
}
