package Fonctions

var A, B []int
var Gestes []string

// pa prend la carte du sommet de B et la pose au sommet de A (rien si B est vide).
func pa() {
	if len(B) > 0 {
		A = append([]int{B[0]}, A...) // Ajoute B[0] au début de A
		B = B[1:]                     // Retire B[0] de B
	}
	Gestes = append(Gestes, "pa")
}

// pb prend la carte du sommet de A et la pose au sommet de B (rien si A est vide).
func pb() {
	if len(A) > 0 {
		B = append([]int{A[0]}, B...) // Ajoute A[0] au début de B
		A = A[1:]                     // Retire A[0] de A
	}
	Gestes = append(Gestes, "pb")
}

