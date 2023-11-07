package yams

import "testing"

var testCases = []struct {
	dice   []int
	figure string
	score  int
}{
	{[]int{2, 1, 3, 2, 2}, "Brelan", 28},
	{[]int{3, 2, 3, 3, 3}, "Carre", 35},
}

func TestAnalyseFigure(t *testing.T) {
	for _, elem := range testCases {
		figure, score := AnalyseFigure(elem.dice)

		if figure != elem.figure {
			t.Errorf("AnalyseFigure(%v) = %v, want %v", elem.dice, figure, elem.figure)
		}
		if score != elem.score {
			t.Errorf("AnalyseFigure(%v) = %v, want %v", elem.dice, score, elem.score)
		}
	}
}
