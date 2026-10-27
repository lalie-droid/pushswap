package Fonctions

var A, B []int
var Gestes []string



// ra fait descendre la carte du sommet de A tout au fond.
func ra() {
	if len(A) > 0 {
		A = append(A[1:], A[0]) // Déplace la première carte à la fin
	}
	Gestes = append(Gestes, "ra")
}

// ra fait descendre la carte du sommet de A tout au fond.
func rb() {
	if len(B) > 0 {
		B = append(B[1:], B[0]) // Déplace la première carte à la fin
	}
	Gestes = append(Gestes, "rb")
}

func rr() {
	ra()
	rb()
}

