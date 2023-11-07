package yams

import "fmt"

type Figure struct {
	name  string
	score int
}

var figures = []Figure{
	{"Brelan", 28},
	{"Carre", 35},
	{"Full", 30},
	{"Grande Suite", 40},
	{"Yams", 50},
	{"Chance", 0},
}

func ParseDice(dices []int) []int {
	var values = []int{0, 0, 0, 0, 0}

	for _, dice := range dices {
		switch dice {
		case 1:
			values[0]++
		case 2:
			values[1]++
		case 3:
			values[2]++
		case 4:
			values[3]++
		case 5:
			values[4]++
		}
	}

	fmt.Println(values)

	return values
}

func AnalyseFigure(dices []int) (string, int) {
	var values = ParseDice(dices)

	if values[0] == 3 || values[1] == 3 || values[2] == 3 || values[3] == 3 || values[4] == 3 {
		return figures[0].name, figures[0].score
	} else if values[0] == 4 || values[1] == 4 || values[2] == 4 || values[3] == 4 || values[4] == 4 {
		return figures[1].name, figures[1].score
	}
	return "", 0
}
