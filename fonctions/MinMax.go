package Fonctions

func min(pile []int) int {
	min := pile[0]
	for i := 0; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]

		}
	}
	return min
}

func max(pile []int) int {
	max := pile[0]
	for i := 0; i < len(pile); i++ {
		if pile[i] > max {
			max = pile[i]

		}
	}
	return max
}
