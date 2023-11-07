package main

import (
	"fmt"
	"tp-cleancode/yams"
)

var testCases = []struct {
	dice   []int
	figure string
	score  int
}{
	{[]int{2, 1, 3, 2, 2}, "Brelan", 28},
	{[]int{3, 2, 3, 3, 3}, "Carre", 35},
}

func main() {
	figure, value := yams.AnalyseFigure(testCases[0].dice)

	fmt.Println(figure, value)
}
