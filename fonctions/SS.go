package Fonctions


func sa() {
	if len(A) >= 2 {
		A[0], A[1] = A[1], A[0] // Échange les deux premières cartes
	}
	Gestes = append(Gestes, "sa")
}

func sb() {
	if len(B) >= 2 {
		B[0], B[1] = B[1], B[0] // Échange les deux premières cartes
	}
	Gestes = append(Gestes, "sb")
}

func ss() {
	sa()
	sb()

}