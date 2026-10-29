package Fonctions

// Appliquer amène B[indexnombre] et sa cible en haut de leur pile, puis fait pa.
func Appliquer(A, B []int, indexnombre int) ([]int, []int) {
	// 1. On cherche la cible AVANT de bouger quoi que ce soit
	_, indexcible := TrouverCible(indexnombre, B, A)

	// 2. On remonte le nombre au sommet de B
	if indexnombre <= len(B)/2 {
		for i := 0; i < indexnombre; i++ {
			RB(B)
		}
	} else {
		for i := 0; i < len(B)-indexnombre; i++ {
			RRB(B)
		}
	}

	// 3. On remonte la cible au sommet de A
	if indexcible <= len(A)/2 {
		for i := 0; i < indexcible; i++ {
			RA(A)
		}
	} else {
		for i := 0; i < len(A)-indexcible; i++ {
			RRA(A)
		}
	}

	// 4. On pousse. PA(&A, &B) modifie les copies locales,
	// qu'on renvoie ensuite au main.
	PA(&A, &B)
	return A, B
}
