package Fonctions

func min(pile []int) (int, int) {
	min := pile[0]
	indexmin := 0
	for i := 0; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]

		}
	}
	return min, indexmin
}

func max(pile []int) (int, int) {
	max := pile[0]
	indexmax := 0
	for i := 0; i < len(pile); i++ {
		if pile[i] > max {
			max = pile[i]

		}
	}
	return max, indexmax
}

func med(pile []int) int {
	max := pile[0]
	for i := 0; i < len(pile); i++ {
		if pile[i] > max {
			max = pile[i]

		}
	}
	return max / 2
}
