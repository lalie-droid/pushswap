package Fonctions

import "testing"

func TestPa(t *testing.T) {
	A := []int{1, 2, 3}
	B := []int{4, 5}
	PA(&A, &B)
	if len(A) != 4 || A[0] != 4 || A[1] != 1 || A[2] != 2 || A[3] != 3 {
		t.Errorf("Test 1 échoué : A = %v, attendu [4 1 2 3]", A)
	}
	if len(B) != 1 || B[0] != 5 {
		t.Errorf("Test 1 échoué : B = %v, attendu [5]", B)
	}

	// Test 2 : Ajout d'un élément depuis un tableau source vide
	A = []int{1, 2, 3}
	B = []int{}
	PA(&A, &B)
	if len(A) != 3 || A[0] != 1 || A[1] != 2 || A[2] != 3 {
		t.Errorf("Test 2 échoué : A = %v, attendu [1 2 3]", A)
	}
	if len(B) != 0 {
		t.Errorf("Test 2 échoué : B = %v, attendu []", B)
	}
}

//__________________________________________________________________________________________

func TestPb(t *testing.T) {
	// Test 1 : Ajout d'un élément depuis un tableau source non vide
	A := []int{1, 2, 3}
	B := []int{4, 5}
	PB(&A, &B)
	if len(B) != 3 || B[0] != 1 || B[1] != 4 || B[2] != 5 {
		t.Errorf("Test 1 échoué : B = %v, attendu [1 4 5]", B)
	}
	if len(A) != 2 || A[0] != 2 || A[1] != 3 {
		t.Errorf("Test 1 échoué : A = %v, attendu [2 3]", A)
	}

	// Test 2 : Ajout d'un élément depuis un tableau source vide
	A = []int{1, 2, 3}
	B = []int{}
	PB(&A, &B)
	if len(B) != 1 || B[0] != 1 {
		t.Errorf("Test 2 échoué : B = %v, attendu [1]", B)
	}
	if len(A) != 2 || A[0] != 2 || A[1] != 3 {
		t.Errorf("Test 2 échoué : A = %v, attendu [2 3]", A)
	}
}

func TestSa(t *testing.T) {
	// Test 1 : Échange de deux éléments
	A := []int{1, 2, 3}
	SA(A)
	if A[0] != 2 || A[1] != 1 || A[2] != 3 {
		t.Errorf("Test 1 échoué : A = %v, attendu [2 1 3]", A)
	}

	// Test 2 : Tableau avec un seul élément
	A = []int{5}
	SA(A)
	if A[0] != 5 {
		t.Errorf("Test 2 échoué : A = %v, attendu [5]", A)
	}

	// Test 3 : Tableau vide
	A = []int{}
	SA(A)
	if len(A) != 0 {
		t.Errorf("Test 3 échoué : A = %v, attendu []", A)
	}
}

func TestSb(t *testing.T) {
	// Test 1 : Échange de deux éléments
	B := []int{1, 2, 3}
	SB(B)
	if B[0] != 2 || B[1] != 1 || B[2] != 3 {
		t.Errorf("Test 1 échoué : B = %v, attendu [2 1 3]", B)
	}

	// Test 2 : Tableau avec un seul élément
	B = []int{5}
	SB(B)
	if B[0] != 5 {
		t.Errorf("Test 2 échoué : B = %v, attendu [5]", B)
	}

	// Test 3 : Tableau vide
	B = []int{}
	SB(B)
	if len(B) != 0 {
		t.Errorf("Test 3 échoué : B = %v, attendu []", B)
	}
}

func TestSs(t *testing.T) {
	// Test 1 : Échange simultané sur deux tableaux
	A := []int{1, 2, 3}
	B := []int{4, 5, 6}
	SS(A, B)
	if A[0] != 2 || A[1] != 1 || A[2] != 3 {
		t.Errorf("Test 1 échoué : A = %v, attendu [2 1 3]", A)
	}
	if B[0] != 5 || B[1] != 4 || B[2] != 6 {
		t.Errorf("Test 1 échoué : B = %v, attendu [5 4 6]", B)
	}
}

func TestRa(t *testing.T) {
	// Test 1 : Rotation vers le haut
	A := []int{9, 2, 4}
	RA(A)
	if len(A) != 3 || A[0] != 2 || A[1] != 4 || A[2] != 9 {
		t.Errorf("Test 1 échoué : A = %v, attendu [2 4 9]", A)
	}

	// Test 2 : Tableau avec un seul élément
	A = []int{5}
	RA(A)
	if len(A) != 1 || A[0] != 5 {
		t.Errorf("Test 2 échoué : A = %v, attendu [5]", A)
	}

}

func TestRb(t *testing.T) {
	// Test 1 : Rotation vers le haut
	B := []int{1, 2, 3}
	RB(B)
	if len(B) != 3 || B[0] != 2 || B[1] != 3 || B[2] != 1 {
		t.Errorf("Test 1 échoué : B = %v, attendu [2 3 1]", B)
	}

	// Test 2 : Tableau avec un seul élément
	B = []int{5}
	RB(B)
	if len(B) != 1 || B[0] != 5 {
		t.Errorf("Test 2 échoué : B = %v, attendu [5]", B)
	}

}

func TestRr(t *testing.T) {
	// Test 1 : Rotation simultanée vers le haut
	A := []int{1, 2, 3}
	B := []int{4, 5, 6}
	RR(A, B)
	if len(A) != 3 || A[0] != 2 || A[1] != 3 || A[2] != 1 {
		t.Errorf("Test 1 échoué : A = %v, attendu [2 3 1]", A)
	}
	if len(B) != 3 || B[0] != 5 || B[1] != 6 || B[2] != 4 {
		t.Errorf("Test 1 échoué : B = %v, attendu [5 6 4]", B)
	}
}

func TestRra(t *testing.T) {
	// Test 1 : Rotation vers le bas
	A := []int{1, 2, 3}
	RRA(A)
	if len(A) != 3 || A[0] != 3 || A[1] != 1 || A[2] != 2 {
		t.Errorf("Test 1 échoué : A = %v, attendu [3 1 2]", A)
	}

	// Test 2 : Tableau avec un seul élément
	A = []int{5}
	RRA(A)
	if len(A) != 1 || A[0] != 5 {
		t.Errorf("Test 2 échoué : A = %v, attendu [5]", A)
	}

	// Test 3 : Tableau vide
	A = []int{}
	RRA(A)
	if len(A) != 0 {
		t.Errorf("Test 3 échoué : A = %v, attendu []", A)
	}
}

func TestRrb(t *testing.T) {
	// Test 1 : Rotation vers le bas
	B := []int{1, 2, 3}
	RRB(B)
	if len(B) != 3 || B[0] != 3 || B[1] != 1 || B[2] != 2 {
		t.Errorf("Test 1 échoué : B = %v, attendu [3 1 2]", B)
	}

	// Test 2 : Tableau avec un seul élément
	B = []int{5}
	RRB(B)
	if len(B) != 1 || B[0] != 5 {
		t.Errorf("Test 2 échoué : B = %v, attendu [5]", B)
	}

	// Test 3 : Tableau vide
	B = []int{}
	RRB(B)
	if len(B) != 0 {
		t.Errorf("Test 3 échoué : B = %v, attendu []", B)
	}
}

func TestRrr(t *testing.T) {
	// Test 1 : Rotation simultanée vers le bas
	A := []int{1, 2, 3}
	B := []int{4, 5, 6}
	RRR(A, B)
	if len(A) != 3 || A[0] != 3 || A[1] != 1 || A[2] != 2 {
		t.Errorf("Test 1 échoué : A = %v, attendu [3 1 2]", A)
	}
	if len(B) != 3 || B[0] != 6 || B[1] != 4 || B[2] != 5 {
		t.Errorf("Test 1 échoué : B = %v, attendu [6 4 5]", B)
	}
}
