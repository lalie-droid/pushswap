package Fonctions

// rra fait remonter la carte du fond de A jusqu'au sommet.
func rra() {
	if len(A) > 0 {
		fond := A[len(A)-1]
		A = append([]int{fond}, A[:len(A)-1]...) // Déplace la dernière carte au début
	}
	Gestes = append(Gestes, "rra")
}

func rrb() {
	if len(B) > 0 {
		fond := B[len(B)-1]
		B = append([]int{fond}, B[:len(B)-1]...) // Déplace la dernière carte au début
	}
	Gestes = append(Gestes, "rrb")
}

func rrr() {
	rra()
	rrb()
}
